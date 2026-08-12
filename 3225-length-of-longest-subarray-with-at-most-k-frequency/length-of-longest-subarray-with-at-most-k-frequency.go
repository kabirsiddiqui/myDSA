func maxSubarrayLength(nums []int, k int) int {
    freq:=map[int]int{}
    left:=0
    right:=0
    maxLength:=0
    for right<len(nums){
        freq[nums[right]]+=1
        for freq[nums[right]]>k{
            freq[nums[left]]=freq[nums[left]]-1
            left++
        }
        currentLength:=right-left+1
        if currentLength>maxLength{
            maxLength=currentLength
        }
        right++
    }
    return maxLength
}