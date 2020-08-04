package _0_Valid_Parentheses

import (
	"testing"
)

//Test
func TestIsValid(t *testing.T) {
	tables := []struct {
		x string
		y bool
	}{
		{"([)]", false},
		{"{}[]()", true},
	}

	for _, table := range tables {
		y := isValid(table.x)
		if y != table.y {
			t.Fatalf("error " + table.x)
		}
	}
}
