import os
import grpc
import time

import pickle
import pb.svc.audio.audio_pb2 as audio_pb2
import pb.svc.audio.audio_pb2_grpc as audio_pb2_grpc

from bark import generate_audio, preload_models
options = [
    # ('grpc.keepalive_time_ms', 900000),
    ('grpc.keepalive_permit_without_calls', True)
]

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
            auth=audio_pb2.Auth(token=token, who="")
        )
    )

    return response


def call_sending_result(stub: audio_pb2_grpc.AudioServiceStub, audio, token, who, content, speaker, id):
    response = stub.SendingResult(
        audio_pb2.Audio(
            data=audio
        ),
        audio_pb2.Auth(
            token=token,
            who=who
        ),
        audio_pb2.Job(
            content=content,
            speaker=speaker,
            id=id
        )
    )

    return response


def main():
    stub = gen_grpc_stubs()

    while True:
        response = call_checking_jobs(stub)
        if len(str(response.error)) > 1:
            print(str(response.error))
        else:
            job = response.job
            print("generate audio", job.content, job.speaker)

            audio = generate_audio(job.content, history_prompt=job.speaker)

            serialized_audio = pickle.dumps(audio)
            err = call_sending_result(
                stub, serialized_audio, token, who, job.content, job.speaker, job.id)
            if len(str(err.msg)) <= 1:
                print("sended, ", len(serialized_audio))

        time.sleep(60)


if __name__ == '__main__':
    # nltk.download('punkt')
    preload_models()
    print("loaded!")
    main()
