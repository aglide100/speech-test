#!/bin/bash

python3 -m grpc_tools.protoc -I. --python_out=. --grpc_python_out=. --pyi_out=. ./pb/svc/audio/audio.proto

protoc --go_out=../../../.. --go-grpc_out=../../../.. ./pb/svc/audio/audio.proto
