require 'json'

arr = "[{\"order_id\": 4},{\"order_id\": 5},{\"order_id\": 6},{\"order_id\": 7}]"

narr = JSON.parse(arr).map{|a| a['order_id']}
for n in narr do
    puts n
end

def test
    true
end

puts test
puts test.to_i
