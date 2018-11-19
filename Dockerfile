FROM golang:1.11-alpine3.8 as builder

ENV buildir=/go/src/github.com/mindera/ready-set-go-workshop

COPY app $buildir/app

WORKDIR $buildir/app/cmd/webapp

RUN apk add --update gcc g++ && \
    go build --ldflags '-s -w -linkmode external -extldflags "-static"' -o ../../build/webapp

FROM scratch
LABEL maintainer="Tiago Pires <tiago.pires@mindera.com>"

ENV buildir=/go/src/github.com/mindera/ready-set-go-workshop

WORKDIR /app

COPY --from=builder ["$buildir/app/build/webapp", "/app/"]

ENTRYPOINT ["/app/webapp"]
