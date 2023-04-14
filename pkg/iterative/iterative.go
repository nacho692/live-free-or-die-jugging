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
// The proof of the algorithm is left as an exercise for the reader.
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

import (
	"errors"
	"github.com/nacho692/live-free-or-die-jugging/pkg/models"
)

type action string

const (
	transferTo action = "transfer_to"
	fillFrom   action = "fill"
	emptyTo    action = "empty"
)

type step func(act action, from, to models.Jug)

// Solve solves the water jugs riddle iteratively.
//
// An error ErrNoSolution is returned if no solution exists.
func Solve(baseState models.State, z int) (models.Solution, error) {

	x := baseState.X.Capacity
	y := baseState.Y.Capacity
	if z > x && z > y {
		return models.Solution{}, errors.New("z must be smaller than either x or y")
	}
	if z < 0 {
		return models.Solution{}, errors.New("z must be zero or greater")
	}
	if y <= 0 || x <= 0 {
		return models.Solution{}, errors.New("both x and z must be positive")
	}

	// We derive two solutions, first pouring from X to Y, secondly from Y to X
	// we keep the minimum of both
	s1 := models.Solution{}
	err := solveFromTo(
		baseState.X,
		baseState.Y,
		// The callback adds a solution step, knowing that the From Jug is X
		// and the To Jug is Y.
		func(act action, from, to models.Jug) {
			s := models.Step{
				State: models.State{
					X: from,
					Y: to,
				},
			}
			switch act {
			case fillFrom:
				s.Action = models.ActionFillX
			case emptyTo:
				s.Action = models.ActionEmptyY
			case transferTo:
				s.Action = models.ActionTransferY
			}
			s1.Steps = append(s1.Steps, s)
		},
		z)

	if err != nil {
		return models.Solution{}, err
	}

	s2 := models.Solution{}
	err = solveFromTo(
		baseState.Y,
		baseState.X,
		func(act action, from, to models.Jug) {
			s := models.Step{
				State: models.State{
					X: to,
					Y: from,
				},
			}
			switch act {
			case fillFrom:
				s.Action = models.ActionFillY
			case emptyTo:
				s.Action = models.ActionEmptyX
			case transferTo:
				s.Action = models.ActionTransferX
			}
			s2.Steps = append(s2.Steps, s)
		},
		z)
	if err != nil {
		return models.Solution{}, err
	}

	if len(s1.Steps) < len(s2.Steps) {
		return s1, nil
	}
	return s2, nil
}

// solveFromTo helps abstract the algorithm from the expected Solution format.
// It solves the water jugs riddle by transfering water from the "from" Jug to
// the "to" Jug.
//
// The idea is to call this method twice, inverting the from and to values.
//
// It receives a "step" method which is basically a callback whenever a new step
// is generated.
// This callback allows and helps formatting the Solution correctly avoiding too
// much code repetition.
func solveFromTo(
	from models.Jug, to models.Jug,
	newStep step,
	z int) error {

	// If z is 0 we already have a solution, and that is doing nothing
	if z == 0 {
		return nil
	}

	toTransfer := 0

	type tuple struct {
		from, to models.Jug
	}

	// We start by filling the from.
	// This allows checking for the winning condition, the only time this action
	// "wins" is the first time we fill the from jug.
	from.Amount = from.Capacity
	newStep(fillFrom, from, to)

	visitedTuples := map[tuple]bool{}
	for from.Amount != z && to.Amount != z && !visitedTuples[tuple{from: from, to: to}] {

		visitedTuples[tuple{from: from, to: to}] = true

		if to.Amount == to.Capacity {
			to.Amount = 0
			newStep(emptyTo, from, to)
		}

		if from.Amount == 0 {
			from.Amount = from.Capacity
			newStep(fillFrom, from, to)
		}

		toTransfer = min(from.Amount, to.Capacity-to.Amount)
		from.Amount -= toTransfer
		to.Amount += toTransfer
		newStep(transferTo, from, to)
	}

	if visitedTuples[tuple{from: from, to: to}] {
		return models.ErrNoSolution
	}
	return nil
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
