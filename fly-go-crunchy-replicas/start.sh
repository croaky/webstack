#!/bin/sh

# Connect machine to Tailscale tailnet.
/app/tailscaled --state=/var/lib/tailscale/tailscaled.state --socket=/var/run/tailscale/tailscaled.sock &
/app/tailscale up --authkey="${TAILSCALE_AUTHKEY}" --hostname="fly-${FLY_REGION}"

# Auto-remove immediately when machine shuts down.
# https://tailscale.com/blog/ephemeral-logout/
# Instead of a 30m-45h window.
# https://tailscale.com/kb/1111/ephemeral-nodes/
/app/tailscale logout --authkey="${TAILSCALE_AUTHKEY}"

# Run server
/app/server
