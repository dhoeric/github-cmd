FROM golang:1.9.3-alpine3.6 AS build-env

WORKDIR /go/src/github.com/dhoeric/github-cmd

RUN apk add --no-cache git

COPY . .
RUN go get -v -d && \
	go build -o app

FROM alpine:3.6
COPY --from=build-env /go/src/github.com/dhoeric/github-cmd/app .
ENTRYPOINT ["/app"]
