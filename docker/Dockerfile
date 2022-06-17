##
## Build
##
FROM golang:1.18-alpine as builder

WORKDIR /app

ADD . .

RUN go mod download
RUN apk --no-cache add git

RUN go build -o argos -v github.com/deyring/argos/cmd

##
## Deploy
##
FROM alpine:3.16
RUN apk add --no-cache ca-certificates && \
    adduser -D -u 12345 -g 12345 argos

COPY --from=builder ./app/argos /usr/bin/argos

USER 12345
WORKDIR /home/argos

ENTRYPOINT [ "argos" ]