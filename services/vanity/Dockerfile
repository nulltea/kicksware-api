FROM golang:alpine AS builder
LABEL Name=vanity-urlhttps://hub.docker.com/orgs
# Helps using module proxy for caching deps for faster build
ARG GO111MODULE=off
ARG GOPROXY
RUN apk add git
RUN go get -d github.com/GoogleCloudPlatform/govanityurls

WORKDIR $GOPATH/src/github.com/GoogleCloudPlatform/govanityurls
RUN go get -v .

FROM alpine:3.9.2
COPY --from=builder /go/bin/govanityurls /go/bin/
WORKDIR /app/
EXPOSE 8080
ENTRYPOINT [ "/go/bin/govanityurls" ]
