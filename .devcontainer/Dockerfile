FROM golang:1.19.1

WORKDIR /go/src/app

ENV GO111MODULE=on

# VSCodeの拡張機能で使うもの
RUN go install -v golang.org/x/tools/gopls@latest \
    && go install -v golang.org/x/tools/gopls@latest