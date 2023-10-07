ARG GO_VERSION=1.21-alpine3.18
ARG FROM_IMAGE=alpine:3.18

FROM golang:${GO_VERSION} AS builder

ARG TARGETOS
ARG TARGETARCH
ARG VERSION

LABEL org.opencontainers.image.source="https://github.com/omegion/vault-unseal"

WORKDIR /app

COPY ./ /app

RUN apk update && \
  apk add ca-certificates gettext git make curl unzip && \
  rm -rf /tmp/* && \
  rm -rf /var/cache/apk/* && \
  rm -rf /var/tmp/*


RUN make build TARGETOS=$TARGETOS TARGETARCH=$TARGETARCH VERSION=$VERSION

FROM ${FROM_IMAGE}

COPY --from=builder /app/dist/vault-unseal /bin/vault-unseal

ENTRYPOINT ["vault-unseal"]
