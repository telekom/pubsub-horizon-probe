# Copyright 2024 Deutsche Telekom IT GmbH
#
# SPDX-License-Identifier: Apache-2.0

FROM golang:1.22-alpine AS build

ARG HTTP_PROXY
ARG HTTPS_PROXY

ENV HTTP_PROXY=$HTTP_PROXY
ENV HTTPS_PROXY=$HTTPS_PROXY

WORKDIR /build
COPY . .
RUN apk add --no-cache build-base
RUN apk add --no-cache --update ca-certificates
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s -extldflags=-static" -o ./out/probe

FROM scratch

ENV PROBE_ARG_MESSAGE_COUNT=3
ENV PROBE_ARG_TIMEOUT=30s
ENV PROBE_ARG_MAX_LATENCY=5s
ENV PROBE_ARG_TEMPLATE=template.json

COPY --from=build /build/out/probe .
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["./probe"]
CMD ["start",\
     "--message-count=$PROBE_ARG_MESSAGE_COUNT",\
     "--timeout=$PROBE_ARG_TIMEOUT",\
     "--max-latency=$PROBE_ARG_MAX_LATENCY",\
     "--template=$PROBE_ARG_TEMPLATE"\
 ]