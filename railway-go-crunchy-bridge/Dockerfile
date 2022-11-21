# Build app
FROM golang:1.19-alpine as builder
WORKDIR /app
COPY . ./
RUN go build -o /app/server

# Build Tailscale
ENV TSFILE=tailscale_1.32.2_amd64.tgz
RUN wget https://pkgs.tailscale.com/stable/${TSFILE} && tar xzf ${TSFILE} --strip-components=1

# Second of multi-stage Docker build
FROM alpine:latest

# Tailscaled can't run as root user, cuz of an SO_MARK issue.
# See: https://github.com/tailscale/tailscale/issues/634.
ARG USER=default
ENV HOME /home/$USER
RUN adduser -D $USER \
  && echo "$USER ALL=(ALL) NOPASSWD: ALL" > /etc/sudoers.d/$USER \
  && chmod 0440 /etc/sudoers.d/$USER
USER $USER
WORKDIR $HOME

# Copy binaries to production image
COPY --from=builder /app/start.sh /app/start.sh
COPY --from=builder /app/server /app/server
COPY --from=builder /app/tailscaled /app/tailscaled
COPY --from=builder /app/tailscale /app/tailscale

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN mkdir -p /var/run/tailscale /var/cache/tailscale /var/lib/tailscale
RUN mkdir -p $HOME/tailscale-storage

# Run on container startup
CMD ["/app/start.sh"]
