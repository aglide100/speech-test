import grpc
import time
from concurrent import futures
import audio_pb2
import audio_pb2_grpc
from bark import generate_audio, preload_models
import pickle
import argparse
import os

_ONE_DAY_IN_SECONDS = 60 * 60 * 24

TOKEN = os.getenv("TOKEN")


class AudioGenerationServicer(audio_pb2_grpc.AudioGenerationServiceServicer):
    def GenerateAudio(self, request, context):
        if request.token != TOKEN:
            print("wrong token, ", request.token)
            return audio_pb2.Audio(error="not valid")

        text = request.content
        print("received : ", text)
        audio = generate_audio(text, history_prompt=request.speaker)
        serialized_audio = pickle.dumps(audio)
        print("sending : ", len(serialized_audio))

        return audio_pb2.Audio(data=serialized_audio)


def serve(port):
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    audio_pb2_grpc.add_AudioGenerationServiceServicer_to_server(
        AudioGenerationServicer(), server)
    server.add_insecure_port('[::]:' + str(port))
    server.start()
    try:
        while True:
            time.sleep(_ONE_DAY_IN_SECONDS)
    except KeyboardInterrupt:
        server.stop(0)


if __name__ == '__main__':
    parser = argparse.ArgumentParser(description='gRPC Audio Server')
    parser.add_argument('--port', type=int, default=50052, help='Server port')
    args = parser.parse_args()
    preload_models()
    print("loaded models")
    print(f"Starting gRPC server on port {args.port}")
    serve(args.port)
