FROM golang:1.21.13-alpine3.20@sha256:2414035b086e3c42b99654c8b26e6f5b1b1598080d65fd03c7f499552ff4dc94 AS builder
ARG VERSION=dev-snapshot
RUN mkdir /build
RUN apk add --update make
WORKDIR /build
ADD . /build
RUN make BUILD_VERSION=${VERSION}

FROM alpine:3.20.3@sha256:beefdbd8a1da6d2915566fde36db9db0b524eb737fc57cd1367effd16dc0d06d AS runner
LABEL org.opencontainers.image.source="https://github.com/slapula/stratus-red-team/"
COPY --from=builder /build/bin/stratus /stratus
RUN apk add --update git # git is needed for Terraform to download external modules at runtime
ENTRYPOINT ["/stratus"]
CMD ["--help"]
