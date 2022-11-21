#!/bin/sh

HOST="railway-app-${RAILWAY_GIT_COMMIT_SHA:0:8}"

# Connect machine to Tailscale tailnet
/app/tailscaled --tun=userspace-networking --socks5-server="localhost:1055" &
/app/tailscale up --authkey="$TAILSCALE_AUTHKEY" --hostname="$HOST"

# Run app
ALL_PROXY="socks5://localhost:1055/" /app/server
