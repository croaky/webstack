# Fly Ruby

See example deployed at
<https://webstack-ruby-sinatra.fly.dev/>.

## Setup

Develop:

```
bundle
createdb webstack_dev
DATABASE_URL="postgres:///webstack_dev" bundle exec ruby api.rb
```

Install Docker from <https://www.docker.com/get-started/>
(~1.5GB for M1 chip).

Set up Fly:

```
brew install flyctl
flyctl auth login
flyctl launch
```

# Editorial

* With Docker and `flyctl` installed ahead of time,
  the deploy, including a Postgres database linked to the web app,
  worked flawlessly. It took a few minutes for the image to build and
  resulted in a 634 MB image, which feels large for a small Sinatra app?
* It seems like Fly can understand a `Procfile`, which is helpful for not
  needing to brush up on `Dockerfile` layouts, which I'm rusty on.
