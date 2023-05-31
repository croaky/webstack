FROM golang:1.20-alpine as builder
WORKDIR /app
COPY . ./
RUN go build -o /app/server

FROM alpine:latest as tailscale
WORKDIR /app
ENV TSFILE=tailscale_1.42.0_amd64.tgz
RUN wget https://pkgs.tailscale.com/stable/${TSFILE} && tar xzf ${TSFILE} --strip-components=1

FROM alpine:latest
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

# Copy binaries to production image
COPY --from=builder /app/start.sh /app/start.sh
COPY --from=builder /app/server /app/server
COPY --from=tailscale /app/tailscaled /app/tailscaled
COPY --from=tailscale /app/tailscale /app/tailscale
RUN mkdir -p /var/run/tailscale /var/cache/tailscale /var/lib/tailscale

# Run on container startup
CMD ["/app/start.sh"]
