# Here would be great to see the multistage build to demostrate that the app can run without the go dependency in a small container like alpine
FROM golang:1.20

COPY . /app
WORKDIR /app

RUN go build ./cmd/wjug/...

ENTRYPOINT ["./wjug"]
