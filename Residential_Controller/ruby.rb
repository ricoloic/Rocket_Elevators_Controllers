up = "up"
name_list = ["up", "down"]
down = "down"

print(name_list[0] == up)
print(name_list[1] == down)

class Book
    def initialize (pages_number)
        @pages_number = pages_number
        @pages = []

        for i in 1..@pages_number
            @pages.push(i)
        end
    end

    def page
        print(@pages)
    end
end

book = Book.new(25)

book.page

class Column
    attr_accessor :floor_amount, :elevator_per_col
    def initialize(floor_amount, elevator_per_col)
        @floor_amount = floor_amount
        @elevator_per_col = elevator_per_col
        @elevator_list = []

        for i in 1..self.elevator_per_col
            e = Book.new(i)
            self.elevator_list.push(e)
        end
    end
end