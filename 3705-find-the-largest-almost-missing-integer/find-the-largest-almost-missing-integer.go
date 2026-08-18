func largestInteger(nums []int, k int) int {
    if k==len(nums){
        max:=-1
        for _,value := range nums{
            if value>max{
                max=value
            }
        }
        return max
    }
    freq:=map[int]int{}
    for i:=0;i<=len(nums)-k;i++ {
        for j:=0;j<k;j++{
            freq[nums[i+j]]++
        }
    }
    max:=-1
    for _,value:= range nums{
        if freq[value]==1 && value>max{
            max=value
        }
    }
    return max
}