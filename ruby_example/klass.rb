class Wololo

    def initialize(val)
        @val = val
        @test = nil
    end

    def new_var(new_val)
        @new_val = new_val
    end

    def testing
        puts @new_val

        maps = {
            "val": @val,
            "test": @test ? test.ark : nil
        }
        puts maps
    end

    def test_if_else
        if @val == 10
            "Ten"
        else
            "Not Ten"
        end
    end
end

wl = Wololo.new(5)
wl.new_var(new_val = 10)
wl.testing

puts wl.test_if_else

puts nil
