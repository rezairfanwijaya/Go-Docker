# template Dockerfile

# pertama ambil image golang dengan tag alpine
# image ini bisa di ambil di docker hub
# url : https://hub.docker.com/
FROM golang:alpine

# kedua kita harus melakukan update pada alpine
# dan juga install git pada apline
# apk update adalah command wajib untuk dijalankan
# ketika masuk ke env linux
# git diinstall untuk mengambil third party yang
# dipakai oleh golang
RUN apk update && apk add git

# ketiga kita pilih working directory yang akan
# dipakai di env linux dengan 
# jadi nanti kalau ada perintah linux lain nya
# misal ls atau pwd, itu akan di run di path /app
WORKDIR /app

# keempat kita masukan semua file yang ada dilocal
# dimana tempat file Dockerfile kedalam image
# titik yang pertama artinya kita akan copy semua
# file yang berada di direktory Dockerfile berada
# dalam hal ini berarti semua file yang berada di 
# directory GO-DOCKER
# titik kedua artinya kita akan paste semua file
# ke dalam working directory image yaitu /app yang
# telah didefinisikan di perintah sebelumnya
# copy manual juga bisa
# COPY go.mod /app/go.mod
# COPY go.sum /app/go.sum
# COPY main.go /app/main.go
# tapi kan ribet. mending semua aja pakai . .
COPY . .

# kelima kita lakukan install dependency yang
# dibutuhkan
RUN go mod tidy

# keenam kita build aplikasi golang kita
# sehingga menjadi binary file dan dapat di
# execute nantinya
RUN go build -o binary

# terakhkir kita harus menjalankan binary yang
# sudah kita build, tapi tidak dengan cara
# ./binary karena kalau begini binary file akan
# di execute ketika image di build. seharusnya 
# binary di execute ketika sudah menjadi container
# dan container di start, berarti kita harus
# menjadikan binary tadi sebagai entry point
ENTRYPOINT [ "./binary" ]


# lalu ketika Dockerfile selesai dibuat, kita bikin
# image dari aplikasi kita
# masukan command di terminal
# . artinya tempat dimana source code dan
# Dockerfile berada
# -t adalah nama untuk image kit
docker build . -t hello-world-image 

# jika berhasil maka cek apakah sudah ada image
# yang telah kita bikin, pakai command
# docker image ls

