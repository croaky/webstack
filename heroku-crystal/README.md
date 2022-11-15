# Heroku Crystal

Setup:

```bash
shards install
createdb webstack_dev
DATABASE_URL="postgres:///webstack_dev" crystal run src/api.cr
```

Go to <https://dashboard.heroku.com/apps>.
Click "New" > "Create new app".
Give it a name and submit.
Click "GitHub" and select this repo, then "Connect".
Click "Enable Automatic Deploys"

```bash
heroku login
heroku git:remote -a <app>
git remote rename heroku heroku-crystal
heroku buildpacks:add -a <app> https://github.com/lstoll/heroku-buildpack-monorepo
heroku buildpacks:add -a <app> https://github.com/crystal-lang/heroku-buildpack-crystal.git
heroku config:set APP_BASE=heroku-crystal -a <app>
heroku addons:create heroku-postgresql:hobby-dev -a <app>
```
