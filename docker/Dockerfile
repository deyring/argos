##
## Build
##
FROM golang:1.18-bullseye AS build

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o argos -v github.com/deyring/argos/cmd

##
## Deploy
##
FROM alpine:3.16

WORKDIR /

COPY --from=build /app/argos .

USER nonroot:nonroot

ENTRYPOINT ["/argos"]