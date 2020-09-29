def return_positive(n)
    positive = n
    if positive < 0
        positive *= -1
    end
    return positive
end

def bubble_sort(list)
    return list if list.size <= 1 # already sorted
    swapped = true
    while swapped do
        swapped = false
        0.upto(list.size-2) do |i|
            if list[i].points > list[i+1].points
                list[i], list[i+1] = list[i+1], list[i] # swap values
                swapped = true
            end
        end
    end

    list
end

class Elevator
    attr_accessor :points, :ID
    def initialize(_id, _floor_amount)
        @ID = _id
        @floor_amount = _floor_amount
        @btn = []
        @points = 0
        @stop_list = []
        @up_buffer = []
        @down_buffer =  []
        @current_direction
        @previous_direction
        @current_floor = 0
        @previous_floor = 0
        @door = "closed"    # closed / open
        @status = "IDLE"    # IDLE / maintenance / moving
        @name_list = ["up", "down", "IDLE"]
        
        for i in 1..@floor_amount
            @btn.push(i + 1)
        end
    end
    
    def door_state
        puts "the doors are OPEN then closed \n..."
    end

    def check_in(requestedFloor)
        in_list = true

        for i in 0..@stop_list.length
            if requestedFloor == @stop_list[i]
                in_list = false
            end
        end
        return in_list
    end

    def list_sorting()
        if @current_direction == @name_list[0]
            @stop_list.sort!
        
        elsif @current_direction == @name_list[1]
            @stop_list.sort! {|x, y| y <=> x}

        else
            @stop_list.sort!
        end
    end

    def send_request(requestedFloor)
        @stop_list.push(requestedFloor)

        if @current_floor > requestedFloor
            @status = "MOVING"
            @current_direction = "down"
        
        elsif @current_floor < requestedFloor
            @status = "MOVING"
            @current_direction = "up"
        end

        self.list_sorting
        puts "requested floor #{requestedFloor} \n..."
    end

    def add_stop(requestFloor, direction)
        if direction == @current_direction and direction == @name_list[0] and requestFloor >= @current_floor
            @stop_list.push(requestFloor)

        elsif direction == @current_direction and direction == @name_list[1] and requestFloor <= @current_floor
            @stop_list.push(requestFloor)

        elsif @status == @name_list[2]
            @stop_list.push(requestFloor)

        elsif direction == @name_list[0]
            @up_buffer.push(requestFloor)

        elsif direction == @name_list[1]
            @down_buffer.push(requestFloor)
        end

        self.list_sorting
    end

    def points_update(requestFloor, direction)
        difference_last_stop = 0
        if @status != @name_list[2]
            dif_last_stop = @stop_list.last - requestFloor
            difference_last_stop = return_positive(dif_last_stop)
        end

        max_floor_difference = @floor_amount + 1

        dif_floor = @current_floor - requestFloor
        difference_floor = return_positive(dif_floor)

        @points = 0

        if @current_direction == direction and @status != @name_list[2]
            if requestFloor >= @current_floor and direction == @name_list[0] or requestFloor <= @current_floor and direction == @name_list[1]
                @points = difference_floor
                @points += @stop_list.length

            elsif requestFloor < @current_floor and direction == @name_list[0] or requestFloor > @current_floor and direction == @name_list[1]
                @points = max_floor_difference
                @points += difference_last_stop + @stop_list.length
            end

        elsif @status == @name_list[2]
            @points = max_floor_difference
            @points += difference_floor

        elsif @current_direction != direction and @status != @name_list[2]
            @points = max_floor_difference * 2
            @points += difference_last_stop + @stop_list.length
        end
    end

    def stop_switch()
        if @down_buffer.length != 0 and @up_buffer != 0
            puts "elevator #{@ID} is changing direction"
            if @previous_direction == @name_list[0]
                @current_direction = "down"

                for i in @down_buffer
                    @stop_list.push(i)
                    @down_buffer.delete_at(0)
                end
            
            elsif @previous_direction == @name_list[1]
                @current_direction = "up"

                for i in @up_buffer
                    @stop_list.push(i)
                    @up_buffer.delete_at(0)
                end
            end
        
        elsif @down_buffer.length != 0 and @up_buffer.length == 0
            puts "elevator #{@ID} is changing direction"
            @current_direction = "down"

            for i in @down_buffer
                @stop_list.push(i)
                @down_buffer.delete_at(0)
            end

        elsif @down_buffer.length == 0 and @up_buffer.length != 0
            puts "elevator #{@ID} is changing direction"
            @current_direction = "up"

            for i in @up_buffer
                @stop_list.push(i)
                @up_buffer.delete_at(0)
            end

        elsif @down_buffer.length == 0 and @up_buffer.length == 0
            @status = "IDLE"
            @current_direction = "stop"
        end

        if @stop_list.length != 0
            self.list_sorting
            self.run
        end
    end

    def run()
        puts "running elevator : #{@ID}"
        while @stop_list.length != 0
            if @stop_list.length !=0
                while @current_floor != @stop_list[0]
                    if @stop_list[0] < @current_floor
                        @current_direction = "down"
                        @previous_direction = @current_direction
                        @current_floor -= 1
                        @status = "MOVING"

                    elsif @stop_list[0] > @current_floor
                        @current_direction = "up"
                        @previous_direction = @current_direction
                        @current_floor += 1
                        @status = "MOVING"
                    end

                    if @current_floor != @previous_floor
                        puts "floor : #{@current_floor}"
                        @previous_floor = @current_floor
                    end
                end

                if @stop_list[0] == @current_floor
                    print "elevator #{@ID} arrived at floor #{@stop_list[0]}"
                    self.door_state
                    @stop_list.delete_at(0)
                end
            
            elsif @stop_list.length == 0
                self.stop_switch
            end
        end
        
        if @stop_list.length == 0
            self.stop_switch
        end
    end

    def change_value current_floor, stop_list, down_buffer, up_buffer, current_direction, status
        @current_floor = current_floor
        @stop_list = stop_list
        @down_buffer = down_buffer
        @up_buffer = up_buffer
        @current_direction = current_direction
        @status = status
        @list_sorting
    end
end

class Column
    attr_accessor :floor_amount, :elevator_list, :elevator_per_col
    def initialize(_floor_amount, _elevator_per_col)
        @floor_amount = _floor_amount
        @elevator_per_col = _elevator_per_col
        @elevator_list = []

        for i in 1..@elevator_per_col
            e = Elevator.new(i, @floor_amount)
            @elevator_list.push(e)
        end
    end

    def printElev
        puts @elevator_list
    end

    def requestElevator(requestedFloor, direction)
        puts "user request from floor #{requestedFloor} \n...\n"

        for i in 0..(@elevator_per_col - 1)
            @elevator_list[i].points_update requestedFloor, direction
        end
        
        bubble_sort(@elevator_list, points)
        best_elevator = @elevator_list[0]

        puts "SENDING ELEVATOR #{best_elevator.ID} \n..."
        best_elevator.add_stop requestedFloor, direction
        best_elevator.run
        return best_elevator
    end

    def runAll
        for i in 0..(@elevator_list.length - 1)
            @elevator_list[i].run
        end
    end

    def requestFloor(elevator ,requestedFloor)
        elevator.send_request requestedFloor
        elevator.run
    end

    def change_value(elevator, current_floor, stop_list, down_buffer, up_buffer, current_direction, status)
        @elevator_list[elevator].change_value current_floor, stop_list, down_buffer, up_buffer, current_direction, status
    end
end

col = Column.new(10, 2)

def scenario1(col)
    col.change_value(0, 2, [], [], [], "stop", "IDLE")
    col.change_value(1, 6, [], [], [], "stop", "IDLE")

    elevator = col.requestElevator(3, "up")
    col.requestFloor(elevator, 7)
end

def scenario2(col)
    col.change_value 0, 10, [], [], [], "stop", "IDLE"
    col.change_value 1, 3, [], [], [], "stop", "IDLE"

    elevator = col.requestElevator 1, "up"
    col.requestFloor elevator, 6

    elevator = col.requestElevator 3, "up"
    col.requestFloor elevator, 5

    elevator = col.requestElevator 9, "down"
    col.requestFloor elevator, 2
end

def scenario3(col)
    col.change_value 0, 10, [], [], [], "stop", "IDLE"
    col.change_value 1, 3, [6], [], [], "up", "MOVING"

    elevator = col.requestElevator 3, "down"
    col.requestFloor elevator, 2

    for i in 0..(col.elevator_list.length - 1)
        col.elevator_list[i].run
    end

    elevator = col.requestElevator 10, "down"
    col.requestFloor elevator, 3
end

# scenario1(col)
# scenario2(col)
# scenario3(col)

col.runAll