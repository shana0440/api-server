FROM golang:1.12

WORKDIR /app

COPY go.mod /app
COPY go.sum /app

RUN go get -u ./...

COPY . /app

CMD go build -o app && ./app
