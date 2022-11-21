#!/bin/sh

HOST="railway-app-${RAILWAY_GIT_COMMIT_SHA:0:8}"
SOCKET="$HOME/tailscale-storage/tailscale.sock"

# Connect machine to Tailscale tailnet
/app/tailscaled --state=mem: --tun=userspace-networking --socks5-server="localhost:1055" --socket="$SOCKET" &

until /app/tailscale --socket="$SOCKET" up --authkey="$TAILSCALE_AUTHKEY" --hostname="$HOST"
do
  echo "Waiting for Tailscale auth"
  sleep 5
done

# Run server with 'exec' to forward SIGINT & SIGTERM to program for cleanup.
exec ALL_PROXY="socks5://localhost:1055/" /app/server
