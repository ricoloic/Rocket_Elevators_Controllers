## Keyword that could be added in the pseudo code extension

 - PLUS
 - MINUS
 - MULTIPLY
 - DIVIDE
 - NEW
 - MAKE
 - DECREMENT
 - LENGTH
 - GREATER
 - SMALLER
 - OR
 - THAN
 - INDEX
 - FIRST
 - LAST

> - the parenthesis, the brackets and the curly brackets could also used a little highlight

> #### Small remark about the keyword NOT, it only gets highlighted when it is placed in front of the keyword EQUAL

If I think of anything else I'll be updating this file

# hello, I present you the adaptation in code of my algorithm from last week project. the algorithm contain 4 class: the printer, the elevator, the column and the scenario. THE PRINTER. has the name suggest it, will have all the method used for printing thing in the console. the printer is also a parent class. THE ELEVATOR. is a child class, which means that it will be created with all the method and variable of  the parent class.  the elevator has for variable: the id, the number of floor in the column, its current position and direction, its pointing, a stop list, 2 buffer list, and some more. the main part of the algorithm for choosing the best elevator to send is in its points update method. this method will be call whenever someone is at floor and request to go up or down. the method set a certain amount of points to the elevator by looking at its current direction and position and the user current direction and position. for example if we have an elevator with 20 points and one with 13 points we will be choosing the elevator with 13 points. the run method here is called when we add something to its stop list or when needed for the scenario to take place and make it seem like the elevator is always looking if there is something in its stop list and if there is, call the run method of the elevator. and that run method change its current position and direction based on the first index of the stop list. also when the stop list is empty call the stop switch method. The column is also a child of the printer and the column contain the amount of floor in the building, the number of elevator in the column and a list that will be filled fusing a for loop where the number of iteration will be based of the number of elevator contain in that column, and will push new elevator object in the list. the method request elevator is