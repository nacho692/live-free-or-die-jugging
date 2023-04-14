// Package app implements the main flow for the CLI application
// It is interactive and can be automated by piping input into the STDIN
package app

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/nacho692/live-free-or-die-jugging/pkg/models"
)

// SolverFun is a wrapper to simplify the solver interface implementation.
// go does not allow unnamed function types to implement interfaces.
type SolverFun func(state models.State, z int) (models.Solution, error)

// Solve just wraps the internal solver solve.
func (s SolverFun) Solve(state models.State, z int) (models.Solution, error) {
	return s(state, z)
}

// Solver must implement a solution to the water jug riddle.
//
// The solution must be correct if the initial state parameters are correct.
//
// If no solution is found models.ErrNoSolution is expected.
type Solver interface {
	Solve(state models.State, z int) (models.Solution, error)
}

// Configuration is the base configuration for instantiating an interactive App.
type Configuration struct {
	// Output allows configuration for the app output, if nil, stdout is used
	// as default.
	Output io.Writer
	// Input allows configuration for the app input, if nil, stdin is used as
	// default.
	Input io.Reader
	// Silent configures whether the flavour messages to the user are displayed
	// or not.
	// It helps when using the application in an interactive-less way.
	// As it will only output the solution.
	Silent bool
	// Solver must be a valid solver, see Solver for more information.
	Solver Solver
}

// App is an interactive application which guides the user through the water
// jug riddle and then exits.
//
// It requires a working Solver.
type App struct {
	output         writer
	input          reader
	solutionOutput writer
	solver         Solver
}

// New instantiates a new App.
// See Configuration for constraints on the input.
func New(conf Configuration) (App, error) {

	output := conf.Output
	if output == nil {
		output = os.Stdout
	}
	if conf.Silent {
		output = io.Discard
	}

	solutionOutput := conf.Output
	if solutionOutput == nil {
		solutionOutput = os.Stdout
	}

	input := conf.Input
	if input == nil {
		conf.Input = os.Stdin
	}

	if conf.Solver == nil {
		return App{}, errors.New("solver cannot be nil")
	}

	return App{
		input:          reader{bufio.NewReader(conf.Input)},
		output:         writer{output},
		solutionOutput: writer{solutionOutput},
		solver:         conf.Solver,
	}, nil
}

// Run is a blocking operation which expects the x,y,z input on the App input
// and writes the solution to the App output.
//
// The user is expected to write the input and the final input string would look
// like "x\ny\nz\n" as long as x,y and z are valid.
//
// The output varies as messages are written to the user requesting the
// different parameters.
//
// If the App is silent then only the solution output will be written in several
// lines indicating an action taken and the result of that action, until the
// solution is found:
// ACTION  (Fill/TransferTo/Empty; see models.Action)
// (w_x/x, w_y/y)_1 (current amount of water over max capacity for each jug)
//
// Example for 5,4,3:
// Fill Y
// (0/5, 4/4)
// Transfer to X
// (4/5, 0/4)
// Fill Y
// (4/5, 4/4)
// Transfer to X
// (5/5, 3/4)
//
// If no solution exists, "no solution" is written to the output.
func (a *App) Run() error {

	err := a.output.Write(welcome)
	if err != nil {
		return err
	}

	var x, y, z int
	for {
		x, err = a.requestPositiveNumber(requestX)
		if err != nil {
			return fmt.Errorf("requesting positive number: %w", err)
		}

		y, err = a.requestPositiveNumber(requestY)
		if err != nil {
			return fmt.Errorf("requesting positive number: %w", err)
		}

		z, err = a.requestNonNegativeNumber(requestZ)
		valid, err := a.validateParameters(x, y, z)
		if err != nil {
			return fmt.Errorf("validating parameters: %w", err)
		}
		if valid {
			break
		}
	}

	s, err := a.solver.Solve(models.State{
		X: models.Jug{
			Capacity: x,
		},
		Y: models.Jug{
			Capacity: y,
		},
	}, z)

	if err != nil && errors.Is(err, models.ErrNoSolution) {
		return a.solutionOutput.WriteLn(noSolution)
	}
	if err != nil {
		return fmt.Errorf("finding solution: %w", err)
	}

	for _, step := range s.Steps {
		err = a.solutionOutput.Write(
			fmt.Sprintf("%s \n(%d/%d, %d/%d) \n",
				step.Action,
				step.State.X.Amount, step.State.X.Capacity,
				step.State.Y.Amount, step.State.Y.Capacity))
		if err != nil {
			return fmt.Errorf("writing solution to output: %w", err)
		}
	}
	return nil
}

func (a *App) validateParameters(x, y, z int) (bool, error) {
	switch {
	case z > x && z > y:
		return false, a.output.WriteLn(zSmaller)
	case z < 0:
		return false, a.output.WriteLn(zNegative)
	case y <= 0 || x <= 0:
		return false, a.output.WriteLn(xyNotPositive)
	}
	return true, nil
}

func (a *App) requestPositiveNumber(message string) (int, error) {

	for {
		err := a.output.Write(message)
		if err != nil {
			return 0, err
		}

		input, err := a.input.Read()
		if err != nil {
			return 0, fmt.Errorf("requesting int: %w", err)
		}

		number, err := strconv.Atoi(input)
		if number <= 0 || err != nil {
			err = a.output.WriteLn("A positive number was expected")
			if err != nil {
				return 0, err
			}
			continue
		}
		return number, nil
	}
}

func (a *App) requestNonNegativeNumber(message string) (int, error) {

	for {
		err := a.output.Write(message)
		if err != nil {
			return 0, err
		}

		input, err := a.input.Read()
		if err != nil {
			return 0, fmt.Errorf("requesting int: %w", err)
		}

		number, err := strconv.Atoi(input)
		if number < 0 || err != nil {
			err = a.output.WriteLn("A non negative number was expected\n")
			if err != nil {
				return 0, err
			}
			continue
		}
		return number, nil
	}
}
