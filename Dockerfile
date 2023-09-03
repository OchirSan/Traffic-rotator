# build stage
FROM golang:latest AS build

ENV GO111MODULE=on

WORKDIR /go/src/app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/traffic-rotator ./cmd/traffic-rotator

# final stage
FROM alpine:3.11
RUN apk --no-cache add ca-certificates tzdata
WORKDIR /usr/bin
COPY --from=build /go/src/app /go

EXPOSE 80
ENTRYPOINT /go/bin/traffic-rotator

EXPOSE 80