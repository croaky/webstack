require "pg"
require "sinatra"

ENV["APP_ENV"] ||= "dev"

path = "#{__dir__}/env/#{ENV.fetch("APP_ENV")}"

if File.exist?(path)
  File.readlines(path).each do |line|
    key, val = line.split("=", 2)
    ENV[key] = val.chomp
  end
end

db = PG.connect(ENV.fetch("DATABASE_URL"))

get "/" do
  db.exec "SELECT 1"
  content_type :json
  {status: "ok"}.to_json
end
