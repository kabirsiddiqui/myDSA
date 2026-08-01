func predictTheWinner(nums []int) bool {
    i:=0
    j:=len(nums)-1
    if score(i,j,nums)>=0{
        return true
    }else{
        return false
    }
    
}

func score(i,j int,nums []int) int{
    if i==j{
        return nums[i]
    }
    return max(nums[i]-score(i+1,j,nums),nums[j]-score(i,j-1,nums))

}