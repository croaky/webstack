require "db"
require "kemal"
require "json"
require "pg"

DB.open(ENV["DATABASE_URL"]) do |db|
  get "/" do |env|
    db.exec "SELECT 1"
    env.response.content_type = "application/json"
    {"status": "ok"}.to_json
  end

  Kemal.run
end
