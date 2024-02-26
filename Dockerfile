FROM golang:1.21 as build

ENV OOKLA_SPEEDTEST_VERSION=1.2.0

RUN mkdir /speedtest
WORKDIR /speedtest

RUN wget -O speedtest.tgz https://install.speedtest.net/app/cli/ookla-speedtest-${OOKLA_SPEEDTEST_VERSION}-linux-x86_64.tgz && \
  tar -zxvf speedtest.tgz && \
  ls -alh && pwd

COPY ./ /app

WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o ./speedtest-exporter

FROM alpine as ca

RUN apk update && apk add --no-cache ca-certificates

FROM scratch

COPY --from=build /app/speedtest-exporter /speedtest-exporter
COPY --from=build /speedtest/speedtest /speedtest
COPY --from=ca /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

CMD [ "/speedtest-exporter" ]