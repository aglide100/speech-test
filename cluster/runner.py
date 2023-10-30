import gc
import os
import grpc
import time

import pb.svc.audio.audio_pb2 as audio_pb2
import pb.svc.audio.audio_pb2_grpc as audio_pb2_grpc
from scipy.io.wavfile import write as write_wav
from dotenv import load_dotenv
from transformers import AutoProcessor, BarkModel
import torch
import socket
import ffmpeg

options = [
    # ('grpc.keepalive_time_ms', 900000),
    ('grpc.keepalive_permit_without_calls', True)
]

output_wav = 'output_wav.wav'
output_aac = 'output.aac'
output_ts = 'output.ts'


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


def call_sending_result(stub: audio_pb2_grpc.AudioServiceStub, audio, millisec, token, who, content, speaker, id, no):
    request = audio_pb2.SendingResultReq(
        audio=audio_pb2.Audio(
            data=audio,
            millisec=millisec
        ),
        auth=audio_pb2.Auth(
            token=token,
            who=who
        ),
        job=audio_pb2.Job(
            content=content,
            speaker=speaker,
            id=id,
            no=no
        )
    )

    try:
        response = stub.SendingResult(request)
        return response
    except grpc.RpcError as e:
        print(f"Error: {e.code()}: {e.details()}")


def main():
    stub = gen_grpc_stubs()

    while True:
        response = call_checking_jobs(stub)
        if len(str(response.error)) > 1:
            print(str(response.error))
        elif response.job is None or len(response.job.content) == 0:
            # pass
            if model is not None:
                del model
                gc.collect()
                print("gc collected!")

            time.sleep(300)
        else:
            device = "cuda:0" if torch.cuda.is_available() else "cpu"

            if device == "cpu":
                print("load small model")
                processor = AutoProcessor.from_pretrained("suno/bark-small")
            else:
                processor = AutoProcessor.from_pretrained("suno/bark")

            model = BarkModel.from_pretrained("suno/bark")
            model = model.to(device)

            job = response.job
            print("generate audio : /", job.content, "/, ", job.speaker)

            inputs = processor(job.content, voice_preset=job.speaker)

            audio_array = model.generate(**inputs)
            audio_array = audio_array.cpu().numpy().squeeze()
            millisec = (len(audio_array) /
                        model.generation_config.sample_rate) * 1000
            write_wav(output_wav,
                      model.generation_config.sample_rate, audio_array)

            input_wav_stream = ffmpeg.input(output_wav)
            output_aac_stream = ffmpeg.output(
                input_wav_stream, output_aac, codec='aac')
            ffmpeg.run(output_aac_stream, overwrite_output=True)

            audio_stream = ffmpeg.input(output_aac)
            output_ts_stream = ffmpeg.output(audio_stream, output_ts)
            ffmpeg.run(output_ts_stream, overwrite_output=True)

            with open(output_ts, 'rb') as fd:
                serialized_audio = fd.read()
                print("sending : ", len(serialized_audio))
                call_sending_result(stub, serialized_audio, millisec, token,
                                    who, job.content, job.speaker, job.id, job.no)

        time.sleep(60)


if __name__ == '__main__':
    load_dotenv()

    address = os.getenv("SERVER_ADDRESS")
    who = socket.gethostname()
    token = os.getenv("TOKEN")

    while True:
        try:
            main()
        except Exception as e:
            print(e)
            time.sleep(60)
