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
