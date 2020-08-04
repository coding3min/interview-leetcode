package _7

func findContinuousSequence(target int) [][]int {
	var res [][]int
	for n := 2; n < target; n++ {
		a1 := target/n - n/2 - (n%2 - 1)
		if a1 < 1 {
			break
		}
		if target == (a1+(a1+n-1))*n/2  {
			var res_col []int
			for i := a1; i < a1+n; i++ {
				res_col = append(res_col, i)
			}
			res = append(res, res_col)
		}
	}
	var revort_res [][]int
	for i:=len(res)-1;i>=0;i--{
		revort_res = append(revort_res,res[i])
	}
	return revort_res
}
