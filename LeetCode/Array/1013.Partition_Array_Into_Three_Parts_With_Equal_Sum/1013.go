package _013_Partition_Array_Into_Three_Parts_With_Equal_Sum

//暴力
func canThreePartsEqualSum(A []int) bool {
	suma, sumb, sumc := 0, 0, 0
	tmpsumc := 0

	for k := 1; k < len(A); k++ {
		tmpsumc = tmpsumc + A[k]
	}

	for i := 0; i < len(A)-2; i++ {
		suma = suma + A[i]
		sumb = 0
		sumc = tmpsumc
		for j := i + 1; j < len(A)-1; j++ {
			sumb = sumb + A[j]
			sumc = sumc - A[j]
			if suma == sumb && sumb == sumc {
				return true
			}
		}

		tmpsumc = tmpsumc - A[i+1]
	}
	return false
}

//数学
func canThreePartsEqualSum2(A []int) bool {
	tmpSum, sum, singeSum, count := 0, 0, 0, 0

	for i := 0; i < len(A); i++ {
		sum = sum + A[i]
	}
	if sum%3 != 0 {
		return false
	}
	singeSum = sum / 3

	for i := 0; i < len(A); i++ {
		tmpSum = tmpSum + A[i]
		if tmpSum == singeSum {
			count = count + 1
			tmpSum = 0
			if count == 2 && i < len(A)-1 {
				return true
			}
		}
	}
	return false
}
