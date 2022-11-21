#!/bin/sh

# Connect machine to Tailscale tailnet
/app/tailscaled --state=mem: --tun=userspace-networking --socks5-server="localhost:1055" &
ID="${RAILWAY_GIT_COMMIT_SHA:0:8}"

until /app/tailscale up --authkey="${TAILSCALE_AUTHKEY}" --hostname="railway-app-${ID}"
do
  echo "Waiting for Tailscale auth..."
  sleep 5
done

# Run server
ALL_PROXY="socks5://localhost:1055/" /app/server
