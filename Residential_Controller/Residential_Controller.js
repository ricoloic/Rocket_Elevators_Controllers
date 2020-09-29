function return_positive(n) {
    let N = n
    if (N < 0) {
        N *= -1
    }

    return N;
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
        this.previous_floor = 0
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

    send_request (requestedFloor) {
        if (this.check_in(requestedFloor)) {
            this.stop_list.push(requestedFloor)

            if (this.current_floor > requestedFloor) {
                this.status = "MOVING"
                this.current_direction = "down"
            }

            else if (this.current_floor < requestedFloor) {
                this.status = "MOVING"
                this.current_direction = "up"
            }

            this.list_sorting()
            console.log("requested floor " + requestedFloor + "\n...")
        }
    }

    add_stop (requestFloor, direction) {
        if (this.check_in(requestFloor)) {
            if (direction == this.current_direction && direction == "up" && requestFloor >= this.current_floor) {
                this.stop_list.push(requestFloor)

            } else if (direction == this.current_direction && direction == "down" && requestFloor <= this.current_floor) {
                this.stop_list.push(requestFloor)

            } else if (this.status == "IDLE") {
                this.stop_list.push(requestFloor)

            } else if (direction == "up") {
                this.up_buffer.push(requestFloor)

            } else if (direction == "down") {
                this.down_buffer.push(requestFloor)
            }

            this.list_sorting()
        }
    }

    points_update (requestFloor, direction) {
        let difference_last_stop = 0
        if (this.status != "IDLE") {
            let dif_last_stop = this.stop_list[this.stop_list.length - 1] - requestFloor
            difference_last_stop = return_positive(dif_last_stop)
        }

        let max_floor_difference = this.floor_amount + 1

        let dif_floor = this.current_floor - requestFloor
        let difference_floor = return_positive(dif_floor)

        this.points = 0

        if (this.current_direction == direction && this.status != "IDLE") {
            if (requestFloor >= this.current_floor && direction == "up" || requestFloor <= this.current_floor && direction == "down") {
                this.points = difference_floor
                this.points += this.stop_list.length

            } else if (requestFloor < this.current_floor && direction == "up" || requestFloor > this.current_floor && direction == "down") {
                this.points = max_floor_difference
                this.points  += difference_last_stop + this.stop_list.length
            }

        } else if (this.status == "IDLE") {
            this.points = max_floor_difference
            this.points += difference_floor

        } else if (this.current_direction != direction && this.status != "IDLE") {
            this.points = max_floor_difference * 2
            this.points += difference_last_stop + this.stop_list.length
        }
    }

    stop_switch () {
        if (this.down_buffer.length != 0 && this.up_buffer.length != 0) {
            console.log("elevator " + this.ID + " is changing direction")
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
            console.log("elevator " + this.ID + " is changing direction")
            this.current_direction = "down"
            for (let i = 0; i < this.down_buffer.length; i++) {
                this.stop_list.push(this.down_buffer[0])
                this.down_buffer.splice(0, 1)
            }

        } else if (this.down_buffer.length == 0 && this.up_buffer.length != 0) {
            console.log("elevator " + this.ID + " is changing direction")
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
        console.log("running elevator : " + this.ID)
        while (this.stop_list.length != 0) {
            if (this.stop_list.length != 0) {
                while (this.current_floor != this.stop_list[0]) {
                    if (this.stop_list[0] < this.current_floor) {
                        this.current_direction = "down"
                        this.previous_direction = this.current_direction
                        this.current_floor -= 1
                        this.status = "MOVING"

                    } else if (this.stop_list[0] > this.current_floor) {
                        this.current_direction = "up"
                        this.previous_direction = this.current_direction
                        this.current_floor += 1
                        this.status = "MOVING"
                    }

                    if (this.current_floor != this.previous_floor) {
                        console.log("floor : " + this.current_floor)
                        this.previous_floor = this.current_floor
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

    requestElevator (requestFloor, direction) {
        console.log("user request from floor " + requestFloor + "\n...")

        for (let i = 0; i < this.elevator_list.length; i++) {
            this.elevator_list[i].points_update(requestFloor, direction)
        }

        this.elevator_list.sort(function(a, b) {
            return parseFloat(a.points) - parseFloat(b.points);
        })

        let best_elevator = this.elevator_list[0]
        console.log("SENDING ELEVATOR " + best_elevator.ID + "\n...")
        best_elevator.add_stop(requestFloor, direction)
        best_elevator.run()
        return best_elevator;
    }

    requestFloor (Elevator, requestedFloor) {
        Elevator.send_request(requestedFloor)
        Elevator.run()
    }

    change_value (elevator, current_floor, stop_list, down_buffer, up_buffer, current_direction, status) {
        this.elevator_list[elevator].current_floor = current_floor
        this.elevator_list[elevator].stop_list = stop_list
        this.elevator_list[elevator].down_buffer = down_buffer
        this.elevator_list[elevator].up_buffer = up_buffer
        this.elevator_list[elevator].current_direction = current_direction
        this.elevator_list[elevator].status = status
        this.elevator_list[elevator].list_sorting()
    }
}

// test
var col = new Column(Elevator, 10, 2)

// column.change_value(elevator, current_floor, stop_list, down_buffer, up_buffer, current_direction, status)

function scenario1 () {
    col.change_value(0, 2, [], [], [], "stop", "IDLE")
    col.change_value(1, 6, [], [], [], "stop", "IDLE")
    
    let elevator = col.requestElevator(3, "up")
    col.requestFloor(elevator, 7)
}

function scenario2 () {
    col.change_value(0, 10, [], [], [], "stop", "IDLE")
    col.change_value(1, 3, [], [], [], "stop", "IDLE")

    let elevator = col.requestElevator(1, "up")
    col.requestFloor(elevator, 6)

    elevator = col.requestElevator(3, "up")
    col.requestFloor(elevator, 5)

    elevator = col.requestElevator(9, "down")
    col.requestFloor(elevator, 2)

    elevator = col.requestElevator(4, "down")
    col.requestFloor(elevator, 10)
}

function scenario3 () {
    col.change_value(0, 10, [], [], [], "stop", "IDLE")
    col.change_value(1, 3, [6], [], [], "up", "MOVING")

    let elevator = col.requestElevator(3, "down")
    col.requestFloor(elevator, 2)

    for (let i = 0; i < col.elevator_list.length; i++) {
        col.elevator_list[i].list_sorting()
        col.elevator_list[i].run()
    }

    elevator = col.requestElevator(10, "down")
    col.requestFloor(elevator, 3)
}

function customScenario () {
    col.change_value(0, 9, [7, 6, 5, 3], [], [4, 10], "down", "MOVING")
    col.change_value(1, 5, [6, 8, 10], [7, 3], [2, 5], "up", "MOVING")

    let elevator = col.requestElevator(4, "down")
    col.requestFloor(elevator, 10)
}

// scenario1()
// scenario2()
scenario3()

// customScenario()

for (let i = 0; i < col.elevator_list.length; i++) {
    col.elevator_list[i].list_sorting()
    col.elevator_list[i].run()
}