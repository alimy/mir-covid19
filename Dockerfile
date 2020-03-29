FROM golang:alpine AS binaryBuilder
# Install build deps
RUN apk --no-cache --no-progress add --virtual build-deps build-base git
WORKDIR /go/src/github.com/alimy/mir-covid19
COPY . .
RUN export GO111MODULE=on && make build

FROM alpine:latest
# Install system utils & Mir-Music runtime dependencies
ADD https://github.com/tianon/gosu/releases/download/1.11/gosu-amd64 /usr/sbin/gosu
RUN chmod +x /usr/sbin/gosu \
  && echo http://dl-2.alpinelinux.org/alpine/edge/community/ >> /etc/apk/repositories \
  && apk --no-cache --no-progress add \
    bash \
    shadow \
    s6

ENV COVID19_CUSTOM /data/covid19

# Configure LibC Name Service
COPY assets/docker/nsswitch.conf /etc/nsswitch.conf

WORKDIR /app/covid19
COPY assets/docker ./docker
COPY --from=binaryBuilder /go/src/github.com/alimy/mir-covid19/covid .

RUN ./docker/finalize.sh

# Configure Docker Container
VOLUME ["/data"]
EXPOSE 10169
ENTRYPOINT ["/app/covid19/docker/start.sh"]
CMD ["/bin/s6-svscan", "/app/covid19/docker/s6/"]
