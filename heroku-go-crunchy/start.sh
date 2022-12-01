#!/bin/sh

# Connect machine to Tailscale tailnet and auto-remove immediately when machine shuts down.
# https://tailscale.com/blog/ephemeral-logout/
/app/tailscaled --state=mem: --tun=userspace-networking --socks5-server=localhost:1055 &
ID=$(echo "${HEROKU_DYNO_ID}" | cut -d '-' -f 1)
/app/tailscale up --authkey="${TAILSCALE_AUTHKEY}" --hostname="heroku-${ID}"
echo "Tailscale started"

# Run server
ALL_PROXY=socks5://localhost:1055/ /app/server
