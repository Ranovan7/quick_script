
class Webhook
  def url
    'https://webhook.site/02677249-4356-4357-a417-5fcf0cedf8da'
  end
end

def testing(test)
  puts "Test : #{test}"
  return unless test
  puts test
end

def generate(num)
  case num
  when 10
    "test"
  end
end

puts "Hello World"

b = true
a = b ? 1000.0 : 500.0
puts a

testing = "testing "

puts a/3
puts (a/3)*3

wb = Webhook.new

puts wb.url

aa = Hash["a" => 100, 1 => 1000]

puts aa['a'] ? aa['a'] : 0
puts aa['b'] ? aa['b'] : 0
puts aa[1]

testing(generate(10))
testing(generate(1))
