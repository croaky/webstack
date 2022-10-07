# Vercel Edge Functions w/ PlanetScale serverless driver for JavaScript

Set up PlanetScale:

* Create a [PlanetScale](https://planetscale.com) account.
* Create a database.
* Go to "Settings" > "Beta features".
  Click "Enroll" for "PlanetScale serverless driver for JavaScript".
* Go to "Settings" > Passwords".
  Click "New password".
* Compose a `DATABASE_URL` from the newly created password.

I chose to create one environment variable because the driver
will decompose the username, password, and host:

```ts
if (config.url) {
  const url = new URL(config.url)
  this.config.username = url.username
  this.config.password = url.password
  this.config.host = url.hostname
}
```

The host looks like `aws.connect.psdb.cloud`
instead of a region like `us-east-1`.
I'm guessing that's the [global router](https://planetscale.com/blog/introducing-the-planetscale-serverless-driver-for-javascript)?

> Similar to a CDN, global routing reduces latency drastically in situations
> where a client is connecting from a geographically distant location, which is
> common within serverless and edge compute environments. A client connects to
> the closest geographic edge in our network, then backhauls over long-held
> connection pools over our internal network to reach the actual destination.
> Within the US, connections from US West to a database in US East reduce
> latency by 100ms or more in some cases, even more as the distance increases.

Set up Vercel:

* Create a [Vercel](https://vercel.com) account.
* Add a `DATABASE_URL` environment variable from PlanetScale.

Deploy:

```bash
npm i -g vercel
npm i
vercel deploy
vercel --prod
```
