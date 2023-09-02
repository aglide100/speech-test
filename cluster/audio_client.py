import grpc
import audio_pb2
import audio_pb2_grpc

def generate_audio_stub():
    channel = grpc.insecure_channel('localhost:50051')
    return audio_pb2_grpc.AudioGenerationServiceStub(channel)

def main():
    stub = generate_audio_stub()

    text = "Your long text goes here..."
    chunk_size = 1000  # Adjust the chunk size as needed
    chunks = [text[i:i+chunk_size] for i in range(0, len(text), chunk_size)]

    audio_chunks = []
    for chunk in chunks:
        response = stub.GenerateAudio(audio_pb2.TextChunk(content=chunk))
        audio_chunks.append(response.data)

    complete_audio = b''.join(audio_chunks)

    # Process or save the complete_audio

if __name__ == '__main__':
    main()
