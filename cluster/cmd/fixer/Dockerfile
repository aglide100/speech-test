FROM golang:1.18 AS builder

RUN apt update && apt install -y upx

WORKDIR "/go/src/github.com/aglide100/speech-test"

COPY go.mod .
COPY go.sum .

RUN mkdir -p /opt/bin/fixer/

RUN go mod tidy \
    && go get -u -d -v ./...

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-s -w' -o /opt/bin/fixer/main cmd/fixer/main.go \    
    && upx --lzma /opt/bin/fixer/main

FROM debian:stretch-slim AS runtime
COPY --from=builder /opt/bin/fixer /opt/bin/fixer

CMD [ "/opt/bin/fixer/main" ]
