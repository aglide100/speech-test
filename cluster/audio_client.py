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
import os


options = [
    ('grpc.keepalive_time_ms', 900000),
    ('grpc.keepalive_permit_without_calls', True)
]


address = os.getenv("NODE_ADDRESS")
node_addresses = address.split(",")

node_status = {address: False for address in node_addresses}

# output_file = "./output/en_speaker_7.wav"
output_file = os.getenv("PATH")
speaker = os.getenv("SPEAKER")

token = os.getenv("TOKEN")


def set_node_status(address, is_busy):
    node_status[address] = not is_busy


def print_node_status():
    for address, is_busy in node_status.items():
        status = "Busy" if not is_busy else "Idle"
        print(f"Node at {address}: {status}")


def generate_audio_stubs():
    stubs = []
    for address in node_addresses:
        channel = grpc.insecure_channel(address, options=options)
        stub = audio_pb2_grpc.AudioGenerationServiceStub(channel)
        stubs.append((address, stub))
    return stubs


def divide_text_into_sentences(text):
    sentences = nltk.sent_tokenize(text)
    return sentences


def generate_audio(text_chunk, speaker, stub, node_address):
    response = stub.GenerateAudio(audio_pb2.Requirement(
        content=text_chunk, speaker=speaker, token=token))
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
        print(
            f"Node at {node_address} is not available. Retrying with another node.")
        avail_node.put(node_address)
        if avail_node.qsize() > 0:
            assign_work_to_node(sentence, speaker, idx, avail_node,
                                results, result_lock, running_threads)
        else:
            print("No available nodes. Waiting for a node to become available.")
            while avail_node.qsize() <= 0:
                time.sleep(1)
            assign_work_to_node(sentence, speaker, idx, avail_node,
                                results, result_lock, running_threads)
        return

    set_node_status(node_address, is_busy=False)
    avail_node.put(node_address)


def main():
    stubs = generate_audio_stubs()

    print(stubs)
    text = os.getenv("TEXT")
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
