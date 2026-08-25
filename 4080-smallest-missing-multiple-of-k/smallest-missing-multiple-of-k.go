func missingMultiple(nums []int, k int) int {
    seen:=map[int]bool{}
    for _,value := range nums {
        seen[value]=true
    }
    i:=1
    for i>0 {
        if seen[(k*i)]==false{
            return k*i
        }
        i++
    }
    return 0
}