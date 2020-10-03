function returnPositive(n) {
    let N = n;
    if (N < 0) {
        N *= -1;
    };

    return N;
};

class Printer {
    doorStage() {
        console.log("The Door Are Open");
        console.log("The Door Are Closed\n");
    };

    floorRequest(requestedFloor) {
        console.log("...");
        console.log(`Floor Requested :${requestedFloor}\n`);
    };

    changingDirection() {
        console.log(`Elevator ${this.ID} is changing direction`);
    };

    printRequest(requestedFloor, direction, bestElevator) {
        console.log("...");
        console.log(`Request from Floor ${requestedFloor} And Going ${direction}\n`);

        for (let i = 0; i < this.elevatorList.length; i++) {
            console.log(`Elevator ${this.elevatorList[i].ID} has ${this.elevatorList[i].points} Points`);
        };

        console.log(`Elevator ${bestElevator.ID} is sent\n`);
    };

    stage() {
        console.log(`Elevator ${this.ID} has the Direction of ${this.currentDirection} and the Status of ${this.status}. He's at Floor ${this.currentFloor}`);
    };

};

class Elevator extends Printer {
    constructor(_id, _floorAmount) {
        super();
        this.ID = _id + 1;
        this.floorAmount = _floorAmount;
        this.btn = [];
        this.points;
        this.stopList = [];
        this.upBuffer = [];
        this.downBuffer = [];
        this.currentDirection;
        this.previousDirection;
        this.currentFloor = 0;
        this.previousFloor = 0;
        this.door = "closed";
        this.status = "IDLE";

        for (let i = 1; i <= this.floorAmount; i++) {
            this.btn.push(i);
        };
    };

    checkIn(n) {
        let inList = true;

        for (let i = 0; i < this.stopList.length; i++) {
            if (n == this.stopList[i]) {
                inList = false;
            };
        };

        return inList;
    };

    doorState() {
        this.door = "open";
        this.door = "closed";
        this.doorStage();
    };

    listSorting() {
        if (this.currentDirection == "up") {
            this.stopList.sort((a, b) => {
                return a - b
            });

        } else if (this.currentDirection == "down") {
            this.stopList.sort((a, b) => {
                return b - a
            });

        } else {
            this.stopList.sort((a, b) => {
                return a - b
            });
        };
    };

    sendRequest(requestedFloor) {
        if (this.checkIn(requestedFloor)) {
            this.stopList.push(requestedFloor);

            if (this.currentFloor > requestedFloor) {
                this.status = "MOVING";
                this.currentDirection = "down";

            } else if (this.currentFloor < requestedFloor) {
                this.status = "MOVING";
                this.currentDirection = "up";
            };

            this.listSorting();
            this.floorRequest(requestedFloor);
        };
    };

    addStop(requestFloor, direction) {
        if (this.checkIn(requestFloor)) {
            if (direction == this.currentDirection && direction == "up" && requestFloor >= this.currentFloor) {
                this.stopList.push(requestFloor);

            } else if (direction == this.currentDirection && direction == "down" && requestFloor <= this.currentFloor) {
                this.stopList.push(requestFloor);

            } else if (this.status == "IDLE") {
                this.stopList.push(requestFloor);

            } else if (direction == "up") {
                this.upBuffer.push(requestFloor);

            } else if (direction == "down") {
                this.downBuffer.push(requestFloor);
            };

            this.listSorting();
        };
    };

    pointsUpdate(requestFloor, direction) {
        let differenceLastStop = 0;
        if (this.status != "IDLE") {
            let difLastStop = this.stopList[this.stopList.length - 1] - requestFloor;
            differenceLastStop = returnPositive(difLastStop);
        };

        let maxFloorDifference = this.floorAmount + 1;

        let difFloor = this.currentFloor - requestFloor;
        let differenceFloor = returnPositive(difFloor);

        this.points = 0;

        if (this.currentDirection == direction && this.status != "IDLE") {
            if (requestFloor >= this.currentFloor && direction == "up" || requestFloor <= this.currentFloor && direction == "down") {
                this.points = differenceFloor;
                this.points += this.stopList.length;

            } else if (requestFloor < this.currentFloor && direction == "up" || requestFloor > this.currentFloor && direction == "down") {
                this.points = maxFloorDifference;
                this.points += differenceLastStop + this.stopList.length;
            };

        } else if (this.status == "IDLE") {
            this.points = maxFloorDifference;
            this.points += differenceFloor;

        } else if (this.currentDirection != direction && this.status != "IDLE") {
            this.points = maxFloorDifference * 2;
            this.points += differenceLastStop + this.stopList.length;
        };
    };

    stopSwitch() {
        if (this.downBuffer.length != 0 && this.upBuffer.length != 0) {
            this.changingDirection();
            if (this.previousDirection == "up") {
                this.currentDirection = "down";
                for (let i = 0; i < this.downBuffer.length; i++) {
                    this.stopList.push(this.downBuffer[0]);
                    this.downBuffer.splice(0, 1);
                };

            } else if (this.previousDirection == "down") {
                this.currentDirection = "up";
                for (let i = 0; i < this.upBuffer.length; i++) {
                    this.stopList.push(this.upBuffer[0]);
                    this.upBuffer.splice(0, 1);
                };
            };

        } else if (this.downBuffer.length != 0 && this.upBuffer.length == 0) {
            this.changingDirection();
            this.currentDirection = "down";
            for (let i = 0; i < this.downBuffer.length; i++) {
                this.stopList.push(this.downBuffer[0]);
                this.downBuffer.splice(0, 1);
            };

        } else if (this.downBuffer.length == 0 && this.upBuffer.length != 0) {
            this.changingDirection();
            this.currentDirection = "up";
            for (let i = 0; i < this.upBuffer.length; i++) {
                this.stopList.push(this.upBuffer[0]);
                this.upBuffer.splice(0, 1);
            };

        } else if (this.downBuffer.length == 0 && this.upBuffer.length == 0) {
            this.status = "IDLE";
            this.currentDirection = "stop";
        };

        if (this.stopList.length != 0) {
            this.listSorting();
            this.run();
        };
    };

    run() {
        while (this.stopList.length != 0) {
            if (this.stopList.length != 0) {
                while (this.currentFloor != this.stopList[0]) {
                    if (this.stopList[0] < this.currentFloor) {
                        this.currentDirection = "down";
                        this.previousDirection = this.currentDirection;
                        this.currentFloor -= 1;
                        this.status = "MOVING";

                    } else if (this.stopList[0] > this.currentFloor) {
                        this.currentDirection = "up";
                        this.previousDirection = this.currentDirection;
                        this.currentFloor += 1;
                        this.status = "MOVING";
                    };

                    if (this.currentFloor != this.previousFloor) {
                        this.stage();
                        this.previousFloor = this.currentFloor;
                    };
                };

                if (this.stopList[0] == this.currentFloor) {
                    this.doorState();
                    this.stopList.splice(0, 1);
                };

            } else if (this.stopList.length == 0) {
                this.stopSwitch();
            };
        };

        if (this.stopList.length == 0) {
            this.stopSwitch();
        };
    };

    changeValue(currentFloor, stopList, downBuffer, upBuffer, currentDirection, status) {
        this.currentFloor = currentFloor;
        this.stopList = stopList;
        this.downBuffer = downBuffer;
        this.upBuffer = upBuffer;
        this.currentDirection = currentDirection;
        this.status = status;
        this.listSorting();
    };
};

class Column extends Printer {
    constructor(_floorAmount, _elevatorPerColumn) {
        super();
        this.floorAmount = _floorAmount;
        this.elevatorPerColumn = _elevatorPerColumn;
        this.elevatorList = [];

        for (let i = 0; i < this.elevatorPerColumn; i++) {
            let e = new Elevator(i, this.floorAmount);
            this.elevatorList.push(e);
        };
    };

    runElevators() {
        for (let i = 0; i < this.elevatorList.length; i++) {
            this.elevatorList[i].listSorting();
            this.elevatorList[i].run();
        };
    };

    requestElevator(requestFloor, direction) {
        for (let i = 0; i < this.elevatorList.length; i++) {
            this.elevatorList[i].pointsUpdate(requestFloor, direction);
        };

        this.elevatorList.sort((a, b) => {
            return parseFloat(a.points) - parseFloat(b.points);
        });
        const bestElevator = this.elevatorList[0];

        this.printRequest(requestFloor, direction, bestElevator);
        bestElevator.addStop(requestFloor, direction);
        bestElevator.run();
        return bestElevator;
    };

    requestFloor(elevator, requestedFloor) {
        elevator.sendRequest(requestedFloor);
        elevator.run();
    };

    changeValue(elevator, currentFloor, stopList, downBuffer, upBuffer, currentDirection, status) {
        this.elevatorList[elevator].changeValue(currentFloor, stopList, downBuffer, upBuffer, currentDirection, status);
    };
};

class Scenario {
    constructor(_nbFloor, _nbElevator) {
        this.col = new Column(_nbFloor, _nbElevator);
    };

    codeboxx(i) {
        if (i == 1) {
            this.col.changeValue(0, 2, [], [], [], "stop", "IDLE");
            this.col.changeValue(1, 6, [], [], [], "stop", "IDLE");

            let elevator = this.col.requestElevator(3, "up");
            this.col.requestFloor(elevator, 7);

        } else if (i == 2) {
            this.col.changeValue(0, 10, [], [], [], "stop", "IDLE");
            this.col.changeValue(1, 3, [], [], [], "stop", "IDLE");

            let elevator = this.col.requestElevator(1, "up");
            this.col.requestFloor(elevator, 6);

            elevator = this.col.requestElevator(3, "up");
            this.col.requestFloor(elevator, 5);

            elevator = this.col.requestElevator(9, "down");
            this.col.requestFloor(elevator, 2);

        } else if (i == 3) {
            this.col.changeValue(0, 10, [], [], [], "stop", "IDLE");
            this.col.changeValue(1, 3, [6], [], [], "up", "MOVING");

            let elevator = this.col.requestElevator(3, "down");
            this.col.requestFloor(elevator, 2);

            this.col.runElevators();

            elevator = this.col.requestElevator(10, "down");
            this.col.requestFloor(elevator, 3);
        };

        this.col.runElevators();
    };

    custom(i) {
        if (i == 1) {
            this.col.changeValue(0, 9, [7, 6, 5, 3], [], [4, 10], "down", "MOVING");
            this.col.changeValue(1, 5, [6, 8, 10], [7, 3], [2, 5], "up", "MOVING");

            let elevator = this.col.requestElevator(4, "down");
            this.col.requestFloor(elevator, 10);
        };

        this.col.runElevators();
    };
};

let scenario = new Scenario(10, 2);
scenario.codeboxx(2);