ARG GO_VERSION=1.16-alpine3.12
ARG FROM_IMAGE=alpine:3.14.3

FROM golang:${GO_VERSION} AS builder

LABEL org.opencontainers.image.source="https://github.com/omegion/vault-unseal"

WORKDIR /app

COPY ./ /app

RUN apk update && \
  apk add ca-certificates gettext git make curl unzip && \
  rm -rf /tmp/* && \
  rm -rf /var/cache/apk/* && \
  rm -rf /var/tmp/*

RUN make build-for-container

FROM ${FROM_IMAGE}

COPY --from=builder /app/dist/vault-unseal-linux /bin/vault-unseal

ENTRYPOINT ["vault-unseal"]
