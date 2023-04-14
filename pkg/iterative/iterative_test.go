package iterative_test

import (
	"github.com/nacho692/live-free-or-die-jugging/pkg/iterative"
	"github.com/nacho692/live-free-or-die-jugging/pkg/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNoSolution(t *testing.T) {

	_, err := iterative.Solve(newBaseState(9, 3), 4)

	assert.ErrorIs(t, err, models.ErrNoSolution)
}

func TestInvalid(t *testing.T) {

	t.Run("x should be positive", func(t *testing.T) {
		_, err := iterative.Solve(newBaseState(-5, 3), 4)
		assert.Error(t, err)
	})

	t.Run("y should be positive", func(t *testing.T) {
		_, err := iterative.Solve(newBaseState(5, -3), 4)
		assert.Error(t, err)
	})

	t.Run("z should be zero or greater", func(t *testing.T) {
		_, err := iterative.Solve(newBaseState(5, 3), -4)
		assert.Error(t, err)
	})

	t.Run("z should be lower than either x or y", func(t *testing.T) {
		_, err := iterative.Solve(newBaseState(5, 3), 10)
		assert.Error(t, err)
	})
}

func TestSolutions(t *testing.T) {

	t.Run("simple solution, should fill X", func(t *testing.T) {
		solution, err := iterative.Solve(newBaseState(5, 3), 4)
		require.NoError(t, err)

		expectedSolution := models.Solution{
			Steps: []models.Step{
				{
					State: models.State{
						X: models.Jug{Capacity: 5, Amount: 5},
						Y: models.Jug{Capacity: 3, Amount: 0},
					},
					Action: models.ActionFillX,
				},
				{
					State: models.State{
						X: models.Jug{Capacity: 5, Amount: 2},
						Y: models.Jug{Capacity: 3, Amount: 3},
					},
					Action: models.ActionTransferY,
				},
				{
					State: models.State{
						X: models.Jug{Capacity: 5, Amount: 2},
						Y: models.Jug{Capacity: 3, Amount: 0},
					},
					Action: models.ActionEmptyY,
				},
				{
					State: models.State{
						X: models.Jug{Capacity: 5, Amount: 0},
						Y: models.Jug{Capacity: 3, Amount: 2},
					},
					Action: models.ActionTransferY,
				},
				{
					State: models.State{
						X: models.Jug{Capacity: 5, Amount: 5},
						Y: models.Jug{Capacity: 3, Amount: 2},
					},
					Action: models.ActionFillX,
				},
				{
					State: models.State{
						X: models.Jug{Capacity: 5, Amount: 4},
						Y: models.Jug{Capacity: 3, Amount: 3},
					},
					Action: models.ActionTransferY,
				},
			},
		}

		assert.Equal(t, expectedSolution, solution)
	})

	t.Run("simple solution, should fill Y", func(t *testing.T) {
		solution, err := iterative.Solve(newBaseState(3, 5), 4)
		require.NoError(t, err)

		expectedSolution := models.Solution{
			Steps: []models.Step{
				{
					State: models.State{
						X: models.Jug{Capacity: 3, Amount: 0},
						Y: models.Jug{Capacity: 5, Amount: 5},
					},
					Action: models.ActionFillY,
				},
				{
					State: models.State{
						X: models.Jug{Capacity: 3, Amount: 3},
						Y: models.Jug{Capacity: 5, Amount: 2},
					},
					Action: models.ActionTransferX,
				},
				{
					State: models.State{
						X: models.Jug{Capacity: 3, Amount: 0},
						Y: models.Jug{Capacity: 5, Amount: 2},
					},
					Action: models.ActionEmptyX,
				},
				{
					State: models.State{
						X: models.Jug{Capacity: 3, Amount: 2},
						Y: models.Jug{Capacity: 5, Amount: 0},
					},
					Action: models.ActionTransferX,
				},
				{
					State: models.State{
						X: models.Jug{Capacity: 3, Amount: 2},
						Y: models.Jug{Capacity: 5, Amount: 5},
					},
					Action: models.ActionFillY,
				},
				{
					State: models.State{
						X: models.Jug{Capacity: 3, Amount: 3},
						Y: models.Jug{Capacity: 5, Amount: 4},
					},
					Action: models.ActionTransferX,
				},
			},
		}

		assert.Equal(t, expectedSolution, solution)
	})

	t.Run("z = 0 should be measurable in 0 steps", func(t *testing.T) {

		s, err := iterative.Solve(newBaseState(5, 3), 0)
		require.NoError(t, err)
		assert.Len(t, s.Steps, 0)
	})
}

func newBaseState(x, y int) models.State {
	return models.State{
		X: models.Jug{
			Capacity: x,
			Amount:   0,
		},
		Y: models.Jug{
			Capacity: y,
			Amount:   0,
		},
	}
}
