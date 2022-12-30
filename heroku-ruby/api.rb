require "connection_pool"
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

class DB
  def initialize
    connect
  end

  def exec(sql)
    do_exec(sql)
  rescue PG::ConnectionBad
    connect
    do_exec(sql)
  end

  private

  def connect
    @pool = ConnectionPool.new(size: 5, timeout: 5) {
      PG.connect(ENV.fetch("DATABASE_URL"))
    }
  end

  def do_exec(sql)
    @pool.with do |conn|
      conn.exec(sql)
    end
  end
end

db = DB.new

before do
  headers(
    "Accept" => "application/json",
    "Access-Control-Allow-Methods" => "GET",
    "Content-Type" => "application/json",
    "X-Content-Type-Options" => "nosniff"
  )
end

get "/" do
  [
    Thread.new { db.exec "SELECT 1" },
    Thread.new { db.exec "SELECT 1" },
    Thread.new { db.exec "SELECT 1" },
    Thread.new { db.exec "SELECT 1" },
    Thread.new { db.exec "SELECT 1" },
    Thread.new { db.exec "SELECT 1" },
    Thread.new { db.exec "SELECT 1" },
    Thread.new { db.exec "SELECT 1" },
    Thread.new { db.exec "SELECT 1" }
  ].each(&:join)
  content_type :json
  {status: "ok"}.to_json
end
