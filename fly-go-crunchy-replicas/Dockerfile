# Build app
FROM golang:1.19-alpine as builder
WORKDIR /app
COPY . ./
RUN go build -o /app/server

# Build Tailscale
ENV TSFILE=tailscale_1.32.2_amd64.tgz
RUN wget https://pkgs.tailscale.com/stable/${TSFILE} && tar xzf ${TSFILE} --strip-components=1

# Copy binaries to production image
FROM alpine:latest
COPY --from=builder /app/start.sh /app/start.sh
COPY --from=builder /app/server /app/server
COPY --from=builder /app/tailscaled /app/tailscaled
COPY --from=builder /app/tailscale /app/tailscale
RUN apk update && apk add ca-certificates iptables ip6tables && rm -rf /var/cache/apk/*
RUN mkdir -p /var/run/tailscale /var/cache/tailscale /var/lib/tailscale

# Run on container startup
EXPOSE 8080
CMD ["/app/start.sh"]
