func runningSum(nums []int) []int {
    sum:=0
    result:=[]int{}
    for _,value := range nums{
        result=append(result,value+sum)
        sum+=value
    }
    return result
}