import numpy as np
from scipy.io.wavfile import write as write_wav

text_path = 'audio.txt'
output_path = 'output.wav'


f = open(text_path, 'r')
lines = f.readlines()


for line in lines:
    print(line)

sample_rate = 44100

write_wav('output.wav', sample_rate, lines)
