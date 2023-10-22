import os
import grpc
import time

import pb.svc.audio.audio_pb2 as audio_pb2
import pb.svc.audio.audio_pb2_grpc as audio_pb2_grpc
from scipy.io.wavfile import write as write_wav
from dotenv import load_dotenv
from transformers import AutoProcessor, BarkModel
import torch

options = [
    # ('grpc.keepalive_time_ms', 900000),
    ('grpc.keepalive_permit_without_calls', True)
]

load_dotenv()

address = os.getenv("SERVER_ADDRESS")
who = os.getenv("Local")
token = os.getenv("TOKEN")


def gen_grpc_stubs():
    channel = grpc.insecure_channel(address, options=options)
    stub = audio_pb2_grpc.AudioServiceStub(channel)
    return stub


def call_checking_jobs(stub: audio_pb2_grpc.AudioServiceStub):
    response = stub.CheckingJob(
        audio_pb2.CheckingJobReq(
            auth=audio_pb2.Auth(token=token, who=who)
        )
    )

    return response


def call_sending_result(stub: audio_pb2_grpc.AudioServiceStub, audio, token, who, content, speaker, id):
    request = audio_pb2.SendingResultReq(
        audio=audio_pb2.Audio(
            data=audio
        ),
        auth=audio_pb2.Auth(
            token=token,
            who=who
        ),
        job=audio_pb2.Job(
            content=content,
            speaker=speaker,
            id=id
        )
    )

    try:
        response = stub.SendingResult(request)
    except grpc.RpcError as e:
        print(f"Error: {e.code()}: {e.details()}")

    return response


def main(model, processor):
    stub = gen_grpc_stubs()

    while True:
        response = call_checking_jobs(stub)
        if len(str(response.error)) > 1:
            print(str(response.error))
        elif response.job is None or len(response.job.content) == 0:
            # pass
            time.sleep(1)
        else:
            job = response.job
            print("generate audio : /", job.content, "/, ", job.speaker)

            inputs = processor(job.content, voice_preset=job.speaker)

            audio_array = model.generate(**inputs)
            audio_array = audio_array.cpu().numpy().squeeze()
            # audio = generate_audio(job.content, history_prompt=job.speaker)

            write_wav('output.wav',
                      model.generation_config.sample_rate, audio_array)

            with open('output.wav', 'rb') as fd:
                serialized_audio = fd.read()
                print("sending : ", len(serialized_audio))
                call_sending_result(stub, serialized_audio, token,
                                    who, job.content, job.speaker, job.id)

        time.sleep(60)


if __name__ == '__main__':
    print("calling to ", address)

    device = "cuda:0" if torch.cuda.is_available() else "cpu"
    processor = AutoProcessor.from_pretrained("suno/bark")
    if device == "cuda:0":
        model = BarkModel.from_pretrained("suno/bark")
    else:
        model = BarkModel.from_pretrained("suno/bark-small")

    model = model.to(device)

    print("loaded!")
    while True:
        try:
            main(model, processor)
        except Exception as e:
            print(e)
            time.sleep(60)
