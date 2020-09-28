def return_positive(n):
    if n < 0:
        n *= -1
        return n
    
    else:
        return n

class User:
    def __init__(self, _floor, _floor_amount):
        self.floor_amount = _floor_amount
        self.current_floor = _floor
        self.btn = []
        self.current_direction = None

        if self.current_floor == 0:
            self.btn.append("up")
        
        elif self.current_floor == self.floor_amount - 1:
            self.btn.append("down")

        else:
            self.btn.append("up")
            self.btn.append("down")


    
    """ def iterate_btn_list(self):
        for i in self.btn:
            print(i, end = "", flush = True) """
   


class Elevator:
    def __init__(self, _floor_amount):
        self.floor_amount = _floor_amount
        self.btn = []
        self.points = 0
        self.stop_list = []
        self.up_buffer = []
        self.down_buffer=  []
        self.current_direction = None
        self.previous_direction = None
        self.current_floor = 0
        self.door = False    # false = closed / true = open
        self.status = "IDLE"    # IDLE / maintenance / moving

        for i in range(_floor_amount):
            self.btn.append(i)


    def list_sorting(self):
        if self.current_direction == "up":
            self.stop_list.sort()

        elif self.current_direction == "down":
            self.stop_list.sort(reverse = True)


    def add_stop(self, request):
        if request.current_direction == self.current_direction and request.current_direction == "up" and request.current_floor >= self.current_floor:
            self.stop_list.append(request.current_floor)

        elif request.current_direction == self.current_direction and request.current_direction == "down" and request.current_floor <= self.current_floor:
            self.stop_list.append(request.current_floor)

        elif self.status == "IDLE":
            self.stop_list.append(request.current_floor)

        elif request.current_direction == "up":
            self.up_buffer.append(request.current_floor)

        elif request.current_direction == "down":
            self.down_buffer.append(request.current_floor)

        self.list_sorting()


    def points_update(self, RequestFloor, Direction):
        if self.status != "IDLE":
            dif_last_stop = self.stop_list - RequestFloor
            difference_last_stop = return_positive(dif_last_stop)

        max_floor_difference = self.floor_amount + 1

        dif_floor = self.current_floor - RequestFloor
        difference_floor = return_positive(dif_floor)

        points = 0

        if self.current_direction == Direction and self.status != "IDLE":
            if RequestFloor >= self.current_floor and Direction == "up" or RequestFloor <= self.current_floor and Direction == "down":
                self.points = difference_floor
                self.points += len(self.stop_list)
            
            elif RequestFloor < self.current_floor and Direction == "up" or RequestFloor > self.current_floor and Direction == "down":
                self.points = max_floor_difference
                self.points += difference_last_stop + len(self.stop_list)

        elif self.status == "IDLE":
            self.points = max_floor_difference
            self.points += difference_floor

        elif self.current_direction != Direction and self.status != "IDLE":
            self.points = max_floor_difference * 2
            self.points += difference_last_stop + len(self.stop_list)


    def stop_switch(self):
        if len(self.down_buffer) != 0 and len(self.up_buffer) != 0:
            if self.previous_direction == "up":
                for i in self.down_buffer:
                    self.stop_list.append(i)
            
            elif self.previous_direction == "down":
                for i in self.up_buffer:
                    self.stop_list.append(i)

        elif len(self.down_buffer) != 0 and len(self.up_buffer) == 0:
            for i in self.down_buffer:
                self.stop_list.append(i)

        elif len(self.down_buffer) == 0 and len(self.up_buffer) != 0:
            for i in self.up_buffer:
                self.stop_list.append(i)

        elif len(self.down_buffer) == 0 and len(self.up_buffer) == 0:
            self.status = "IDLE"
            self.current_direction = "stop"


    def run(self):
        if len(self.stop_list) != 0:
            if self.stop_list[0] < self.current_floor:
                self.current_direction = "down"
                self.previous_direction = self.current_direction
                self.current_floor -= 1
                self.status = "MOVING"

            elif self.stop_list[0] > self.current_floor:
                self.current_direction = "up"
                self.previous_direction = self.current_direction
                self.current_floor += 1
                self.status = "MOVING"

            elif self.stop_list[0] == self.current_floor:
                del self.stop_list[0]
                print("Arrived at floor : " + self.current_floor)

        elif len(self.stop_list) == 0:
            self.stop_switch()


class Column:
    def __init__(self, Elevator, User, fa, epc):
        self.floor_amount = fa
        self.elevator_per_col = epc
        self.user_list = []
        self.elevator_list = []

        for i in range(self.floor_amount):
            user = User(i, self.floor_amount)
            self.user_list.append(user)

        for i in range(self.elevator_per_col):
            elevator = Elevator(self.floor_amount)
            self.elevator_list.append(elevator)


    def tell_elevator(self):
        j = 1
        for i in self.elevator_list:
            print(i)
            j += 1

    from operator import attrgetter

    def RequestElevator(self, RequestFloor, Direction):
        j = 0
        for i in self.elevator_list:
            self.elevator_list[j].points_update(RequestFloor, Direction)

        

        

    """ def tell_floor_btn(self):
        j = 0
        for i in self.user_list:
            self.user_list[j].iterate_btn_list()
            j += 1 """


col1 = Column(Elevator, User, 10, 2)

#col1.tell_floor_btn()

print(col1.elevator_list[0].btn)

col1.