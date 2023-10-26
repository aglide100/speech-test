import scipy
from transformers import AutoProcessor, BarkModel

device = "cuda:0" if torch.cuda.is_available() else "cpu"

if device == "cpu":
    print("load small model")
    processor = AutoProcessor.from_pretrained("suno/bark-small")
else:
    processor = AutoProcessor.from_pretrained("suno/bark")

model = BarkModel.from_pretrained("suno/bark")
model = model.to(device)

voice_preset = "v2/en_speaker_6"

inputs = processor("Hello, my dog is cute", voice_preset=voice_preset)

audio_array = model.generate(**inputs)
audio_array = audio_array.cpu().numpy().squeeze()

sample_rate = model.generation_config.sample_rate  # 샘플 속도 (샘플/초)
audio_length_seconds = len(audio_array) / sample_rate  # 재생 시간 (초)

print("오디오 재생 시간:", audio_length_seconds * 1000, "초")

# sample_rate = model.generation_config.sample_rate
# scipy.io.wavfile.write("bark_out.wav", rate=sample_rate, data=audio_array)
