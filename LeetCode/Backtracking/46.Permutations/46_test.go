package _6_Permutations

import (
	"reflect"
	"testing"
)

//46. Permutations

//Test
func TestPermute(t *testing.T) {
	answer1 := make([][]int,6)
	answer1[0] = []int{1,2,3}
	answer1[1] = []int{1,3,2}
	answer1[2] = []int{2,1,3}
	answer1[3] = []int{2,3,1}
	answer1[4] = []int{3,1,2}
	answer1[5] = []int{3,2,1}

	nums := []int{1, 2, 3}
	res := permute(nums)
	if !reflect.DeepEqual(res,answer1){
		t.Fatalf("error ")
	}

}
