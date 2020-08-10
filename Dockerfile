FROM golang:1.12.7-alpine3.9
LABEL  ZuoGuocai "zuoguocai@126.com"

WORKDIR $GOPATH/src/github.com/getRealip
ADD . $GOPATH/src/github.com/getRealip
RUN go build .
EXPOSE 12345
ENTRYPOINT  ["./getRealip"]
