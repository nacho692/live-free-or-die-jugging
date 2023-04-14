package iterative

import "testing"

func TestNoSolution(t *testing.T) {

	_, err := Solve(newBaseState(9, 3), 4)
	if err == nil {
		t.Error("error expected")
	}
	if err != ErrNoSolution {
		t.Error("no solution error expected")
	}
}

func TestInvalid(t *testing.T) {

	t.Run("x should be positive", func(t *testing.T) {
		_, err := Solve(newBaseState(-5, 3), 4)
		if err == nil {
			t.Errorf("error expected")
		}
	})

	t.Run("y should be positive", func(t *testing.T) {
		_, err := Solve(newBaseState(5, -3), 4)
		if err == nil {
			t.Errorf("error expected")
		}
	})

	t.Run("z should be zero or greater", func(t *testing.T) {
		_, err := Solve(newBaseState(5, 3), -4)
		if err == nil {
			t.Errorf("error expected")
		}
	})

	t.Run("z should be lower than either x or y", func(t *testing.T) {
		_, err := Solve(newBaseState(5, 3), 10)
		if err == nil {
			t.Errorf("error expected")
		}
	})
}

func TestSolutions(t *testing.T) {

	t.Run("simple solution", func(t *testing.T) {
		s, err := Solve(newBaseState(5, 3), 4)
		if err != nil {
			t.Fatalf("error %s not expected", err.Error())
		}
		if len(s.Steps) != 6 {
			t.Error("6 steps expected")
		}
		if s.Steps[len(s.Steps)-1].State.X.Amount != 4 &&
			s.Steps[len(s.Steps)-1].State.Y.Amount != 4 {
			t.Errorf("solution not found")
		}
	})

	t.Run("z = 0 should be measurable in 0 steps", func(t *testing.T) {
		s, err := Solve(newBaseState(5, 3), 0)
		if err != nil {
			t.Fatalf("error %s not expected", err.Error())
		}
		if len(s.Steps) != 0 {
			t.Error("0 steps expected")
		}
	})
}

func newBaseState(x, y int) State {
	return State{
		X: Jug{
			Capacity: x,
			Amount:   0,
		},
		Y: Jug{
			Capacity: y,
			Amount:   0,
		},
	}
}
