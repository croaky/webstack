# Render Ruby Sinatra

See example deployed at
<https://webstack-ruby-sinatra.render.com/>.

## Setup

```
bundle
createdb webstack_dev
DATABASE_URL="postgres:///webstack_dev" bundle exec ruby api.rb
```

Set up Render:

Go to <https://dashboard.render.com/select-repo?type=web>.
Select this repo.
Select `main` branch.
Name the service.
Select "Ruby" environment.
Set "Build command" to `cd render-ruby-sinatra && bundle install`.
Set "Start command" to `cd render-ruby-sinatra && bundle exec ruby api.rb`.
Set "Plans" to "Free".
