FROM golang:1.17
WORKDIR /go/src
RUN apt-get update && apt install build-essential -y
CMD ["tail", "-f", "/dev/null"]