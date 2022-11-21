# Build app
FROM golang:1.19-alpine as builder
WORKDIR /app
COPY . ./
RUN go build -o /app/server

# Build Tailscale
FROM alpine:latest
RUN curl -fsSL https://tailscale.com/install.sh | sh
RUN apk update && apk add --update sudo ca-certificates && rm -rf /var/cache/apk/*

# Copy binaries to production image
COPY --from=builder /app/start.sh /app/start.sh
COPY --from=builder /app/server /app/server

# Run on container startup
CMD ["/app/start.sh"]
