# build stage
FROM golang:latest AS build

ENV GO111MODULE=on

# COPY ./.netrc /root/.netrc
# RUN chmod 600 /root/.netrc

WORKDIR /go/src/app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/traffic-rotator ./cmd/traffic-rotator

# final stage
FROM alpine:3.13

RUN apk --no-cache add ca-certificates
WORKDIR /usr/bin
COPY --from=build /go/src/app/bin/traffic-rotator traffic-rotator
EXPOSE 8090

ENTRYPOINT CONFIG=config/traffic-rotator/config.yaml /usr/bin/traffic-rotator
