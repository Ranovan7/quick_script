require 'date'

def datetime2int
    puts (DateTime.parse("2022-01-17T10:26:58.507904Z").to_time.to_i rescue nil)
    puts (DateTime.parse("2022-01-17T10:27:32Z").to_time.to_i rescue nil)
end

File.write("log.txt", "Testing")

puts "qwertsdfsdtg".to_s
datetime2int

ex = "Gachagame named "
ex += "Arknights"
puts ex

begin
    puts "Arknights"
rescue
    puts "Error Occured"
else
    puts "Else after Rescue"
end
