func findMissingElements(nums []int) []int {
    smallest:=math.MaxInt
    largest:=math.MinInt
    seen := map[int]bool{}
    for _,value := range nums{
        if value<smallest{
            smallest=value
        }
        if value>largest{
            largest=value
        }
        seen[value]=true
    }
    missing:=[]int{}
    for i:=smallest+1;i<largest;i++{
        if !seen[i]{
            missing=append(missing,i)
        }
    }
    return missing
}