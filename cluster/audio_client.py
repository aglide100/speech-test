import grpc
import audio_pb2
import audio_pb2_grpc
from scipy.io import wavfile
import numpy as np
import pickle
from bark import SAMPLE_RATE
import nltk
import threading
import threading
import queue
import time
import multiprocessing

node_addresses = ['0.0.0.0:50052', '192.168.0.72:50053']
node_status = {address: False for address in node_addresses}

output_file = "en_speaker_7.wav"
speaker = "v2/en_speaker_7"


def set_node_status(address, is_busy):
    node_status[address] = not is_busy


def print_node_status():
    for address, is_busy in node_status.items():
        status = "Busy" if not is_busy else "Idle"
        print(f"Node at {address}: {status}")


def generate_audio_stubs():
    stubs = []
    for address in node_addresses:
        channel = grpc.insecure_channel(address)
        stub = audio_pb2_grpc.AudioGenerationServiceStub(channel)
        stubs.append((address, stub))
    return stubs


def divide_text_into_sentences(text):
    sentences = nltk.sent_tokenize(text)
    return sentences


def generate_audio(text_chunk, speaker, stub, node_address):
    response = stub.GenerateAudio(audio_pb2.Requirement(
        content=text_chunk, speaker=speaker))
    set_node_status(node_address, is_busy=False)
    return response.data


def assign_work_to_node(sentence, speaker, idx, avail_node, results, result_lock, running_threads):
    node_address = avail_node.get()

    set_node_status(node_address, is_busy=True)
    stub = audio_pb2_grpc.AudioGenerationServiceStub(
        grpc.insecure_channel(node_address))
    try:
        response = generate_audio(sentence, speaker, stub, node_address)
        result_lock.acquire()
        results[idx] = response
        running_threads.pop()
        result_lock.release()
    except grpc.RpcError:
        print(f"Node at {node_address} is not available.")
        assign_work_to_node(sentence, idx, avail_node,
                            results, result_lock, running_threads)
        return

    set_node_status(node_address, is_busy=False)
    avail_node.put(node_address)


def main():
    stubs = generate_audio_stubs()

    print(stubs)

    text = """I'm Sarah, a perpetual explorer of life's possibilities! You might hear me chuckling [laughs], or maybe even humming a tune as I dive into creative endeavors. My days are a symphony of curiosity and learning, from coding projects to culinary experiments. Sometimes I break into spontaneous dance moves, and when I'm deep in thought, you might catch me going "hmm" [thoughtful sound]. Adventure calls my name, whether it's through thrilling books or outdoor escapades. My camera is my constant companion, capturing both breathtaking landscapes and candid moments. [clears throat] Life's a canvas, and I'm here to paint it with bold strokes!"""
    manager = multiprocessing.Manager()
    results = manager.dict()
    result_lock = manager.Lock()
    running_threads = []

    avail_node = queue.Queue()
    for address, is_busy in node_status.items():
        if not is_busy:
            avail_node.put(address)

    sentences = divide_text_into_sentences(text)
    for idx, sentence in enumerate(sentences):
        if len(sentence.strip()) > 0:
            if avail_node.qsize() <= 0:
                while avail_node.qsize() <= 0:
                    time.sleep(1)

            print("Text : ", sentence)
            thread = threading.Thread(
                target=assign_work_to_node, args=(sentence, speaker, idx, avail_node, results, result_lock, running_threads))
            thread.start()
            running_threads.append(thread)

    for thread in running_threads:
        thread.join()

    audio_arrays = []
    for val in results.values():
        re = pickle.loads(val)

        audio_arrays.append(re)

    combined_audio = np.concatenate(audio_arrays)

    print("write file")
    wavfile.write(output_file, SAMPLE_RATE, combined_audio)


if __name__ == '__main__':
    nltk.download('punkt')
    main()
