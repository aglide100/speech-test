
import re
import scipy
from transformers import AutoProcessor, BarkModel
import ffmpeg
import torch
from scipy.io.wavfile import write as write_wav

# device = "cuda:0" if torch.cuda.is_available() else "cpu"

# if device == "cpu":
#     print("load small model")
#     processor = AutoProcessor.from_pretrained("suno/bark-small")
# else:
#     processor = AutoProcessor.from_pretrained("suno/bark")

# model = BarkModel.from_pretrained("suno/bark")
# model = model.to(device)

# voice_preset = "v2/en_speaker_6"

# inputs = processor("Hello, my dog is cute", voice_preset=voice_preset)

# audio_array = model.generate(**inputs)
# audio_array = audio_array.cpu().numpy().squeeze()

# sample_rate = model.generation_config.sample_rate  # 샘플 속도 (샘플/초)
# audio_length_seconds = len(audio_array) / sample_rate  # 재생 시간 (초)

# print("오디오 재생 시간:", audio_length_seconds * 1000, "초")


# sample_rate = model.generation_config.sample_rate
# write_wav("bark_out.wav", rate=sample_rate, data=audio_array)


input_wav = 'bark_out.wav'
output_aac = 'output.aac'
output_ts = 'output.ts'

input_wav_stream = ffmpeg.input(input_wav)
output_aac_stream = ffmpeg.output(input_wav_stream, output_aac, codec='aac')
ffmpeg.run(output_aac_stream, overwrite_output=True)


audio_stream = ffmpeg.input(output_aac)

output_ts_stream = ffmpeg.output(audio_stream, output_ts)
ffmpeg.run(output_ts_stream, overwrite_output=True)


pattern = r'#EXTINF:(\d+\.\d+),'
with open(output_ts, 'rb') as file:
    content = file.read().decode('utf-8', errors='ignore')  # 'utf-8' 인코딩 오류를 무시

extinf_match = re.search(pattern, content)

if extinf_match:
    extinf_value = float(extinf_match.group(1))
    print(f"EXTINF value for {output_ts}: {extinf_value} seconds")
else:
    print("EXTINF value not found in the segment file.")

input_stream = ffmpeg.input(output_ts)

probe = ffmpeg.probe(output_ts, v='error',
                     select_streams='a:0', show_entries='format=duration')
duration = float(probe['format']['duration'])

print(f"duration: {duration} 초")
