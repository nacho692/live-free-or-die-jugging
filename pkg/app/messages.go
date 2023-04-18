package app

// Except for the welcome message, and other possible paragraph messages, no
// message has a new line appended to them.
const (
	welcome = `Welcome to Water Jugs Riddle Solver!

The problem statement is:
There is have an X-gallon and a Y-gallon jug that you can fill from a lake. 
Measure Z gallons of water using only an X-gallon and Y-gallon jug.

`
	//here all the specific const could be just templates and send X or Y as a param to avoid copypaste
	requestX = `Insert the value for the "x" jug, remember it must be positive: `
	requestY = `Insert the value for the "y" jug, remember it must be positive: `
	requestZ = `Insert the value for the "z" goal, it must be smaller than either "x" or "y": `

	zSmaller      = "z must be smaller than either x or y"
	zNegative     = "z must be zero or greater"
	xyNotPositive = "both x and z must be positive"

	noSolution = "no solution"
)
