package app

// Except for the welcome message, and other possible paragraph messages, no
// message has a new line appended to them.
const (
	welcome = `Welcome to Water Jugs Riddle Solver!

The problem statement is:
There is have an X-gallon and a Y-gallon jug that you can fill from a lake. 
Measure Z gallons of water using only an X-gallon and Y-gallon jug.

`
	requestX = `Please insert the value for the "x" jug, remember it must be positive: `
	requestY = `Please insert the value for the "y" jug, remember it must be positive: `
	requestZ = `Please insert the value for the "z" jug, it must be smaller than either "x" or "y": `

	zSmaller      = "z must be smaller than either x or y"
	zNegative     = "z must be zero or greater"
	xyNotPositive = "both x and z must be positive"

	noSolution = "no solution"
)
