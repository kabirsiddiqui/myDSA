func findMissingElements(nums []int) []int {
    smallest:=math.MaxInt
    largest:=math.MinInt
    for _,value := range nums{
        if value<smallest{
            smallest=value
        }
        if value>largest{
            largest=value
        }
    }
    missing:=[]int{}
    for i:=smallest+1;i<largest;i++{
        if !present(nums,i){
            missing=append(missing,i)
        }
    }
    return missing
}
func present(nums []int,target int) bool{
    for _,value:=range nums{
        if value==target{
            return true
        }
    }
    return false
}