FROM golang:1.12

# Add Tini
ENV TINI_VERSION v0.18.0
ADD https://github.com/krallin/tini/releases/download/${TINI_VERSION}/tini /tini
RUN chmod +x /tini
ENTRYPOINT ["/tini", "--"]

WORKDIR /app

COPY go.mod /app
COPY go.sum /app

RUN go get -u ./...

COPY . /app

EXPOSE 3000

CMD go build -o app && ./app
