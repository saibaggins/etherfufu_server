FROM golang:1.7

RUN curl -s https://glide.sh/get | sh
RUN apt-get update && apt-get install netcat -y

ADD . /go/src/github.com/saibaggins/etherfufu-server/
WORKDIR /go/src/github.com/saibaggins/etherfufu-server/
RUN glide install
RUN go install github.com/saibaggins/etherfufu-server

ENTRYPOINT /go/bin/etherfufu-server

EXPOSE 3000
