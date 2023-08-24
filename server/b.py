from bark import SAMPLE_RATE, generate_audio, preload_models
import numpy as np
from scipy.io import wavfile
import os

# or export SUNO_ENABLE_MPS=true
os.environ["SUNO_ENABLE_MPS"] = "True"
# os.environ["SUNO_OFFLOAD_CPU"] = "True"
# os.environ["SUNO_USE_SMALL_MODELS"] = "True"

preload_models()

chunk_size = 100

full_text = """
    The exploration of outer space has always fascinated humanity. From the first time we looked up at the stars and wondered about what lies beyond our world, to the incredible achievements of space missions and the quest for knowledge. Space travel has brought us closer to understanding our place in the universe and has paved the way for technological advancements that impact our lives on Earth. As we continue to push the boundaries of space exploration, we open doors to new possibilities and expand our understanding of the cosmos.
     The beauty of nature's landscapes never fails to captivate our senses. Whether it's the serene tranquility of a calm lake reflecting the surrounding mountains, or the vibrant colors of a sunset painting the sky in hues of orange and pink, nature's artistry leaves us in awe. From the smallest details in a delicate flower to the grandeur of a towering forest, the intricate designs of the natural world remind us of the marvels that exist all around us. It's a reminder to appreciate the simple joys of life and the wonders of the Earth we call home.
    Hello, my name is Alex. I'm a passionate learner and explorer of the world around me. From a young age, I've been fascinated by technology and its potential to shape our future. I believe that every challenge is an opportunity for growth, and I'm always excited to take on new challenges head-on. When I'm not immersed in coding or solving problems, you can find me hiking in the great outdoors, capturing moments through my camera lens, or enjoying a good book. I value connections with people and am always open to new experiences and perspectives. Life is a journey, and I'm on a quest to make the most of every moment.
    Greetings, I'm Sarah, a perpetual explorer of life's possibilities! You might hear me chuckling [laughs], or maybe even humming a tune as I dive into creative endeavors. My days are a symphony of curiosity and learning, from coding projects to culinary experiments. Sometimes I break into spontaneous dance moves, and when I'm deep in thought, you might catch me going "hmm" [thoughtful sound]. Adventure calls my name, whether it's through thrilling books or outdoor escapades. My camera is my constant companion, capturing both breathtaking landscapes and candid moments. [clears throat] Life's a canvas, and I'm here to paint it with bold strokes!
    Howdy, I'm Max, a perpetual enthusiast on a quest for life's wonders! You might find me bursting into laughter [laughs] or uttering an "aha" when I stumble upon a new idea. My days are a whirlwind of exploration, whether I'm writing code or crafting new recipes. Sometimes I break into impromptu dance moves, and when I'm in my zone, you'll hear me softly humming along [humming sound]. Adventure beckons, whether it's within the pages of a captivating book or the trails of a scenic hike. My camera is my faithful companion, capturing both breathtaking landscapes and candid moments. [clears throat] Life's a grand adventure, and I'm here to experience every twist and turn!
"""


sentences = full_text.split(". ")
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

audio_arrays = []
for text_chunk in text_chunks:
    audio = generate_audio(text_chunk, history_prompt="v2/en_speaker_1")
    audio_arrays.append(audio)

combined_audio = np.concatenate(audio_arrays)

output_file = "en_speaker_1.wav"
wavfile.write(output_file, SAMPLE_RATE, combined_audio)