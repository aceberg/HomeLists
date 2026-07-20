FROM golang:alpine AS builder

RUN apk add build-base
COPY . /src
RUN cd /src/cmd/HomeLists/ && CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /HomeLists .


FROM scratch

WORKDIR /data/homelists
WORKDIR /app

COPY --from=builder /HomeLists /app/

ENTRYPOINT ["./HomeLists"]