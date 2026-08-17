func removeDuplicates(nums []int) int {
    i:=0

    for j:=1;j<len(nums);j++{
        if nums[i]==nums[j]{
            continue
        }else{
            i++
            nums[i]=nums[j]
        }
    }
    return i+1
}