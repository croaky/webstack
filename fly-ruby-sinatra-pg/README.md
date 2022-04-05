# Fly.io Ruby Sinatra Postgres

Develop app locally:

```
bundle
createdb ruby_fly_dev
DATABASE_URL=postgres://postgres:postgres@localhost:5432/ruby_fly_dev rackup
```

Set up Fly.io:

```
brew install flyctl
flyctl auth login
flyctl launch --image heroku/buildpacks:18
```
