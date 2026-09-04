func firstStableIndex(nums []int, k int) int {
    for i:=0;i<len(nums);i++{
        if maxNum(nums[:i+1])-minNum(nums[i:])<=k{
            return i
        }
    }
    return -1
}
func maxNum(nums []int) int{
    max:=math.MinInt
    for _,value:=range nums{
        if value>max{
            max=value
        }
    }
    return max
}
func minNum(nums []int) int{
    min:=math.MaxInt
    for _,value:=range nums{
        if value<min{
            min=value
        }
    }
    return min
}