import grpc
import time
from concurrent import futures
import audio_pb2
import audio_pb2_grpc

_ONE_DAY_IN_SECONDS = 60 * 60 * 24


class AudioGenerationServicer(audio_pb2_grpc.AudioGenerationServiceServicer):
    def GenerateAudio(self, request, context):
        text_chunk = request.content
        # Generate audio using text_chunk and return Audio object
        audio_data = b"Generated audio data"  # Replace with actual audio data
        return audio_pb2.Audio(data=audio_data)


def divide_text_into_sentences(text, chunk_size):
    sentences = text.split(". ")
    text_chunks = []
    chunk = ""
    for sentence in sentences:
        if len(chunk) + len(sentence) < chunk_size:
            chunk += sentence + ". "
        else:
            text_chunks.append(chunk)
            chunk = sentence + ". "
    if chunk:
        text_chunks.append(chunk)
    return text_chunks


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    audio_pb2_grpc.add_AudioGenerationServiceServicer_to_server(
        AudioGenerationServicer(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    try:
        while True:
            time.sleep(_ONE_DAY_IN_SECONDS)
    except KeyboardInterrupt:
        server.stop(0)


if __name__ == '__main__':
    text = "testing"
    chunk_size = 100
    sentences = divide_text_into_sentences(text, chunk_size)
