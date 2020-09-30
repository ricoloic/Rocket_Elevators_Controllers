from operator import attrgetter
import time

def returnPositive(n):
    if n < 0:
        n *= -1

    return n

def wait(t):
    t = 0.1
    time.sleep(t)


class Printer:
    def doorStage(self):
        print("The Door Are Open")
        wait(3)
        print("The Door Are Closed")
        print("\n")

    
    def floorRequest(self, requestedFloor):
            print("...")
            wait(1)
            print("Floor Requested :{}\n".format(requestedFloor))


    def changingDirection(self):
        print(" Elevator {} is changing direction".format(self.ID))


    def requestedUser(self, requestedFloor, direction):
        print("...")
        wait(2)
        print("Request from Floor {} And Going {}\n".format(requestedFloor, direction))


    def elevatorChosen(self, bestElevator):
        print("Elevator {} is sent")


    def pointingStage(self):
        for i in range(len(self.elevatorList)):
            print("Elevator {} has {} Points".format(self.elevatorList[i].ID, self.elevatorList[i].points))
        print("\n")


    def stage(self):
        print("Elevator {} has the Direction of {} and the Status of {}. He's at Floor {}.".format(self.ID, self.currentDirection, self.status, self.currentFloor))



class Elevator(Printer):
    def __init__(self, _id, _floorAmount):
        self.ID = _id + 1
        self.floorAmount = _floorAmount
        self.btn = []
        self.points = 0
        self.StopList = []
        self.UpBuffer = []
        self.DownBuffer =  []
        self.currentDirection = None
        self.previousDirection = None
        self.currentFloor = 0
        self.previousFloor = 0
        self.door = "closed"    # false = closed / true = open
        self.status = "IDLE"    # IDLE / maintenance / moving

        for i in range(self.floorAmount):
            self.btn.append(i + 1)


    def doorState(self):
        self.door = "open"
        self.door = "closed"
        self.doorStage()


    def checkIn(self, n):
        inList = True

        for i in range(len(self.StopList)):
            if n == self.StopList[i]:
                inList = False
        return inList


    def listSorting(self):
        if self.currentDirection == "Up":
            self.StopList.sort()

        elif self.currentDirection == "Down":
            self.StopList.sort(reverse = True)

        else:
            self.StopList.sort()


    def sendRequest(self, requestedFloor):
        if self.checkIn(requestedFloor):
            self.StopList.append(requestedFloor)

            if self.currentFloor > requestedFloor:
                self.status = "MOVING"
                self.currentDirection = "Down"
            
            elif self.currentFloor < requestedFloor:
                self.status = "MOVING"
                self.currentDirection = "Up"

            self.listSorting()
            self.floorRequest(requestedFloor)



    def addStop(self, requestFloor, direction):
        if self.checkIn(requestFloor):
            if direction == self.currentDirection and direction == "Up" and requestFloor >= self.currentFloor:
                self.StopList.append(requestFloor)

            elif direction == self.currentDirection and direction == "Down" and requestFloor <= self.currentFloor:
                self.StopList.append(requestFloor)

            elif self.status == "IDLE":
                self.StopList.append(requestFloor)

            elif direction == "Up":
                self.UpBuffer.append(requestFloor)

            elif direction == "Down":
                self.DownBuffer.append(requestFloor)

            self.listSorting()


    def pointsUpdate(self, requestFloor, direction):
        differenceLastStop = 0
        if self.status != "IDLE":
            difLastStop = self.StopList[-1] - requestFloor
            differenceLastStop = returnPositive(difLastStop)

        maxFloorDifference = self.floorAmount + 1

        difFloor = self.currentFloor - requestFloor
        differenceFloor = returnPositive(difFloor)

        self.points = 0

        if self.currentDirection == direction and self.status != "IDLE":
            if requestFloor >= self.currentFloor and direction == "Up" or requestFloor <= self.currentFloor and direction == "Down":
                self.points = differenceFloor
                self.points += len(self.StopList)
            
            elif requestFloor < self.currentFloor and direction == "Up" or requestFloor > self.currentFloor and direction == "Down":
                self.points = maxFloorDifference
                self.points += differenceLastStop + len(self.StopList)

        elif self.status == "IDLE":
            self.points = maxFloorDifference
            self.points += differenceFloor

        elif self.currentDirection != direction and self.status != "IDLE":
            self.points = maxFloorDifference * 2
            self.points += differenceLastStop + len(self.StopList)


    def StopSwitch(self):
        if len(self.DownBuffer) != 0 and len(self.UpBuffer) != 0:
            self.changingDirection()
            if self.previousDirection == "Up":
                self.currentDirection = "Down"
                for i in self.DownBuffer:
                    self.StopList.append(i)
                    del self.DownBuffer[0]
            
            elif self.previousDirection == "Down":
                self.currentDirection = "Up"
                for i in self.UpBuffer:
                    self.StopList.append(i)
                    del self.UpBuffer[0]

        elif len(self.DownBuffer) != 0 and len(self.UpBuffer) == 0:
            self.changingDirection()
            self.currentDirection = "Down"
            for i in self.DownBuffer:
                self.StopList.append(i)
                del self.DownBuffer[0]

        elif len(self.DownBuffer) == 0 and len(self.UpBuffer) != 0:
            self.changingDirection()
            self.currentDirection = "Up"
            for i in self.UpBuffer:
                self.StopList.append(i)
                del self.UpBuffer[0]

        elif len(self.DownBuffer) == 0 and len(self.UpBuffer) == 0:
            self.status = "IDLE"
            self.currentDirection = "Stop"

        if len(self.StopList) != 0:
            self.listSorting()
            self.run()


    def run(self):
        while len(self.StopList) != 0:
            if len(self.StopList) != 0:
                while self.currentFloor != self.StopList[0]:
                    if self.StopList[0] < self.currentFloor:
                        self.currentDirection = "Down"
                        self.previousDirection = self.currentDirection
                        self.currentFloor -= 1
                        self.status = "MOVING"

                    elif self.StopList[0] > self.currentFloor:
                        self.currentDirection = "Up"
                        self.previousDirection = self.currentDirection
                        self.currentFloor += 1
                        self.status = "MOVING"
                    
                    if self.currentFloor != self.previousFloor:
                        self.stage()
                        self.previousFloor = self.currentFloor
                
                if self.StopList[0] == self.currentFloor:
                    self.doorState()
                    del self.StopList[0]

            elif len(self.StopList) == 0:
                self.StopSwitch()

        if len(self.StopList) == 0:
            self.StopSwitch()


    def resetPointing(self):
        self.points = 0



class Column(Printer):
    def __init__(self, Elevator, _floorAmount, _elevatorPerColumn):
        self.floorAmount = _floorAmount
        self.elevatorPerColumn = _elevatorPerColumn
        self.elevatorList = []

        for i in range(self.elevatorPerColumn):
            e = Elevator(i, self.floorAmount)
            self.elevatorList.append(e)


    def requestElevator(self, requestedFloor, direction):
        for i in range(len(self.elevatorList)):
            self.elevatorList[i].pointsUpdate(requestedFloor, direction)

        bestElevator = min(self.elevatorList, key = attrgetter('points'))
        self.requestedUser(requestedFloor, direction)

        self.pointingStage()
        self.elevatorChosen(bestElevator)
        bestElevator.addStop(requestedFloor, direction)        
        bestElevator.run()
        return bestElevator


    def requestFloor(self, Elevator, requestedFloor):
        Elevator.sendRequest(requestedFloor)
        Elevator.run()

    
    def changeValue(self, elevator, currentFloor, StopList, DownBuffer, UpBuffer, currentDirection, status):
        self.elevatorList[elevator].currentFloor = currentFloor
        self.elevatorList[elevator].StopList = StopList
        self.elevatorList[elevator].DownBuffer = DownBuffer
        self.elevatorList[elevator].UpBuffer = UpBuffer
        self.elevatorList[elevator].currentDirection = currentDirection
        self.elevatorList[elevator].status = status
        self.elevatorList[elevator].listSorting()




# test
col = Column(Elevator, 10, 2)

# column.changeValue(elevator, currentFloor, StopList, DownBuffer, UpBuffer, currentDirection, status)

def scenario1 ():
    col.changeValue(0, 2, [], [], [], "Stop", "IDLE")
    col.changeValue(1, 6, [], [], [], "Stop", "IDLE")

    elevator = col.requestElevator(3, "Up")
    col.requestFloor(elevator, 7)

def scenario2 ():
    col.changeValue(0, 10, [], [], [], "Stop", "IDLE")
    col.changeValue(1, 3, [], [], [], "Stop", "IDLE")

    elevator = col.requestElevator(1, "Up")
    col.requestFloor(elevator, 6)

    elevator = col.requestElevator(3, "Up")
    col.requestFloor(elevator, 5)

    elevator = col.requestElevator(9, "Down")
    col.requestFloor(elevator, 2)

def scenario3 ():
    col.changeValue(0, 10, [], [], [], "Stop", "IDLE")
    col.changeValue(1, 3, [6], [], [], "Up", "MOVING")

    elevator = col.requestElevator(3, "Down")
    col.requestFloor(elevator, 2)

    for elevator in col.elevatorList:
        elevator.listSorting()
        elevator.run()

    elevator = col.requestElevator(10, "Down")
    col.requestFloor(elevator, 3)

def customScenario ():
    col.changeValue(0, 9, [7, 6, 5, 3], [], [4, 10], "Down", "MOVING")
    col.changeValue(1, 5, [6, 8, 10], [7, 3], [2, 5], "Up", "MOVING")

    elevator = col.requestElevator(4, "Down")
    col.requestFloor(elevator, 10)

# scenario1()
scenario2()
# scenario3()

# customScenario()

for elevator in col.elevatorList:
    elevator.listSorting()
    elevator.run()