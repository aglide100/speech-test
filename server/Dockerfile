FROM python
WORKDIR /usr/src/app

RUN apt-get update && apt-get install -y ffmpeg

RUN pip3 install TTS 

COPY main.py .


CMD ["main.py"]
ENTRYPOINT ["python3"]