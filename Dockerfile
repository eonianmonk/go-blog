FROM golang:1.21.3-bullseye

WORKDIR /

COPY . . 

RUN make build