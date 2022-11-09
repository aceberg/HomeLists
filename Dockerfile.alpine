FROM golang:alpine AS builder

RUN apk add build-base
COPY src /src
RUN cd /src && CGO_ENABLED=0 go build .


FROM alpine

WORKDIR /app

RUN apk add --no-cache tzdata \
    && mkdir -p /data/homelists

COPY --from=builder /src/HomeLists /app/

ENTRYPOINT ["./HomeLists"]