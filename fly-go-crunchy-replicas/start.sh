#!/bin/sh

# Connect machine to Tailscale tailnet and auto-remove immediately when machine shuts down.
# https://tailscale.com/blog/ephemeral-logout/
/app/tailscaled --state=/var/lib/tailscale/tailscaled.state --socket=/var/run/tailscale/tailscaled.sock &
ID=$(echo "${FLY_ALLOC_ID}" | cut -d '-' -f 1)
/app/tailscale up --authkey="${TAILSCALE_AUTHKEY}" --hostname="fly-${FLY_REGION}-${ID}"
echo "Tailscale started"

# Run server
/app/server
