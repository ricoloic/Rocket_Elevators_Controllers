function return_positive(n) {
    if (n < 0) {
        n *= -1
    }

    return n;
};

class Elevator {
    constructor (_id, _floor_amount) {
        this.ID = _id + 1
        this.floor_amount = _floor_amount
        this.btn = []
        this.points
        this.stop_list = []
        this.up_buffer = []
        this.down_buffer =  []
        this.current_direction
        this.previous_direction
        this.current_floor = 0
        this.door = "closed"    // false = closed / true = open
        this.status = "IDLE"    // IDLE / maintenance / moving
    
        for (let i = 1; i <= this.floor_amount; i++) {
            this.btn.push(i)
        }
    }

    check_in (n) {
        let in_list = true

        for (let i = 0; i < this.stop_list.length; i++) {
            if (n == this.stop_list[i]) {
                in_list = false
            }
        }

        return in_list;
    }

    door_state () {
        this.door = "open"
        console.log("the door is " + this.door + "\n...")

        this.door = "closed"
        console.log("the door is " + this.door + "\n...")
    }

    list_sorting () {
        if (this.current_direction == "up") {
            this.stop_list.sort(function(a, b) {return a - b})

        } else if (this.current_direction == "down") {
            this.stop_list.sort(function(a, b) {return b - a})

        } else {
            this.stop_list.sort(function(a, b) {return a - b})
        }
    }

    send_request (RequestedFloor) {
        if (this.check_in(RequestedFloor)) {
            this.stop_list.push(RequestedFloor)

            if (this.current_floor > RequestedFloor) {
                this.status = "MOVING"
                this.current_direction = "down"
            }

            else if (this.current_floor < RequestedFloor) {
                this.status = "MOVING"
                this.current_direction = "up"
            }

            this.list_sorting()
            console.log("requested floor " + RequestedFloor + "\n...")
        }
    }

    add_stop (RequestFloor, Direction) {
        if (this.check_in(RequestFloor)) {
            if (Direction == this.current_direction && Direction == "up" && RequestFloor >= this.current_floor) {
                this.stop_list.push(RequestFloor)

            } else if (Direction == this.current_direction && Direction == "down" && RequestFloor <= this.current_floor) {
                this.stop_list.push(RequestFloor)

            } else if (this.status == "IDLE") {
                this.stop_list.push(RequestFloor)

            } else if (Direction == "up") {
                this.up_buffer.push(RequestFloor)

            } else if (Direction == "down") {
                this.down_buffer.push(RequestFloor)
            }

            this.list_sorting()
        }
    }

    points_update (RequestFloor, Direction) {
        let difference_last_stop = 0
        if (this.status == "IDLE") {
            let dif_last_stop = this.stop_list[this.stop_list.length] - RequestFloor
            difference_last_stop = return_positive(dif_last_stop)
        }

        let max_floor_difference = this.floor_amount + 1

        let dif_floor = this.current_floor - RequestFloor
        let difference_floor = return_positive(dif_floor)

        this.points = 0

        if (this.current_direction == Direction && this.status == "IDLE") {
            if (RequestFloor >= this.current_floor && Direction == "up" || RequestFloor <= this.current_floor && Direction == "down") {
                this.points = difference_floor
                this.points += this.stop_list.length

            } else if (RequestFloor < this.current_floor && Direction == "up" || RequestFloor > this.current_floor && Direction == "down") {
                this.points = max_floor_difference
                this.points  += difference_last_stop + this.stop_list.length
            }

        } else if (this.status == "IDLE") {
            this.points = max_floor_difference
            this.points += difference_floor

        } else if (this.current_direction == Direction && this.status == "IDLE") {
            this.points = max_floor_difference * 2
            this.points += difference_last_stop + this.stop_list.length
        }
    }

    stop_switch () {
        if (this.down_buffer.length != 0 && this.up_buffer.length != 0) {
            if (this.previous_direction == "up") {
                this.current_direction = "down"
                for (let i = 0; i < this.down_buffer.length; i++) {
                    this.stop_list.push(this.down_buffer[0])
                    this.down_buffer.splice(0, 1)
                }

            } else if (this.previous_direction == "down") {
                this.current_direction = "up"
                for (let i = 0; i < this.up_buffer.length; i++) {
                    this.stop_list.push(this.up_buffer[0])
                    this.up_buffer.splice(0, 1)
                }
            }

        } else if (this.down_buffer.length != 0 && this.up_buffer.length == 0) {
            this.current_direction = "down"
            for (let i = 0; i < this.down_buffer.length; i++) {
                this.stop_list.push(this.down_buffer[0])
                this.down_buffer.splice(0, 1)
            }

        } else if (this.down_buffer.length == 0 && this.up_buffer.length != 0) {
            this.current_direction = "up"
            for (let i = 0; i < this.up_buffer.length; i++) {
                this.stop_list.push(this.up_buffer[0])
                this.up_buffer.splice(0, 1)
            }

        } else if (this.down_buffer.length == 0 && this.up_buffer.length == 0) {
            this.status = "IDLE"
            this.current_direction = "stop"
        }

        if (this.stop_list.length != 0) {
            this.list_sorting()
            this.run()
        }
    }

    run () {
        while (this.stop_list.length != 0) {
            if (this.stop_list.length != 0) {
                while (this.current_floor != this.stop_list[0]) {
                    if (this.stop_list[0] < this.current_floor) {
                        this.current_direction = "down"
                        this.previous_direction = this.current_direction
                        this.current_floor -= 0.25
                        this.status = "MOVING"

                    } else if (this.stop_list[0] > this.current_floor) {
                        this.current_direction = "up"
                        this.previous_direction = this.current_direction
                        this.current_floor += 0.25
                        this.status = "MOVING"
                    }
                }

                if (this.stop_list[0] == this.current_floor) {
                    console.log("elevator " + this.ID + " arrived at floor " + this.stop_list[0] + "\n...")
                    this.door_state()
                    this.stop_list.splice(0, 1)
                }

            } else if (this.stop_list.length == 0) {
                this.stop_switch()
            }
        } 
        
        if (this.stop_list.length == 0) {
            this.stop_switch()
        }
    }
}

class Column {
    constructor (Elevator, _floor_amount, _elevator_per_col) {
        this.floor_amount = _floor_amount
        this.elevator_per_col = _elevator_per_col
        this.elevator_list = []

        for (let i = 0; i < this.elevator_per_col; i++) {
            let e = new Elevator(i, this.floor_amount)
            this.elevator_list.push(e)
        }
    }

    RequestElevator (RequestFloor, Direction) {
        console.log("user request from floor " + RequestFloor + "\n...")

        for (let i = 0; i < this.elevator_list.length; i++) {
            this.elevator_list[i].points_update(RequestFloor, Direction)
        }

        let best_elevator = this.elevator_list.reduce((prev, current) => (prev.points > current.points) ? prev : current)
        console.log("sending elevator " + best_elevator.ID + "\n...")
        best_elevator.add_stop(RequestFloor, Direction)
        best_elevator.run()
        return best_elevator;
    }

    RequestFloor (Elevator, RequestedFloor) {
        Elevator.send_request(RequestedFloor)
        Elevator.run()
    }
}

let col = new Column(Elevator, 10, 2)

col.elevator_list[0].current_floor = 5
col.elevator_list[0].stop_list = [6, 8, 10]
col.elevator_list[0].down_buffer = [7, 3]
col.elevator_list[0].current_direction = "up"
col.elevator_list[0].status = "MOVING"
col.elevator_list[1].current_floor = 3
col.elevator_list[1].stop_list = [4, 5, 6, 7]
col.elevator_list[1].current_direction = "up"
col.elevator_list[1].status = "MOVING"

let elevator = col.RequestElevator(4, "down")
col.RequestFloor(elevator, 10)

for (let i = 0; i < col.elevator_list.length; i++) {
    col.elevator_list[i].run()
}