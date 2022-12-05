FROM golang:alpine

RUN apk update && apk add git

ARG author="reza irfan wijaya"

ENV AUTHOR=${author}

RUN echo "Selamat datang diaplikasi pertama saya"
RUN echo "created by ${author}"

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o binary

ENTRYPOINT [ "./binary" ]