package _6_Permutations

//46. Permutations

var ResRow [][]int

func checkIndex(indexs []int, curret int) bool {
	for i := 0; i < curret; i++ {
		if indexs[i] == indexs[curret] {
			return true
		}
	}
	return false
}

func run(indexs []int, nums []int, current int, len int) {
	for i := 0; i < len; i++ {
		indexs[current] = nums[i]
		if checkIndex(indexs, current) {
			continue
		}

		if current == len-1 {
			res_col := make([]int, len)
			for i := 0; i < len; i++ {
				res_col[i] = indexs[i]
			}
			ResRow = append(ResRow,res_col)
			indexs[current] = -1
			return
		}
		newCurrent := current + 1
		run(indexs, nums, newCurrent, len)
		indexs[current] = -1

	}
}

func permute(nums []int) [][]int {
	indexs := make([]int, len(nums))

	//init num_indexs
	for i := 0; i < len(nums); i++ {
		indexs[i] = -1
	}

	//run
	run(indexs, nums, 0, len(nums))
	var tmpRow = ResRow
	ResRow = [][]int{}
	return tmpRow
}
