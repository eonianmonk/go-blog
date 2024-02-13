FROM golang:1.21.3-bullseye

WORKDIR /

USER root
RUN apt-get update \
 && DEBIAN_FRONTEND=noninteractive \
    apt-get install --no-install-recommends --assume-yes \
      docker.io

COPY blogo blogo 

#RUN make build