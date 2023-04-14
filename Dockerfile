FROM golang:1.18

ARG APP

WORKDIR /usr/src/app

COPY go.mod .

COPY go.sum .

RUN go mod download

RUN go mod verify

COPY . .

RUN go build -v -o /usr/local/bin/app

CMD ["app"]
