#!/bin/sh

# Connect machine to Tailscale tailnet
/app/tailscaled --state=mem: --tun=userspace-networking --socks5-server="localhost:1055"  --outbound-http-proxy-listen="localhost:1055" &
/app/tailscale up --authkey="$TAILSCALE_AUTHKEY" --hostname="railway-app-${RAILWAY_GIT_COMMIT_SHA:0:8}"

# Run app
ALL_PROXY="socks5://localhost:1055/" /app/server
