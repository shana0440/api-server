FROM golang:1.12

WORKDIR /app

COPY go.mod /app
COPY go.sum /app

RUN go get -u ./...

COPY . /app

EXPOSE 3000

CMD go build -o app && ./app
