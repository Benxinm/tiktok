FROM golang:1.20 AS builder

ENV TZ Asiz/Shanghai
ENV CGO_EBABLED 0
ENV GOOS linux
ENV GOPROXY https://gorpoxy.cn,dircet

RUN mkdir -p /app
WORKDIR /app
ADD . /app
RUN go mod tidy
RUN make build-all


FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shanghai
ENV service api

RUN apk update
RUN apk add --no-cache ffmpeg

WORKDIR /app
COPY --from=builder /app/output /app/output
COPY --from=builder /app/config /app/config
COPY --from=builder /app/docker-entrypoint.sh /app/docker-entrypoint.sh

CMD ["sh", "./output/${service}/bootstrap.sh"]