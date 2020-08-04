package _7

import (
	"reflect"
	"testing"
)

//Test
func TestFindContinuousSequence(t *testing.T) {

	answer1 := make([][]int, 2)
	answer1[0] = []int{2, 3, 4}
	answer1[1] = []int{4, 5}

	answer2 := make([][]int, 3)
	answer2[0] = []int{1, 2, 3, 4, 5}
	answer2[1] = []int{4, 5, 6}
	answer2[2] = []int{7, 8}
	tables := []struct {
		x int
		y [][]int
	}{
		{9, answer1},
		{15, answer2},
	}

	for _, table := range tables {
		y := findContinuousSequence(table.x)
		if !reflect.DeepEqual(y, table.y) {
			t.Fatalf("error %d ", table.x)
		}
	}
}
