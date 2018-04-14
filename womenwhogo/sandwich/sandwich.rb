class Butler 
    def get_peanut_butter
        sleep 1
        return "peanut butter"
    end

    def get_jelly
        sleep 1
        return "jelly"
    end
end

b1 = Butler.new
b2 = Butler.new

t1 = Time.new
pb = b1.get_peanut_butter
j = b2.get_jelly
t2 = Time.new
puts "took #{t2 - t1} seconds to get the ingredients"