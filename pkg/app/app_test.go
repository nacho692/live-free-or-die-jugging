package app_test

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/nacho692/live-free-or-die-jugging/pkg/app"
	"github.com/nacho692/live-free-or-die-jugging/pkg/models"
)

// We are only testing the solution output and the inputs
// The flavour text and hints can change at anytime, and, in my opinion
// do not make sense to test.
func TestRun(t *testing.T) {

	t.Run("happy path with x=3, y=2, z=1", func(t *testing.T) {

		expected := func(state models.State, z int) (models.Solution, error) {
			assert.Equal(t, models.State{
				X: models.Jug{Capacity: 3},
				Y: models.Jug{Capacity: 2},
			}, state)
			assert.Equal(t, 1, z)
			return models.Solution{
				Steps: []models.Step{
					{
						State: models.State{
							X: models.Jug{Capacity: 3, Amount: 3},
							Y: models.Jug{Capacity: 2, Amount: 0},
						},
						Action: models.ActionFillX,
					},
					{
						State: models.State{
							X: models.Jug{Capacity: 3, Amount: 1},
							Y: models.Jug{Capacity: 2, Amount: 2},
						},
						Action: models.ActionTransferY,
					},
				},
			}, nil
		}
		input := "3\n2\n1\n"
		output := &bytes.Buffer{}
		a, err := app.New(app.Configuration{
			Input:  bytes.NewReader([]byte(input)),
			Output: output,
			Silent: true,
			Solver: app.SolverFun(expected),
		})
		require.NoError(t, err)

		err = a.Run()
		require.NoError(t, err)

		assert.Equal(t, output.String(),
			"Fill X \n(3/3, 0/2) \n"+
				"Transfer to Y \n(1/3, 2/2) \n")
	})

	t.Run("no solution x=3, y=9, z=4", func(t *testing.T) {

		expected := func(state models.State, z int) (models.Solution, error) {
			assert.Equal(t, models.State{
				X: models.Jug{Capacity: 3},
				Y: models.Jug{Capacity: 9},
			}, state)
			assert.Equal(t, 4, z)
			return models.Solution{}, models.ErrNoSolution
		}
		input := "3\n9\n4\n"
		output := &bytes.Buffer{}
		a, err := app.New(app.Configuration{
			Input:  bytes.NewReader([]byte(input)),
			Output: output,
			Silent: true,
			Solver: app.SolverFun(expected),
		})
		require.NoError(t, err)

		err = a.Run()
		require.NoError(t, err)

		assert.Equal(t, output.String(),
			"no solution\n")
	})

}
