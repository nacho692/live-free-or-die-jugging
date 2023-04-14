package iterative

import "testing"

func TestNoSolution(t *testing.T) {
	_, err := Solve(9, 3, 4)
	if err == nil {
		t.Error("error expected")
	}
	if err.Error() != "no solution" {
		t.Error("no solution error expected")
	}
}

func TestSimpleSolution(t *testing.T) {
	s, err := Solve(5, 3, 4)
	if err != nil {
		t.Fatalf("error %s not expected", err.Error())
	}
	if len(s.Action) != 6 {
		t.Error("6 steps expected")
	}
}
