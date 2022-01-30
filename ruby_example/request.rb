require 'uri'
require 'net/http'

uri = URI('https://upb-bengawansolo.id/api/bendungan/volume')
# params = { :api_key => 'your_api_key' }
# uri.query = URI.encode_www_form()

res = Net::HTTP.get_response(uri)
puts res.body if res.is_a?(Net::HTTPSuccess)
