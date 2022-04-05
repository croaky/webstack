# Fly Ruby Sinatra

## Setup

Develop locally:

```
bundle
createdb webstack_dev
DATABASE_URL="postgres:///webstack_dev" bundle exec ruby api.rb
```

Set up Fly.io:

```
brew install flyctl
flyctl auth login
flyctl launch --image heroku/buildpacks:20
```
