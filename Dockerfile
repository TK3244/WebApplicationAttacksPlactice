FROM golang:latest

WORKDIR /go

ENV GOPATH=

EXPOSE 8080

RUN cd /go

CMD ["go", "run", "."]