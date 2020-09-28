from operator import attrgetter
import time

def return_positive(n):
    if n < 0:
        n *= -1
        return n
    
    else:
        return n

def wait(t):
    time.sleep(t)
    print("...")


class Elevator:
    def __init__(self, _id, _floor_amount):
        self.ID = _id + 1
        self.floor_amount = _floor_amount
        self.btn = []
        self.points = 0
        self.stop_list = []
        self.up_buffer = []
        self.down_buffer=  []
        self.current_direction = None
        self.previous_direction = None
        self.current_floor = 0
        self.door = "closed"    # false = closed / true = open
        self.status = "IDLE"    # IDLE / maintenance / moving

        for i in range(self.floor_amount):
            self.btn.append(i)


    def door_state(self):
        self.door = "open"
        print("the doors are", self.door)
        wait(5)

        self.door = "closed"
        print("the doors are", self.door)
        wait(1)


    def list_sorting(self):
        if self.current_direction == "up":
            self.stop_list.sort()

        elif self.current_direction == "down":
            self.stop_list.sort(reverse = True)

        else:
            self.stop_list.sort()


    def send_request(self, RequestedFloor):
        requestFloor = RequestedFloor - 1
        self.stop_list.append(RequestedFloor)

        if self.current_floor > RequestedFloor:
            self.status = "MOVING"
            self.current_direction = "down"
        
        elif self.current_floor < RequestedFloor:
            self.status = "MOVING"
            self.current_direction = "up"

        self.list_sorting()


    def add_stop(self, RequestFloor, Direction):
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
        requestFloor = RequestFloor
        if self.status != "IDLE":
            dif_last_stop = self.stop_list[-1] - requestFloor
            difference_last_stop = return_positive(dif_last_stop)

        max_floor_difference = self.floor_amount + 1

        dif_floor = self.current_floor - requestFloor
        difference_floor = return_positive(dif_floor)

        self.points = 0

        if self.current_direction == Direction and self.status != "IDLE":
            if requestFloor >= self.current_floor and Direction == "up" or requestFloor <= self.current_floor and Direction == "down":
                self.points = difference_floor
                self.points += len(self.stop_list)
            
            elif requestFloor < self.current_floor and Direction == "up" or requestFloor > self.current_floor and Direction == "down":
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
                for i in len(self.down_buffer):
                    self.stop_list.append(i)
            
            elif self.previous_direction == "down":
                for i in len(self.up_buffer):
                    self.stop_list.append(i)

        elif len(self.down_buffer) != 0 and len(self.up_buffer) == 0:
            for i in len(self.down_buffer):
                self.stop_list.append(i)

        elif len(self.down_buffer) == 0 and len(self.up_buffer) != 0:
            for i in len(self.up_buffer):
                self.stop_list.append(i)

        elif len(self.down_buffer) == 0 and len(self.up_buffer) == 0:
            self.status = "IDLE"
            self.current_direction = "stop"

        if len(self.stop_list) != 0:
            self.list_sorting()
            self.run()


    def run(self):
        while len(self.stop_list) != 0:
            if len(self.stop_list) != 0:
                while self.current_floor != self.stop_list[0]:
                    if self.stop_list[0] < self.current_floor:
                        self.current_direction = "down"
                        self.previous_direction = self.current_direction
                        self.current_floor -= 0.25
                        self.status = "MOVING"

                    elif self.stop_list[0] > self.current_floor:
                        self.current_direction = "up"
                        self.previous_direction = self.current_direction
                        self.current_floor += 0.25
                        self.status = "MOVING"

                if self.stop_list[0] == self.current_floor:
                    print("elevator", self.ID, "arrived at floor", self.stop_list[0])
                    wait(1)
                    self.door_state()
                    del self.stop_list[0]

            elif len(self.stop_list) == 0:
                self.stop_switch()
                
        if len(self.stop_list) == 0:
            self.stop_switch()



class Column:
    def __init__(self, Elevator, fa, epc):
        self.floor_amount = fa
        self.elevator_per_col = epc
        self.elevator_list = []

        for i in range(self.elevator_per_col):
            elevator = Elevator(i, self.floor_amount)
            self.elevator_list.append(elevator)


    def RequestElevator(self, RequestFloor, Direction):
        wait(1)
        print("user request from floor", RequestFloor)
        wait(1)

        j = 0
        for i in self.elevator_list:
            self.elevator_list[j].points_update(RequestFloor, Direction)
            j += 1

        best_elevator = min(self.elevator_list, key = attrgetter('points'))
        print("sending elevator", best_elevator.ID)
        best_elevator.add_stop(RequestFloor, Direction)
        wait(1)
        best_elevator.run()
        return best_elevator


    def RequestFloor(self, Elevator, RequestedFloor):
        print("Requested floor", RequestedFloor)
        wait(1)
        Elevator.send_request(RequestedFloor)
        Elevator.run()


col = Column(Elevator, 10, 2)

col.elevator_list[0].current_floor = 4
col.elevator_list[1].current_floor = 3
col.elevator_list[1].stop_list = [5, 6, 7]
col.elevator_list[1].current_direction = "up"
col.elevator_list[1].status = "MOVING"


elevator = col.RequestElevator(4, "up")
col.RequestFloor(elevator, 10)