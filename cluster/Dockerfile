FROM python:3.10

RUN python -m pip install --upgrade pip

RUN python -m pip install grpcio nltk

RUN pip install --upgrade google-api-python-client

RUN pip install git+https://github.com/suno-ai/bark.git

COPY . .
