#!/bin/sh

# Connect machine to Tailscale tailnet
/app/tailscaled --tun=userspace-networking --socks5-server=localhost:1055 &
ID="${RAILWAY_GIT_COMMIT_SHA:0:8}"
/app/tailscale up --authkey="${TAILSCALE_AUTHKEY}" --hostname="railway-app-${ID}"
/app/tailscale status

# Run server
/app/server
