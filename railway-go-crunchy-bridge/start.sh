#!/bin/sh

# Connect machine to Tailscale tailnet
/app/tailscale up --authkey="$TAILSCALE_AUTHKEY" --hostname="railway-app-${RAILWAY_GIT_COMMIT_SHA:0:8}"
/app/tailscale logout

# Run app
/app/server
