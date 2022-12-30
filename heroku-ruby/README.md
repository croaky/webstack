# Heroku Ruby

Setup:

```bash
bundle
createdb webstack_dev
DATABASE_URL="postgres:///webstack_dev" bundle exec ruby api.rb
```

Go to <https://dashboard.heroku.com/apps>.
Click "New" > "Create new app".
Give it a name and submit.
Click "GitHub" and select this repo, then "Connect".
Click "Enable Automatic Deploys"

```bash
heroku login
heroku git:remote -a <app>
git remote rename heroku heroku-ruby
heroku buildpacks:add -a <app> https://github.com/lstoll/heroku-buildpack-monorepo
heroku buildpacks:add -a <app> heroku/ruby
heroku config:set APP_BASE=heroku-ruby -a <app>
heroku config:set APP_ENV=prod -a <app>
heroku addons:create heroku-postgresql:hobby-dev -a <app>
```
