// Package app implements the main flow for the CLI application
package app

import (
	"errors"
	"fmt"
	"io"
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
	Output         io.Writer
	Input          io.Reader
	SolutionOutput io.Writer
	Solver         Solver
}

type App struct {
	output         writer
	input          reader
	solutionOutput writer
	solver         Solver
}

func New(conf Configuration) (App, error) {

	if conf.Output == nil {
		conf.Output = io.Discard
	}

	if conf.SolutionOutput == nil {
		conf.SolutionOutput = conf.Output
	}

	if conf.Input == nil {
		return App{}, errors.New("configuration Input stream must be set")
	}

	if conf.Solver == nil {
		return App{}, errors.New("solver cannot be nil")
	}

	return App{
		input:          reader{conf.Input},
		output:         writer{conf.Output},
		solutionOutput: writer{conf.SolutionOutput},
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
		return a.output.WriteLn(noSolution)
	}
	if err != nil {
		return fmt.Errorf("finding solution: %w", err)
	}

	for _, step := range s.Steps {
		err = a.solutionOutput.Write(
			fmt.Sprintf("(%d, %d) \t %s \n",
				step.State.X.Amount, step.State.Y.Amount, step.Action))
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
