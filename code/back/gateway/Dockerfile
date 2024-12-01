FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN apk add --no-cache make
RUN go build -ldflags "-s -w" -o ./bin/gateway ./cmd

FROM alpine:latest AS runner
WORKDIR /app
COPY --from=builder /app/bin/gateway .
CMD ["./gateway"]
