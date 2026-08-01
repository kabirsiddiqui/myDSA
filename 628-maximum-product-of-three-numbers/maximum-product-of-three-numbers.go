func maximumProduct(nums []int) int {
    sort.Ints(nums)
    i:=len(nums)-1
    return max(nums[0]*nums[1]*nums[i],
    nums[i]*nums[i-1]*nums[i-2])
}
