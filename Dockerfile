FROM golang:alpine AS binaryBuilder
# Install build deps
RUN apk --no-cache --no-progress add --virtual build-deps build-base git
WORKDIR /go/src/github.com/alimy/echo-music
COPY . .
RUN export GO111MODULE=on && make build

FROM alpine:latest
# Install system utils & Gin-Music runtime dependencies
ADD https://github.com/tianon/gosu/releases/download/1.11/gosu-amd64 /usr/sbin/gosu
RUN chmod +x /usr/sbin/gosu \
  && echo http://dl-2.alpinelinux.org/alpine/edge/community/ >> /etc/apk/repositories \
  && apk --no-cache --no-progress add \
    bash \
    shadow \
    s6

ENV ECHOMUSIC_CUSTOM /data/echomusic

# Configure LibC Name Service
COPY hack/docker/nsswitch.conf /etc/nsswitch.conf

WORKDIR /app/echomusic
COPY hack/docker ./docker
COPY --from=binaryBuilder /go/src/github.com/alimy/echo-music/echo-music .

RUN ./docker/finalize.sh

# Configure Docker Container
VOLUME ["/data"]
EXPOSE 8013
ENTRYPOINT ["/app/echomusic/docker/start.sh"]
CMD ["/bin/s6-svscan", "/app/ginmusic/docker/s6/"]