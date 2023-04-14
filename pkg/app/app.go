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

type SolverFun func(state models.State, z int) (models.Solution, error)

func (s SolverFun) Solve(state models.State, z int) (models.Solution, error) {
	return s(state, z)
}

type Solver interface {
	Solve(state models.State, z int) (models.Solution, error)
}

type Configuration struct {
	Output io.Writer
	Input  io.Reader
	Silent bool
	Solver Solver
}

type App struct {
	output         writer
	input          reader
	solutionOutput writer
	solver         Solver
}

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
		}
		return number, nil
	}
}
