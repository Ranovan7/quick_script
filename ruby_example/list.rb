require 'json'

def testing(a, b=nil)
    puts "a = #{a}"

    puts "b = #{b}" unless b.nil?
end

arr = "[{\"order_id\": 4},{\"order_id\": 5},{\"order_id\": 6},{\"order_id\": 7}]"

narr = JSON.parse(arr).map{|a| a['order_id']}
for n in narr do
    puts n
end

def test
    true
end

puts test

testing(10)
testing(50, 100)

asoiaf = nil
puts asoiaf.this_function_does_not_exists
