FROM golang:alpine

RUN apk update && apk add git

ENV AUTHOR="reza irfan wijaya"

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o binary

ENTRYPOINT [ "./binary" ]