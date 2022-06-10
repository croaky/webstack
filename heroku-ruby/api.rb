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
    @pool.with do |conn|
      conn.exec(sql)
    end
  rescue PG::ConnectionBad
    connect
    @pool.with do |conn|
      conn.exec(sql)
    end
  end

  private def connect
    if @pool.nil?
      @pool = ConnectionPool.new(size: 5, timeout: 5) {
        db = PG.connect(ENV.fetch("DATABASE_URL"))
        db.type_map_for_results = PG::BasicTypeMapForResults.new(db)
        db
      }
    else
      @pool.reload { |conn| conn&.close }
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
  db.exec "SELECT 1"
  content_type :json
  {status: "ok"}.to_json
end
