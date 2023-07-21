FROM golang:1.19.11-alpine

WORKDIR /bem-te-vi

COPY . /bem-te-vi

RUN go mod download && \
    go build -o /usr/local/bin/app

ENTRYPOINT ["/usr/local/bin/app"]
