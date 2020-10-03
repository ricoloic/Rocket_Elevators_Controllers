def positive n
    positive = n
    if positive < 0
        positive *= -1
    end
    return positive
end

class Printer
    def doorStage
        print "Elevator #{@ID} has arrived at Floor #{@currentFloor}\n"
        print "The Door Are Open\nThe Door Are Closed\n"
    end

    def floorRequest requestedFloor
        print "...\nFloor Requested :#{requestedFloor}\n\n"
    end

    def changingDirection
        print "Elevator #{@ID} is changing direction"
    end

    def printRequest requestedFloor, direction, bestElevator
        print "...\nRequest from Floor #{requestedFloor} And Going #{direction}\n\n"

        for i in @elevatorList
            print "Elevator #{i.ID} has #{i.points} Points\n"
        end

        print "Elevator #{bestElevator.ID} is sent\n\n"
    end

    def stage
        print "Elevator #{@ID} has the Direction of #{@currentDirection} and the Status of #{@status}. He's at Floor #{@currentFloor}\n"
    end
end

class Elevator < Printer
    attr_accessor :points, :ID, :stopList
    def initialize _id, _floorAmount
        @ID = _id
        @floorAmount = _floorAmount
        @btn = []
        @points = 0
        @stopList = []
        @upBuffer = []
        @downBuffer = []
        @currentDirection
        @previousDirection
        @currentFloor = 0
        @previousFloor = 0
        @door = "closed"
        @status = "IDLE"
        @wordList = ["up", "down", "IDLE"]

        for i in 1..@floorAmount
            @btn.push i + 1
        end
    end

    def listSorting
        if @currentDirection == @wordList[0]
            @stopList.sort!

        elsif @currentDirection == @wordList[1]
            @stopList.sort! {|x, y| y <=> x}

        else
            @stopList.sort!
        end
    end

    def sendRequest requestedFloor
        @stopList.push requestedFloor

        if @currentFloor > requestedFloor
            @status = "MOVING"
            @currentDirection = "down"

        elsif @currentFloor < requestedFloor
            @status = "MOVING"
            @currentDirection = "up"
        end

        self.listSorting
        floorRequest requestedFloor
    end

    def addStop requestFloor, direction
        if direction == @currentDirection and direction == @wordList[0] and requestFloor >= @currentFloor
            @stopList.push requestFloor

        elsif direction == @currentDirection and direction == @wordList[1] and requestFloor <= @currentFloor
            @stopList.push requestFloor

        elsif @status == @wordList[2]
            @stopList.push requestFloor

        elsif direction == @wordList[0]
            @upBuffer.push requestFloor

        elsif direction == @wordList[1]
            @downBuffer.push requestFloor
        end

        self.listSorting
    end

    def pointsUpdate requestFloor, direction
        differenceLastStop = 0
        if @stopList.length != 0
            difLastStop = @stopList[-1] - requestFloor
            differenceLastStop = positive difLastStop
        end

        maxFloorDifference = @floorAmount + 1
        difFloor = @currentFloor - requestFloor
        differenceFloor = positive difFloor

        @points = 0
        if @currentDirection == direction and @status != @wordList[2]
            if requestFloor >= @currentFloor and direction == @wordList[0] or requestFloor <= @currentFloor and direction == @wordList[1]
                @points = differenceFloor
                @points += @stopList.length

            elsif requestFloor < @currentFloor and direction == @wordList[0] or requestFloor > @currentFloor and direction == @wordList[1]
                @points = maxFloorDifference
                @points += differenceLastStop + @stopList.length
            end

        elsif @status == @wordList[2]
            @points = maxFloorDifference
            @points += differenceFloor

        elsif @currentDirection != direction and @status != @wordList[2]
            @points = maxFloorDifference * 2
            @points += differenceLastStop + @stopList.length
        end
    end

    def stopSwitch
        if @downBuffer.length != 0 and @upBuffer.length != 0
            changingDirection
            if @previousDirection == @wordList[0]
                @currentDirection = "down"

                for i in @downBuffer
                    @stopList.push i
                    @downBuffer.delete_at 0
                end

            elsif @previousDirection == @wordList[1]
                @currentDirection = "up"

                for i in @upBuffer
                    @stopList.push i
                    @upBuffer.delete_at 0
                end
            end

        elsif @downBuffer.length != 0 and @upBuffer.length == 0
            changingDirection
            @currentDirection = "down"

            for i in @downBuffer
                @stopList.push i
                @downBuffer.delete_at 0
            end

        elsif @downBuffer.length == 0 and @upBuffer.length != 0
            changingDirection
            @currentDirection = "up"

            for i in @upBuffer
                @stopList.push i
                @upBuffer.delete_at 0
            end

        elsif @downBuffer.length == 0 and @upBuffer.length == 0
            @status = @wordList[2]
            @currentDirection = "stop"
        end

        if @stopList.length != 0
            self.listSorting
            self.run
        end
    end

    def run
        while @stopList.length != 0
            if @stopList.length != 0
                while @currentFloor != @stopList[0]
                    if @stopList[0] < @currentFloor
                        @currentDirection = "down"
                        @previousDirection = @currentDirection
                        @currentFloor -= 1
                        @status = "MOVING"

                    elsif @stopList[0] > @currentFloor
                        @currentDirection = "up"
                        @previousDirection = @currentDirection
                        @currentFloor += 1
                        @status = "MOVING"
                    end

                    if @currentFloor != @previousFloor and @stopList[0] != @currentFloor
                        stage
                        @previousFloor = @currentFloor
                    end
                end

                if @stopList[0] == @currentFloor
                    doorStage
                    @stopList.delete_at 0
                end

            elsif @stopList.length == 0
                slef.stopSwitch
            end
        end

        if @stopList.length == 0
            self.stopSwitch
        end
    end

    def resetPointing
        @points = 0
    end

    def changeValue currentFloor, stopList, downBuffer, upBuffer, currentDirection, status
        @currentFloor = currentFloor
        @stopList = stopList
        @downBuffer = downBuffer
        @upBuffer = upBuffer
        @currentDirection = currentDirection
        @status = status
        self.listSorting
    end
end

class Column < Printer
    attr_accessor :floorAmount, :elevatorList, :elevatorPerColumn
    def initialize _floorAmount, _elevatorPerColumn
        @floorAmount = _floorAmount
        @elevatorPerColumn = _elevatorPerColumn
        @elevatorList = []

        for i in 1..@elevatorPerColumn
            e = Elevator.new i, @floorAmount
            @elevatorList.push e
        end
    end

    def requestElevator requestedFloor, direction
        for i in @elevatorList
            i.pointsUpdate requestedFloor, direction
        end

        swapped = true
        while swapped do
            swapped = false
            0.upto @elevatorList.size-2 do |i|
                if @elevatorList[i].points > elevatorList[i+1].points
                    elevatorList[i], elevatorList[i+1] = elevatorList[i+1], elevatorList[i] # swap values
                    swapped = true
                end
            end
        end

        bestElevator = @elevatorList[0]
        printRequest requestedFloor, direction, bestElevator

        for i in @elevatorList
            i.resetPointing
        end

        bestElevator.addStop requestedFloor, direction
        bestElevator.run
        return bestElevator
    end

    def runElevators
        for i in 0..@elevatorList.length - 1
            @elevatorList[i].listSorting
            @elevatorList[i].run
        end
    end

    def requestFloor elevator ,requestedFloor
        elevator.sendRequest requestedFloor
        elevator.run
    end

    def changeValue elevator, currentFloor, stopList, downBuffer, upBuffer, currentDirection, status
        @elevatorList[elevator].changeValue currentFloor, stopList, downBuffer, upBuffer, currentDirection, status
    end
end

class Scenario
    def initialize _floorAmount, _elevatorPerColumn
        @col = Column.new _floorAmount, _elevatorPerColumn
    end

    def codeboxx n
        if n == 1
            @col.changeValue 0, 2, [], [], [], "stop", "IDLE"
            @col.changeValue 1, 6, [], [], [], "stop", "IDLE"

            elevator = @col.requestElevator 3, "up"
            @col.requestFloor elevator, 7

        elsif n == 2
            @col.changeValue 0, 10, [], [], [], "stop", "IDLE"
            @col.changeValue 1, 3, [], [], [], "stop", "IDLE"

            elevator = @col.requestElevator 1, "up"
            @col.requestFloor elevator, 6

            elevator = @col.requestElevator 3, "up"
            @col.requestFloor elevator, 5

            elevator = @col.requestElevator 9, "down"
            @col.requestFloor elevator, 2

        elsif n == 3
            @col.changeValue 0, 10, [], [], [], "stop", "IDLE"
            @col.changeValue 1, 3, [6], [], [], "up", "MOVING"

            elevator = @col.requestElevator 3, "down"
            @col.requestFloor elevator, 2

            @col.runElevators

            elevator = @col.requestElevator 10, "down"
            @col.requestFloor elevator, 3
        end

        @col.runElevators
    end
end

scenario = Scenario.new 10, 2

scenario.codeboxx 2