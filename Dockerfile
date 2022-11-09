FROM golang:alpine AS builder

RUN apk add build-base
COPY src /src
RUN cd /src && CGO_ENABLED=0 go build .


FROM scratch

WORKDIR /data/homelists
WORKDIR /app

COPY --from=builder /src/HomeLists /app/

ENTRYPOINT ["./HomeLists"]