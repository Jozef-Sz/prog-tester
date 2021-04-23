package testcase

import "testing"

func strlistcmp(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, elm := range a {
		if elm != b[i] {
			return false
		}
	}
	return true
}

func TestRegularArgsSplit(t *testing.T) {
	args := "run -r hello --lower"
	expect := []string{"run", "-r", "hello", "--lower"}
	result := splitArguments(args)
	if !strlistcmp(result, expect) {
		t.Errorf("got: %v\n expected: %v\n", result, expect)
	}
}

func TestNestedArgsSplit(t *testing.T) {
	args := "run -r \"hello, world\" --lower"
	expect := []string{"run", "-r", "hello, world", "--lower"}
	result := splitArguments(args)
	if !strlistcmp(result, expect) {
		t.Errorf("got: %v\n expected: %v\n", result, expect)
	}
}
