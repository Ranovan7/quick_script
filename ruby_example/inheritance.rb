class Car

  def fuel
    puts "Gasoline"
  end
end

class Tesla < Car

  def fuel
    puts "Electricity"
  end
end


model_x = Tesla.new()
model_x.fuel
