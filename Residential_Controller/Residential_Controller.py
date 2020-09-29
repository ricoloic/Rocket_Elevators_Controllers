from operator import attrgetter
import time

def return_positive(n):
    if n < 0:
        n *= -1

    return n

def wait(t):
    t = 0.1
    print("...")
    time.sleep(t)


class Elevator:
    def __init__(self, _id, _floor_amount):
        self.ID = _id + 1
        self.floor_amount = _floor_amount
        self.btn = []
        self.points = 0
        self.stop_list = []
        self.up_buffer = []
        self.down_buffer =  []
        self.current_direction = None
        self.previous_direction = None
        self.current_floor = 0
        self.previous_floor = 0
        self.door = "closed"    # false = closed / true = open
        self.status = "IDLE"    # IDLE / maintenance / moving

        for i in range(self.floor_amount):
            self.btn.append(i + 1)


    def door_state(self):
        """ self.door = "open"
        print("the doors are", self.door)
        wait(5)

        self.door = "closed"
        print("the doors are", self.door)
        wait(1) """
        print("the doors are OPEN then closed")
        wait(0)



    def check_in(self, n):
        in_list = True

        for i in range(len(self.stop_list)):
            if n == self.stop_list[i]:
                in_list = False

        return in_list
    


    def list_sorting(self):
        if self.current_direction == "up":
            self.stop_list.sort()

        elif self.current_direction == "down":
            self.stop_list.sort(reverse = True)

        else:
            self.stop_list.sort()


    def send_request(self, RequestedFloor):
        if self.check_in(RequestedFloor):
            self.stop_list.append(RequestedFloor)

            if self.current_floor > RequestedFloor:
                self.status = "MOVING"
                self.current_direction = "down"
            
            elif self.current_floor < RequestedFloor:
                self.status = "MOVING"
                self.current_direction = "up"

            self.list_sorting()
            print("requested floor", RequestedFloor)
            wait(1)


    def add_stop(self, RequestFloor, Direction):
        if self.check_in(RequestFloor):
            if Direction == self.current_direction and Direction == "up" and RequestFloor >= self.current_floor:
                self.stop_list.append(RequestFloor)

            elif Direction == self.current_direction and Direction == "down" and RequestFloor <= self.current_floor:
                self.stop_list.append(RequestFloor)

            elif self.status == "IDLE":
                self.stop_list.append(RequestFloor)

            elif Direction == "up":
                self.up_buffer.append(RequestFloor)

            elif Direction == "down":
                self.down_buffer.append(RequestFloor)

            self.list_sorting()


    def points_update(self, RequestFloor, Direction):
        difference_last_stop = 0
        if self.status != "IDLE":
            dif_last_stop = self.stop_list[-1] - RequestFloor
            difference_last_stop = return_positive(dif_last_stop)

        max_floor_difference = self.floor_amount + 1

        dif_floor = self.current_floor - RequestFloor
        difference_floor = return_positive(dif_floor)

        self.points = 0

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
            print("elevator", self.ID, "is changing direction")
            if self.previous_direction == "up":
                self.current_direction = "down"
                for i in self.down_buffer:
                    self.stop_list.append(i)
                    del self.down_buffer[0]
            
            elif self.previous_direction == "down":
                self.current_direction = "up"
                for i in self.up_buffer:
                    self.stop_list.append(i)
                    del self.up_buffer[0]

        elif len(self.down_buffer) != 0 and len(self.up_buffer) == 0:
            print("elevator", self.ID, "is changing direction")
            self.current_direction = "down"
            for i in self.down_buffer:
                self.stop_list.append(i)
                del self.down_buffer[0]

        elif len(self.down_buffer) == 0 and len(self.up_buffer) != 0:
            print("elevator", self.ID, "is changing direction")
            self.current_direction = "up"
            for i in self.up_buffer:
                self.stop_list.append(i)
                del self.up_buffer[0]

        elif len(self.down_buffer) == 0 and len(self.up_buffer) == 0:
            self.status = "IDLE"
            self.current_direction = "stop"

        if len(self.stop_list) != 0:
            self.list_sorting()
            self.run()


    def run(self):
        print("running elevator :", self.ID)
        while len(self.stop_list) != 0:
            if len(self.stop_list) != 0:
                while self.current_floor != self.stop_list[0]:
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
                    
                    if self.current_floor != self.previous_floor:
                        print("floor :", self.current_floor)
                        self.previous_floor = self.current_floor
                
                if self.stop_list[0] == self.current_floor:
                    print("elevator", self.ID, "arrived at floor", self.stop_list[0])
                    self.door_state()
                    del self.stop_list[0]

            elif len(self.stop_list) == 0:
                self.stop_switch()

        if len(self.stop_list) == 0:
            self.stop_switch()


class Column:
    def __init__(self, Elevator, _floor_amount, _elevator_per_col):
        self.floor_amount = _floor_amount
        self.elevator_per_col = _elevator_per_col
        self.elevator_list = []

        for i in range(self.elevator_per_col):
            e = Elevator(i, self.floor_amount)
            self.elevator_list.append(e)


    def RequestElevator(self, RequestedFloor, Direction):
        wait(1)
        print("user request from floor", RequestedFloor)
        wait(1)

        j = 0
        for i in self.elevator_list:
            self.elevator_list[j].points_update(RequestedFloor, Direction)
            j += 1

        best_elevator = min(self.elevator_list, key = attrgetter('points'))
        print("sending elevator", best_elevator.ID)
        best_elevator.add_stop(RequestedFloor, Direction)
        wait(1)
        best_elevator.run()
        return best_elevator


    def RequestFloor(self, Elevator, RequestedFloor):
        Elevator.send_request(RequestedFloor)
        Elevator.run()

    
    def change_value(self, elevator, current_floor, stop_list, down_buffer, up_buffer, current_direction, status):
        self.elevator_list[elevator].current_floor = current_floor
        self.elevator_list[elevator].stop_list = stop_list
        self.elevator_list[elevator].down_buffer = down_buffer
        self.elevator_list[elevator].up_buffer = up_buffer
        self.elevator_list[elevator].current_direction = current_direction
        self.elevator_list[elevator].status = status
        self.elevator_list[elevator].list_sorting()
        

# test
col = Column(Elevator, 10, 2)

# column.change_value(elevator, current_floor, stop_list, down_buffer, up_buffer, current_direction, status)

# Scenario Custom
""" col.change_value(0, 9, [7, 6, 5, 3], [], [4, 10], "down", "MOVING")
col.change_value(1, 5, [6, 8, 10], [7, 3], [2, 5], "up", "MOVING")

elevator = col.RequestElevator(4, "down")
col.RequestFloor(elevator, 10) """

# Scenario 1
""" col.change_value(0, 2, [], [], [], "stop", "IDLE")
col.change_value(1, 6, [], [], [], "stop", "IDLE")

elevator = col.RequestElevator(3, "up")
col.RequestFloor(elevator, 7) """

# Scenario 2
col.change_value(0, 10, [], [], [], "stop", "IDLE")
col.change_value(1, 3, [], [], [], "stop", "IDLE")

elevator = col.RequestElevator(1, "up")
col.RequestFloor(elevator, 6)

elevator = col.RequestElevator(3, "up")
col.RequestFloor(elevator, 5)

elevator = col.RequestElevator(9, "down")
col.RequestFloor(elevator, 2)

# Scenario 3
""" col.change_value(0, 10, [], [], [], "stop", "IDLE")
col.change_value(1, 3, [6], [], [], "up", "MOVING")

elevator = col.RequestElevator(3, "down")
col.RequestFloor(elevator, 2)

for elevator in col.elevator_list:
    elevator.list_sorting()
    elevator.run()

elevator = col.RequestElevator(10, "down")
col.RequestFloor(elevator, 3) """


for elevator in col.elevator_list:
    elevator.list_sorting()
    elevator.run()