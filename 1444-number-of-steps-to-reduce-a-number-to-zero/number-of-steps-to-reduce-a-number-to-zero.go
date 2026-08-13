func numberOfSteps(num int) int {
    return helper(num,0)
}
func helper(num int,c int)int{
    if num==0{
        return c
    }else if num%2==0{
        return helper(num/2,c+1)
    }else{
        return helper(num-1,c+1)
    }
}