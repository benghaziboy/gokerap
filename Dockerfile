FROM ubuntu:14.04

RUN apt-get update -qq
RUN apt-get install -y curl git bzr mercurial

RUN curl -s https://go.googlecode.com/files/go1.2.linux-amd64.tar.gz | tar -v -C /usr/local/ -xz

ENV PATH  /usr/local/go/bin:/usr/local/bin:/usr/local/sbin:/usr/bin:/usr/sbin:/bin:/sbin
ENV GOPATH  /go
ENV GOROOT  /usr/local/go

WORKDIR /go/src/gokerap
ADD . /go/src/gokerap

RUN go get github.com/lib/pq
RUN go get github.com/gorilla/mux
RUN go get code.google.com/p/go.crypto/bcrypt
RUN go install gokerap

EXPOSE 8080

ENTRYPOINT /go/bin/gokerap
