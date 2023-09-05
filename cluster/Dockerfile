FROM python:3.10

RUN  python -m pip install --upgrade pip

RUN python -m pip install grpcio

RUN pip install git+https://github.com/suno-ai/bark.git
RUN pip install --upgrade google-api-python-client

COPY . .
