# Build app
FROM golang:1.20-alpine as builder
WORKDIR /app
COPY . ./
RUN go build -o /app/server

# Build Tailscale
ENV TSFILE=tailscale_1.40.0_amd64.tgz
RUN wget https://pkgs.tailscale.com/stable/${TSFILE} && tar xzf ${TSFILE} --strip-components=1

# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM alpine:latest
RUN apk update && apk add ca-certificates iptables ip6tables && rm -rf /var/cache/apk/*

# Copy binary to production image
COPY --from=builder /app/start.sh /app/start.sh
COPY --from=builder /app/server /app/server
COPY --from=builder /app/tailscaled /app/tailscaled
COPY --from=builder /app/tailscale /app/tailscale
RUN mkdir -p /var/run/tailscale /var/cache/tailscale /var/lib/tailscale

# Run on container startup
EXPOSE 8080
CMD ["/app/start.sh"]
