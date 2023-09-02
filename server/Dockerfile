FROM python:3.10
WORKDIR /usr/src/app

RUN apt-get update && apt-get install -y ffmpeg



RUN pip3 install git+https://github.com/huggingface/transformers.git

COPY b.py .


CMD ["b.py"]
ENTRYPOINT ["python3"]