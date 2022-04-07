# Heroku Crystal Kemal

See example deployed at
<https://webstack-crystal-kemal.herokuapp.com/>.

## Setup

```
shards install
createdb webstack_dev
DATABASE_URL="postgres:///webstack_dev" crystal run api.cr
```

Go to <https://dashboard.heroku.com/apps>.
Click "New" > "Create new app".
Give it a name and submit.
Click "GitHub" and select this repo, then "Connect".
Click "Enable Automatic Deploys"

```
heroku login
heroku git:remote -a <app>
git remote rename heroku heroku-crystal-kemal
heroku buildpacks:add -a <app> https://github.com/lstoll/heroku-buildpack-monorepo
heroku buildpacks:add -a <app> https://github.com/crystal-lang/heroku-buildpack-crystal.git
heroku config:set APP_BASE=heroku-crystal-kemal -a <app>
heroku addons:create heroku-postgresql:hobby-dev -a <app>
```
