# Build app
FROM golang:1.20-alpine as builder
WORKDIR /app
COPY . ./
RUN go build -o /app/server

# Build Tailscale
ENV TSFILE=tailscale_1.42.0_amd64.tgz
RUN wget https://pkgs.tailscale.com/stable/${TSFILE} && tar xzf ${TSFILE} --strip-components=1

# Copy binaries to production image
FROM alpine:latest
COPY --from=builder /app/start.sh /app/start.sh
COPY --from=builder /app/tailscaled /app/tailscaled
COPY --from=builder /app/tailscale /app/tailscale
COPY --from=builder /app/server /app/server
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN mkdir -p /var/run/tailscale /var/cache/tailscale /var/lib/tailscale

# Run on container startup
CMD ["/app/start.sh"]
