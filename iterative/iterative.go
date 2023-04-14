// Package iterative implements an iterative solution to the water jug puzzle.
//
// Let's take x = 5, y = 3, z = 4 as an example and consider the tuple
// (w_x, w_y) as the current gallons in each jug.
//
// We start by willing one of them up.
//
// Fill x -> (5, 0)
//
// Now we can either fill y, empty x or transfer to y.
// Filling the second one is pretty useless unless we are measuring exactly y.
// Emptying x is useless, we are going back.
//
// Transfer to y -> (2, 3)
//
// Now we can either fill x, transfer to y, empty x or empty y.
// Filling x is useless, we could have gotten there easier, transferring to y
// does nothing, emptying x is the same as filling y at the start.
//
// Empty y -> (2, 0)
//
// Using the same deductions, transferring to y is the reasonable choice.
//
// Transfer to y -> (0, 2)
// Fill x -> (5, 2)
// Transfer to y -> (4, 3)
//
// We got there!
//
// # The same idea can be used by filling first the y jug and transferring to x
//
// (0, 3) -> (3, 0) -> (3, 3) -> (5, 1) -> (0, 1) -> (1, 0) -> (1, 3) -> (4, 0)
//
// Assuming X is the jug we are filling first the algorithm would look like:
// Step 1: Fill X if empty
// Step 2: Transfer to Y
// Step 3: If Y is full empty it
// Step 5: Repeat from Step 1 until winning condition is met
//
// There are problems with no solutions, for example
// x=9, y=3, z=5
//
// (9, 0) -> (6, 3) -> (6, 0) -> (3, 3) -> (3, 0) -> (0, 3) -> (0, 0) -> (9, 0)
// (0, 3) -> (3, 3) -> (6, 0) -> (6, 3) -> (9, 0) -> (9, 3) -> (0, 3)
//
// Given that the amount of states we are traversing through is limited, one
// easy way to check for problems without solutions is checking if we are
// cycling.
package iterative

import "errors"

type Solution struct {
	Steps  []State
	Action []string
}

type Jug struct {
	Capacity uint
	Amount   uint
}

type State struct {
	X Jug
	Y Jug
}

func Solve(x, y, z uint) (Solution, error) {
	if z < x && z < y {
		return Solution{}, errors.New("z must be smaller than either x or y")
	}
	if y == 0 || x == 0 {
		return Solution{}, errors.New("both y and x must be positive")
	}
	s := Solution{}
	state := State{
		X: Jug{
			Capacity: x,
			Amount:   0,
		},
		Y: Jug{
			Capacity: y,
			Amount:   0,
		},
	}
	s.Steps = append(s.Steps, state)
	toTransfer := uint(0)
	for state.X.Amount != z && state.Y.Amount != z {

		if state.Y.Amount == state.Y.Capacity {
			state.Y.Amount = 0
			s.Steps = append(s.Steps, state)
			s.Action = append(s.Action, "Empty Y")
		}

		if state.X.Amount == 0 {
			state.X.Amount = state.X.Capacity
			s.Steps = append(s.Steps, state)
			s.Action = append(s.Action, "Fill X")
		}

		toTransfer = min(state.X.Amount, state.Y.Capacity-state.Y.Amount)
		state.X.Amount -= toTransfer
		state.Y.Amount += toTransfer
		s.Steps = append(s.Steps, state)
		s.Action = append(s.Action, "Transfer to Y")
	}

	return s, nil
}

func min(x, y uint) uint {
	if x < y {
		return x
	}
	return y
}
