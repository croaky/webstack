#!/bin/sh

# Connect machine to Tailscale tailnet and auto-remove immediately when machine shuts down.
# https://tailscale.com/blog/ephemeral-logout/
/app/tailscaled --state=mem: --tun=userspace-networking --socks5-server=localhost:1055 &
/app/tailscale up --authkey="${TAILSCALE_AUTHKEY}" --hostname="railway-app"
echo "Tailscale started"

# Run server
ALL_PROXY=socks5://localhost:1055/ /app/server
