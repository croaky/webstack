#!/bin/sh

HOST="railway-app-${RAILWAY_GIT_COMMIT_SHA:0:8}"

# Connect machine to Tailscale tailnet
/app/tailscaled --tun=userspace-networking --socks5-server="localhost:1055" &

until /app/tailscale up --authkey="$TAILSCALE_AUTHKEY" --hostname="$HOST"
do
  sleep 0.1
done

# Run app
ALL_PROXY="socks5://localhost:1055/" /app/server
