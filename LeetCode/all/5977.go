func minSwaps(nums []int) int {
    var count int
    for _,v := range nums{
        if v == 1{
            count++
        }
    }
    lenOne := count
    count = 0
    for i:=0;i<lenOne;i++{
        if nums[i]==0{
            count++
        }
    }
    var minCount = count
    for i,j:=0,lenOne;j<len(nums)+lenOne;i,j=i+1,j+1{
        if nums[i%len(nums)] == 0{
            count--
        }
        if nums[j%len(nums)] == 0{
            count++
        }
        if minCount>count{
            minCount = count
        }
    }
    return minCount
}