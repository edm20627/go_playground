FROM golang:1.16

WORKDIR /go/src/github.com/edm20627/go_playground

RUN apt-get update -qq && apt-get install -y vim git

RUN go install github.com/pilu/fresh@latest
RUN go install golang.org/x/tools/cmd/goimports@latest
RUN go install github.com/kisielk/errcheck@latest
RUN go install github.com/gostaticanalysis/skeleton@latest

CMD ["fresh"]