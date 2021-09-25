FROM golang:latest
WORKDIR /go/src
RUN GOBIN=/tmp/ go get github.com/go-delve/delve/cmd/dlv@master \
 && mv /tmp/dlv $GOPATH/bin/dlv-dap
##    go get github.com/uudashr/gopkgs/v2/cmd/gopkgs \
##    go get github.com/ramya-rao-a/go-outline \
##    go get github.com/cweill/gotests/... \
##    go get github.com/fatih/gomodifytags \
##    go get github.com/josharian/impl \
##    go get github.com/haya14busa/goplay/cmd/goplay \
##    go get github.com/go-delve/delve/cmd/dlv \
##    go get github.com/golangci/golangci-lint/cmd/golangci-lint \
##    go get golang.org/x/tools/gopls
## ADD . /workspace