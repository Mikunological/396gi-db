FROM golang:alpine as build-reflex
RUN apk update && \
    apk upgrade && \
    apk add bash git && \
    rm -rf /var/cache/apk/*

RUN go get github.com/cespare/reflex

FROM golang:alpine
COPY --from=build-reflex /go/bin/reflex /go/bin/reflex

RUN apk update && \
    apk upgrade && \
    apk add bash git build-base jq && \
    rm -rf /var/cache/apk/*

WORKDIR /app/backend/go-app
