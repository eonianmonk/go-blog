FROM golang:1.21.3-alpine

WORKDIR /


COPY blogo blogo 

#RUN make build