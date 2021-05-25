FROM golang:1.16.4 AS builder
WORKDIR /go/src/gitlab.com/yuccastream/alertmanager-webhook-telegram-go/
COPY . .

ENV GO111MODULE="off"
RUN set -xe \
    && go build -o awt-go main.go

FROM debian:10-slim
LABEL maintainer="Yucca Stream https://yucca.app"
RUN apt-get update \
    && apt-get install -y ca-certificates \
    && apt-get autoremove -y \
    && apt-get autoclean \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

COPY --from=builder /go/src/gitlab.com/yuccastream/alertmanager-webhook-telegram-go/awt /usr/local/bin/awt

EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/awt"]
