# An **INSANE** algorithm, a bit like Me
Hi, So now that we are in Odyssey what is there to do ? **WEEK 3, 4, 5  - ELEVATORS CONTROLLERS ALGORITHMS** An algorithm what for ? Not again with the elevator theme ? All right, let's see what we're dealing with here. Mmmm, so you're telling me I have to control all the elevators in a building, oh you shouldn't have. **Moua ha ha**, Just kidding. This was a really intensive week for my brain. It will have been the biggest brainteaser my brain has ever had to solve. A lightweight algorithm you say, well that's not what my algorithm is, but if we're talking about its ability to solve an elevator request problem, let me tell you that it is more than capable of solving the scenarios given by the Codeboxx team.

## Language

 - [Go](https://golang.org/)
 - [C#](https://docs.microsoft.com/en-us/learn/paths/csharp-first-steps/)
 - [JavaScript](https://www.javascript.com/)
 - [Ruby](https://www.ruby-lang.org/)
 - [Python](https://www.python.org/)
 
# [Video Week 2](https://youtu.be/VQ67JKKJKUE/)
# [Video Week 3](https://youtu.be/XX_-aWSohq0/)
![PRINTING THE PROGRAM](https://i.gyazo.com/c2cb644a8f3683df7d039b70cec7ca30.gif)

# New Addition "THE SO CALLED PRINTER"
take a good look at that baby

## A brief description

 hello, I present you the adaptation in code of my algorithm from last week project. the algorithm contain 4 class: the printer, the elevator, the column and the scenario. THE PRINTER. has the name suggest it, will have all the method used for printing thing in the console. the printer is also a parent class. THE ELEVATOR. is a child class, which means that it will be created with all the method and variable of  the parent class.  the elevator has for variable: the id, the number of floor in the column, its current position and direction, its pointing, a stop list, 2 buffer list, and some more. the main part of the algorithm for choosing the best elevator to send is in its points update method. this method will be call whenever someone is at floor and request to go up or down. the method set a certain amount of points to the elevator by looking at its current direction and position and the user current direction and position. for example if we have an elevator with 20 points and one with 13 points we will be choosing the elevator with 13 points. the run method here is called when we add something to its stop list or when needed for the scenario to take place and make it seem like the elevator is always looking if there is something in its stop list and if there is, call the run method of the elevator. and that run method change its current position and direction based on the first index of the stop list. also when the stop list is empty call the stop switch method. The column is also a child of the printer and the column contain the amount of floor in the building, the number of elevator in the column and a list that will be filled fusing a for loop where the number of iteration will be based of the number of elevator contain in that column, and will push new elevator object in the list. the method request elevator is used for calling the points update method in all the elevator in the column using a for loop then it will sort the list of elevators based on there points from smallest to largest. it will also call some method from the best elevator which is set to the first index of the list of elevators then it will return the best elevator. the method request floor is used to add a stop to the stop list by calling some method from the best elevator object past as parameter. THE SCENARIO is a class created for the scenario used if needed you can add some more scenario, but there is 4 at the moment which are in a method named codeboxx or custom. if you want to run a scenario you just have to create a scenario object then call its method codeboxx with the wanted scenario from one to tree.
