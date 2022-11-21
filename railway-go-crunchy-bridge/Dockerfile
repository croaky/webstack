# Build app
FROM golang:1.19-alpine as builder
WORKDIR /app
COPY . ./
RUN go build -o /app/server

# Build Tailscale
RUN curl -fsSL https://tailscale.com/install.sh | sh

FROM alpine:latest
RUN apk update && apk add --update sudo ca-certificates && rm -rf /var/cache/apk/*

# Copy binaries to production image
COPY --from=builder /app/start.sh /app/start.sh
COPY --from=builder /app/server /app/server
COPY --from=builder /app/tailscaled /app/tailscaled
COPY --from=builder /app/tailscale /app/tailscale
RUN mkdir -p /var/run/tailscale /var/cache/tailscale /var/lib/tailscale

# Run on container startup
CMD ["/app/start.sh"]
