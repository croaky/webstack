ENV["APP_ENV"] = "test"

require "rspec"
require "rack/test"

# # http://rubydoc.info/gems/rspec-core/RSpec/Core/Configuration
RSpec.configure do |config|
  config.include Rack::Test::Methods

  def app
    Sinatra::Application
  end
end
