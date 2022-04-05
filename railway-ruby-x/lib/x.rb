require "async"
require "async/http/endpoint"
require "async/http/server"
require "connection_pool"
require "json"
require "pg"

module X
  class Env
    def self.load(root_dir:)
      # fallbacks
      ENV["APP_ENV"] ||= "dev"
      ENV["DATABASE_POOL_SIZE"] ||= "5"
      ENV["PORT"] ||= "2000"

      path = "#{root_dir}/env/#{ENV.fetch("APP_ENV")}"

      if File.exist?(path)
        File.readlines(path).each do |line|
          key, val = line.split("=", 2)
          ENV[key] = val.chomp
        end
      end
    end
  end

  class API
    def self.serve(&block)
      db = X::Database.new(
        ENV.fetch("DATABASE_URL"),
        ENV.fetch("DATABASE_POOL_SIZE")
      )
      url = "http://0.0.0.0:#{ENV.fetch("PORT")}"
      router = X::Router.new
      router.instance_exec(db, &block)
      endpoint = Async::HTTP::Endpoint.parse(url)
      server = Async::HTTP::Server.new(router, endpoint)

      Async do |task|
        task.async do
          puts "Listening on #{url}"
          server.run
        end
      end
    end
  end

  class Database
    def initialize(url, pool_size)
      @pool = ConnectionPool.new(size: pool_size, timeout: 5) {
        conn = PG.connect(url)
        conn.type_map_for_results = PG::BasicTypeMapForResults.new(conn)
        conn
      }
    end

    def exec(sql, params = nil)
      @pool.with do |conn|
        if params.nil?
          conn.exec(sql)
        else
          conn.exec_params(sql, params)
        end
      end
    end

    def self.migrate
      db_url = ENV.fetch("DATABASE_URL")
      db = new(db_url, ENV.fetch("DATABASE_POOL_SIZE"))
      ready = db.exec(<<~SQL).first
        SELECT EXISTS (
          SELECT FROM information_schema.tables
          WHERE table_schema = 'public' AND table_name = 'migrations'
        )
      SQL
      if !ready
        puts "The migrations table does not exist. Try: x db init"
        exit 1
      end
      migrations = db.exec(<<~SQL).to_a
        SELECT
          version
        FROM
          migrations
        ORDER BY
          version DESC
      SQL
      Dir.glob("db/migrate/**.sql").each do |mig|
        version = mig.split("/")[-1].split(".")[0]
        if migrations.include?(version)
          next
        end
        sql = File.read(mig)
        db.exec(sql)
        db.exec(
          "INSERT INTO migrations (version) VALUES ($1)",
          [version]
        )
      end
      if ENV.fetch("APP_ENV") == "dev"
        db_name = db_url.split("/")[-1].split("?")[0]
        `pg_dump --schema-only #{db_name} > db/schema.sql`
      end
    end
  end

  class Request
    attr_accessor :headers, :body

    def initialize(headers, body)
      @headers = {}
      headers.each do |header|
        key, val = header.flatten.to_s.split(" ", 2)
        @headers[key] = val
      end
      @body = body || {}
    end

    def blank?(*keys)
      keys.any? { |k| @body[k].nil? || @body[k] == "" }
    end
  end

  class Response
    attr_accessor :status, :headers, :body

    def initialize
      @headers = {
        "Accept" => "application/json",
        "Access-Control-Allow-Methods" => "OPTIONS,GET,POST",
        "Content-Type" => "application/json"
      }
    end
  end

  class Router
    def initialize
      @routes = {}
    end

    def get(path, &handler)
      @routes[path] = handler
    end

    def post(path, &handler)
      @routes[path] = handler
    end

    def halt(status, msg, hint: nil)
      throw :halt, [status, {err: msg, hint: hint}.compact.to_json]
    end

    def call(req)
      t = Process.clock_gettime(Process::CLOCK_MONOTONIC)
      resp = X::Response.new

      if !["OPTIONS", "GET", "POST"].include?(req.method)
        resp.status = 405
        resp.body = %({"err": "#{req.method} method not accepted"})
      end

      if req.method == "OPTIONS"
        # Cache pre-flight requests for 2 hours
        # https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Max-Age
        resp.headers = resp.headers.merge("Access-Control-Max-Age" => "7200")
        resp.status = 204
      end

      if req.method == "GET" || req.method == "POST"
        handler = @routes[req.path]

        if handler.nil?
          resp.status = 404
          resp.body = %({"err": "HTTP method not supported. Try OPTIONS, GET, or POST."})
        end
      end

      if req.method == "POST"
        if resp.status.nil? && req.body && req.body.length > 0
          begin
            parsed_body = JSON.parse(req.body.read)
          rescue
            resp.status = 400
            resp.body = %({err: "JSON request body is invalid."})
          end
        end
      end

      if resp.status.nil?
        resp.status, resp.headers, resp.body = catch(:halt) {
          throw :halt, handler.call(X::Request.new(req.headers, parsed_body), resp)
        }
      end

      if resp.status.nil?
        resp.status = 500
      end

      elapsed = (Process.clock_gettime(Process::CLOCK_MONOTONIC) - t).round(3)
      puts "#{elapsed}s #{resp.status} #{req.method} #{req.path}"
      Protocol::HTTP::Response[resp.status, resp.headers, [resp.body]]
    rescue => err
      if err.is_a?(ConnectionPool::TimeoutError)
        resp.status = 504
        resp.body = %({err: "timeout"})
      else
        resp.status = 500
        resp.body = %({err: "something went wrong"})
      end

      elapsed = (Process.clock_gettime(Process::CLOCK_MONOTONIC) - t).round(3)
      puts "#{elapsed}s #{resp.status} #{req.method} #{req.path} #{err.message}"
      Protocol::HTTP::Response[resp.status, resp.headers, [resp.body]]
    end
  end
end
