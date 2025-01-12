# Content managed by Project Forge, see [projectforge.md] for details.
FROM golang:alpine

LABEL "org.opencontainers.image.authors"="Kyle"
LABEL "org.opencontainers.image.source"="https://github.com/kyleu/projectforge"
LABEL "org.opencontainers.image.vendor"="kyleu"
LABEL "org.opencontainers.image.title"="Project Forge"
LABEL "org.opencontainers.image.description"="Build and maintain feature-rich applications using Golang"

RUN apk add --update --no-cache ca-certificates tzdata bash curl htop libc6-compat

RUN apk add --no-cache ca-certificates dpkg gcc git musl-dev \
    && mkdir -p "$GOPATH/src" "$GOPATH/bin" \
    && chmod -R 777 "$GOPATH"

RUN go install github.com/go-delve/delve/cmd/dlv@latest

SHELL ["/bin/bash", "-c"]

# main http port
EXPOSE 40000
# marketing port
EXPOSE 40001

WORKDIR /

ENTRYPOINT ["/projectforge", "-a", "0.0.0.0"]

COPY projectforge /
