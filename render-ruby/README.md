# Render Ruby

See example deployed at
<https://ruby-sinatra.onrender.com>.

## Setup

```
bundle
createdb webstack_dev
DATABASE_URL="postgres:///webstack_dev" bundle exec ruby api.rb
```

Go to <https://dashboard.render.com/new/database>
Name the database "webstack".
Set "Plans" to "Free".
Copy "Internal Connection String".

Go to <https://dashboard.render.com/select-repo?type=web>.
Select this repo.
Select `main` branch.
Name the service.
Select "Ruby" environment.
Set "Build command" to `cd render-ruby && bundle install`.
Set "Start command" to `cd render-ruby && bundle exec ruby api.rb`.
Set "Plans" to "Free".
Click "Advanced".
Click "Add Environment Variable".
Paste "Internal Connection String" value from above into new `DATABASE_URL` env var.

## Editorial

* The [Render docs](https://render.com/docs/deploy-rails#additional-notes) say
  "By default, Render uses the latest LTS version of Ruby" but using
  `ruby "~> 3"` in the `Gemfile` resulted in build errors:
  "Your Ruby version is 2.6.8, but your Gemfile specified ~> 3".
* Databases are backed up daily, not continuously.
* HA databases are not available.
