class Webhook
  def url
    'https://webhook.site/02677249-4356-4357-a417-5fcf0cedf8da'
  end
end

puts "Hello World"

b = true
a = b ? 1000.0 : 500.0
puts a

puts a/3
puts (a/3)*3

wb = Webhook.new

puts wb.url

aa = Hash["a" => 100]

puts aa['a'] ? aa['a'] : 0
puts aa['b'] ? aa['b'] : 0
