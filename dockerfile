FROM golang:1.20

COPY . /app
WORKDIR /app

RUN go build ./cmd/wjug/...

ENTRYPOINT ["./wjug"]