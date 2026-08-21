func maxAdjacentDistance(nums []int) int {
    max:=nums[0]-nums[len(nums)-1]
    if max<0{
        max=-max
    }
    for i:=1;i<len(nums);i++{
        diff:=nums[i]-nums[i-1]
        if diff<0{
            diff=-diff
        }
        if diff>max{
            max=diff
        }
    }
    return max
}