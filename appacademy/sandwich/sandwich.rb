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

class Customer
    def get_sandwich(butler1, butler2)
        ingredient_one = butler1.get_peanut_butter
        ingredient_two = butler2.get_jelly

        puts "putting together #{ingredient_one} and #{ingredient_two}"
    end
end

w1 = Butler.new
w2 = Butler.new
c = Customer.new

t1 = Time.new
c.get_sandwich(w1, w2)
t2 = Time.new
puts "took #{t2 - t1} seconds"