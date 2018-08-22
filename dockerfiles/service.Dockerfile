# Build image
FROM golang:1.10.3-alpine3.8 as builder

# Build dependencies
RUN apk add --no-cache \
    git \
    gcc \
    musl-dev

COPY . /go/src/github.com/im-kulikov/atlantio-task
WORKDIR /go/src/github.com/im-kulikov/atlantio-task

RUN \
    export VERSION=$(git rev-parse --verify HEAD) && \
    export BUILD=$(date -u +%s%N) && \
    export LDFLAGS="-w -s -X main.BuildVersion=${VERSION} -X main.BuildTime=${BUILD} -extldflags \"-static\"" && \
    export CGO_ENABLED=0 && \
    go build -v -ldflags "${LDFLAGS}" -o /go/bin/atlant ./

# Executable image
FROM scratch

WORKDIR /

COPY --from=builder /go/bin/atlant /atlant
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/src/github.com/im-kulikov/atlantio-task/config.yml /config.yml

CMD ["/atlant"]