FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN apk add --no-cache make
RUN go build -ldflags "-s -w" -o ./bin/gateway ./code/back/gateway/cmd

FROM alpine:latest AS runner
WORKDIR /app
CMD ["./bin/gateway"]
