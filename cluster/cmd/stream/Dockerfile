FROM golang:1.20-alpine3.17 AS builder

RUN apk update && apk --no-cache add upx ca-certificates ffmpeg

WORKDIR "/go/src/github.com/aglide100/speech-test"

COPY go.mod .
COPY go.sum .

RUN mkdir -p /opt/bin/stream/

RUN go mod tidy \
    && go get -u -d -v ./...

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-s -w' -o /opt/bin/stream/main cmd/stream/main.go \    
    && upx --lzma /opt/bin/stream/main

FROM debian:stretch-slim AS runtime
COPY --from=builder /opt/bin/stream /opt/bin/stream
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

CMD [ "/opt/bin/stream/main" ]