// Package models are the glue that connect input and the algorithms
// that solve the problem.
//
// The input must generate a valid model to feed the solver and the solver must
// generate a valid model to output.
//
// It is on the input and solver side to actually construct a valid model, as
// it is not easy to disallow invalid states.
package models

import "errors"

// ErrNoSolution indicates that there is no solution for the puzzle
var ErrNoSolution = errors.New("no solution")

// Action is a user-friendly text indicating the action taken
type Action string

const (
	ActionFillX     = "Fill X"
	ActionFillY     = "Fill Y"
	ActionTransferX = "Transfer to X"
	ActionTransferY = "Transfer to Y"
	ActionEmptyX    = "Empty X"
	ActionEmptyY    = "Empty Y"
)

// State a State indicates the current state of the X and Y Jugs
type State struct {
	X Jug
	Y Jug
}

// Step simply joins a state and the action that got there.
type Step struct {
	// State indicates the step after the action is taken
	State State
	// Action indicates the action that arrived at this state.
	// It can be deduced by checking the previous step.
	Action Action
}

type Solution struct {
	// Steps are the series of steps required to arrive to the solution.
	// Each state in the step list must be consistent with the next one, as in
	// it must follow a possible action.
	// The last step must be a solution to the problem.
	// Meaning that either the X Jug or the Y Jug have z amount of water in
	// them.
	Steps []Step
}

// Jug a Jug carries a certain amount of water.
type Jug struct {
	Capacity int
	Amount   int
}
