FROM python:3.10

WORKDIR /app

COPY ./requirements.txt .

RUN python -m pip install --upgrade pip

RUN pip install -r requirements.txt

RUN pip install --upgrade google-api-python-client

# RUN pip install git+https://github.com/suno-ai/bark.git

COPY . .
