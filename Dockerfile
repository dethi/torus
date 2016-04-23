FROM gliderlabs/alpine
MAINTAINER Thibault Deutsch <thibault.deutsch@gmail.com>

EXPOSE 80
EXPOSE 50007

VOLUME /static
WORKDIR /app

RUN apk-install ca-certificates

COPY torrent_service /app/
ENTRYPOINT ["./torrent_service"]
