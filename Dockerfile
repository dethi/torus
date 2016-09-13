FROM gliderlabs/alpine
MAINTAINER Thibault Deutsch <thibault.deutsch@gmail.com>

EXPOSE 80
EXPOSE 50007

VOLUME /static
WORKDIR /app

RUN apk-install ca-certificates sqlite-libs
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

COPY torus /app/
ENTRYPOINT ["./torus"]
