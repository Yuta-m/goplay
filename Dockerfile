FROM golang:latest
RUN GOBIN=/tmp/ go get github.com/go-delve/delve/cmd/dlv@master \
 && mv /tmp/dlv $GOPATH/bin/dlv-dap
RUN go get github.com/golangci/golangci-lint/cmd/golangci-lint
RUN go get -v golang.org/x/tools/gopls
WORKDIR /go
