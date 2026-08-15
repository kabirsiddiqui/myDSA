func longestSubsequence(nums []int) int {
    result:=0
    allZero:=true
    for _,value := range nums {
        result=result^value
        if value!=0{
            allZero=false
        }
    }
    if allZero==true{
        return 0
    }
    if result!=0{
        return len(nums)
    }
    for i:=0;i<len(nums);i++ {
        arr:=append(nums[:i],nums[i+1:]...)
        if helper(arr)!=0{
            return len(arr)
        }
    }
    return 0
}
func helper(nums []int) int{
    result:=0
    for _,value := range nums{
        result=result^value
    }
    return result
}
