// var nb = [[10, 100, 1000], 20, 30, 40, 50,]

// console.log(nb[0][1]);

/* for (var i of nb) {
    for (var j of i) {
        console.log(j + "\n")
    }
} */

/* for (var i = 0; i < nb.length; i++) {
    if (i == 0) {
        for (var j of nb[i]) {
            console.log(j + "\n")
        }
    } else {
        console.log(nb[i] + "\n")
    }
}

for (var i of nb) {
    console.log(i + "\n")
}

var i;
while (nb.length != 0) {
    i = nb.length
    nb.splice(0, 1)
    console.log(i)
} */
class Column {
    constructor(floorAmount, elevatorAmount) {
        console.log('column constructor', floorAmount, elevatorAmount);
        this.floorAmount = floorAmount;
        this.elevatorAmount = elevatorAmount;
        this.elevatorList = [];
        this.callButtonList = [];
        this.floorRequestButtonList = [];
        this.bestElevator;

        for (let i = 1; i <= this.elevatorAmount; i++) {
            let elevator = new Elevator("elevator"+i, "idle", 1, "none", 10, "closed");
            this.elevatorList.push(elevator);
        }
        
        for (let i = 1; i <= this.floorAmount; i++) {
            if (i == 1) {
                let callButton = new CallButton("callButtonUP"+i, i, "up", false);
                this.callButtonList.push(callButton);
            } 
            else if (i == 10) {
                let callButton = new CallButton("callButtonDOWN"+i, i, "down", false);
                this.callButtonList.push(callButton);
            }    
            else {
                this.callButtonList.push(new CallButton("callButtonDOWN"+i, i, "down", false));
                this.callButtonList.push(new CallButton("callButtonUP"+i, i, "up", false));
            }
        }
        
        for (let i = 1; i <= this.floorAmount; i++) {
            let floorRequestButton = new FloorRequestButton("floorButton"+i, i);
            this.floorRequestButtonList.push(floorRequestButton);
        }
                
        console.log("ELEVATORS LIST");
        console.table(this.elevatorList);
        console.log("CALL BUTTON LIST");
        console.table(this.callButtonList);
        console.log("FLOOR REQUEST BUTTON LIST");
        console.table(this.floorRequestButtonList);
        
    }
    
    //CallButton pressed on floor
    requestElevator(floor, direction) {

        console.log("LOOKING FOR BEST ELEVATOR...");

        this.bestElevator = this.findBestElevator(floor, direction);
        
        console.log("BEST ELEVATOR FOUND...");
        console.log("Elevator " + this.bestElevator.id + " has been requested at Floor " + floor);
        
        this.bestElevator.requestList.push(floor);
        this.bestElevator.moveElevator();

        console.log(this.bestElevator.id);
        return this.bestElevator;
    }
        
    //Find the Best Elevator to Send to Call Button Floor 
    findBestElevator(floor, direction) {
        
        for (let i = 0; i < this.elevatorList.length; i++) {
            
            if (floor === this.elevatorList[i].position && this.elevatorList[i].direction === "idle") {
                var bestCase = this.elevatorList[i];
            }
            else if (direction === "up" && (this.elevatorList[i].direction === "up" || this.elevatorList[i].direction === "idle") && this.elevatorList[i].position <= floor){
                var bestCase = this.elevatorList[i];
                
            }
            else if (direction === "down" && (this.elevatorList[i].direction === "down" || this.elevatorList[i].direction === "idle") && this.elevatorList[i].position >= floor){
                var bestCase = this.elevatorList[i];
            }
        }

        let minDistance = 999;
        let bestDistance = 11;
        for (let  i = 0; i < this.elevatorList.length; i++) {
            let distance = Math.abs(this.elevatorList[i].position - floor);
            
            if (this.elevatorList[i].direction === "idle" && bestDistance >= distance) {
                minDistance = distance;
                var bestIdle = this.elevatorList[i];
            }
        }
        
        if (bestCase != null) {
            return bestCase;
        } else {
            return bestIdle;
        }
         
    }

 
    //FloorButton pressed inside Elevator
    requestFloor(elevator, floor) {
        console.log("\n---> REQUESTED FLOOR : " + floor + "\n");
        console.log(elevator.id);
        elevator.requestList.push(floor);
        elevator.moveElevator();
    }
}


class Elevator {
    constructor(id, status, position, direction, floor, doors) {
        console.log('elevator constructor', id, status, position, direction, floor, doors);
        this.id = id;
        this.status = status;
        this.position = position;
        this.direction = direction;
        this.floor = floor;
        this.doors = doors;
        this.requestList = [];
        
        
        
    }

    //Move Elevator
    moveElevator() {
        var previousPosition = this.position;
        while (this.requestList.length != 0) {
            if (this.position > this.requestList[0]) {
                this.position--;
            } 
            else if (this.position < this.requestList[0]) {
                this.position++;
            } 
            else if (this.position == this.requestList[0]) {
                console.log("Elevator " + this.id + " arrived at Floor " + this.position);
                this.requestList.splice(0, 1);
            }
            
            if (previousPosition != this.position) {
                console.log("Elevator " + this.id + " is moving... (Floor " + this.position + ")");
                previousPosition = this.position;
            }
            
        }
    }
}

class CallButton {
    constructor(id, floor, direction) {
        this.id = id;
        this.direction = direction;
        this.floor = floor;
    }
}


class FloorRequestButton {
    constructor(id, floorAmount) {
        this.id = id;
        this.floorAmount = floorAmount;
    }
}



/* #################################### TEST ZONE ####################################*/
var column1 = new Column(10, 2);

/* SCENARIO 1 */
function scenario1() {
    console.log("\n-----------------------\n");
    console.log("SCENARIO 1\n");
    column1.elevatorList[0].id = "A";
    column1.elevatorList[0].position = 2;
    column1.elevatorList[0].direction = 'idle';
    column1.elevatorList[0].status = 'idle';
    column1.elevatorList[0].floor = 3;
    column1.elevatorList[1].id = "B";
    column1.elevatorList[1].position = 6;
    column1.elevatorList[1].direction = 'idle';
    column1.elevatorList[1].status = 'idle';
    column1.elevatorList[1].floor = 3;

    console.log("FIRST CALL \n");
    var elevator = column1.requestElevator(3, "up");
    column1.requestFloor(elevator, 7);
      
};


/* SCENARIO 2 */
function scenario2() {
    console.log("\n-----------------------\n");
    console.log("SCENARIO 2\n");
    column1.elevatorList[0].id = "A";
    column1.elevatorList[0].position = 10;
    column1.elevatorList[0].direction = 'idle';
    column1.elevatorList[0].status = 'idle';
    column1.elevatorList[0].floor = 1;
    column1.elevatorList[1].id = "B";
    column1.elevatorList[1].position = 3;
    column1.elevatorList[1].direction = 'idle';
    column1.elevatorList[1].status = 'idle';
    column1.elevatorList[1].floor = 1;

    console.log("FIRST CALL \n");
    var elevator = column1.requestElevator(1, "up");
    column1.requestFloor(elevator, 6);

    /* for (var i of column1.elevatorList) {
        i.moveElevator();
    }; */

    console.log("\nSECOND CALL \n");
    elevator = column1.requestElevator(3, "up");
    column1.requestFloor(elevator, 5);

    /* for (var i of column1.elevatorList) {
        i.moveElevator();
    }; */

    console.log("\nTHIRD CALL \n");
    elevator = column1.requestElevator(9, "down");
    column1.requestFloor(elevator, 2);
    
};


/* SCENARIO 3 */
function scenario3() {
    console.log("\n-----------------------\n");
    console.log("SCENARIO 3\n");
    column1.elevatorList[0].id = "A";
    column1.elevatorList[0].position = 10;
    column1.elevatorList[0].direction = 'idle';
    column1.elevatorList[0].status = 'idle';
    column1.elevatorList[0].floor = 3;
    column1.elevatorList[1].id = "B";
    column1.elevatorList[1].position = 3;
    column1.elevatorList[1].requestList = [6];
    column1.elevatorList[1].direction = 'up';
    column1.elevatorList[1].status = 'moving';
    column1.elevatorList[1].floor = 6;

    console.log("FIRST CALL \n");
    var elevator = column1.requestElevator(3, "down");
    column1.requestFloor(elevator, 2);

    for (var i of column1.elevatorList) {
        i.moveElevator();
    };

    console.log("\nSECOND CALL \n");
    elevator = column1.requestElevator(10, "down");
    column1.requestFloor(elevator, 3);
    
};


//scenario1();
scenario2();
//scenario3();