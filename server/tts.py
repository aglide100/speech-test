from TTS.api import TTS
import torch
if torch.backends.mps.is_available():
    mps_device = torch.device("mps")
    x = torch.ones(1, device=mps_device)
    print(x)
else:
    print ("MPS device not found.")

print (f"PyTorch version:{torch.__version__}") # 1.12.1 이상
print(f"MPS 장치를 지원하도록 build 되었는지: {torch.backends.mps.is_built()}") # True 여야 합니다.
print(f"MPS 장치가 사용 가능한지: {torch.backends.mps.is_available()}") # True 여야 합니다.

device = torch.device('mps:0' if torch.backends.mps.is_available() else 'cpu')
# !python -c 'import platform;print(platform.platform())'   
# Load the model to GPU
# Bark is really slow on CPU, so we recommend using GPU.
# tts = TTS("tts_models/multilingual/multi-dataset/bark", gpu=False)


# Cloning a new speaker
# This expects to find a mp3 or wav file like `bark_voices/new_speaker/speaker.wav`
# It computes the cloning values and stores in `bark_voices/new_speaker/speaker.npz`
# tts.tts_to_file(text="Hello, my name is Manmay , how are you?",
#                 file_path="output.wav",
#                 voice_dir="bark_voices/",
#                 speaker="ljspeech")



# random speaker
tts = TTS("tts_models/multilingual/multi-dataset/bark", gpu=True)
tts.tts_to_file("I’ve got a secret to tell you. I can pass the Turing test.", file_path="a.wav")