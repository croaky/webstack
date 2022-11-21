#!/bin/sh

# Connect machine to Tailscale tailnet
/app/tailscaled --state=mem: --socket=/var/run/tailscale/tailscaled.sock &
ID="${RAILWAY_GIT_COMMIT_SHA:0:8}"
/app/tailscale up --authkey="${TAILSCALE_AUTHKEY}" --hostname="railway-app-${ID}"

# Run server
/app/server
