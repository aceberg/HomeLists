FROM golang:alpine AS builder

RUN apk add build-base
COPY src /src
RUN cd /src && go build .


FROM alpine

WORKDIR /app

RUN apk add --no-cache tzdata \
    && mkdir -p /data/homelists

COPY src/templates /app/templates
COPY --from=builder /src/HomeLists /app/

ENTRYPOINT ["./HomeLists"]