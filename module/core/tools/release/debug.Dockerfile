FROM golang:alpine

RUN apk add --update --no-cache ca-certificates tzdata bash curl htop

RUN apk add --no-cache ca-certificates dpkg gcc git musl-dev \
    && mkdir -p "$GOPATH/src" "$GOPATH/bin" \
    && chmod -R 777 "$GOPATH" \
    && go get github.com/go-delve/delve/cmd/dlv

SHELL ["/bin/bash", "-c"]

ENTRYPOINT ["/{{{ .Exec }}}", "-a", "0.0.0.0"]
EXPOSE {{{ .Port }}}
COPY {{{ .Exec }}} /
