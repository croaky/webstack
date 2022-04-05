require "rubygems"
require "bundler"

Bundler.setup(:default)

require "pg"
require "sinatra"

db = PG.connect(ENV.fetch("DATABASE_URL"))

get "/" do
  db.exec "SELECT 1"
  content_type "application/json"
  {status: "ok"}.to_json
end
