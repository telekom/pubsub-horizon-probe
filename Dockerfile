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
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s -extldflags=-static" -o ./out/probe

FROM scratch
COPY --from=build /build/out/probe probe
ENTRYPOINT ["./probe", "start", "--template", "template.json"]