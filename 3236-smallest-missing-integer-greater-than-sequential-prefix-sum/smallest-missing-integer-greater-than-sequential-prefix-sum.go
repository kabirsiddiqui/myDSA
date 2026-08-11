func missingInteger(nums []int) int {
    seen := map[int]bool{nums[0]:true}
    for _,value := range nums {
        seen[value]=true
    }
    newNums:=[]int{}
    for i:=1;i<len(nums);i++ {
        if nums[i]!=nums[i-1]+1 {
            newNums=nums[0:i]
            break
        }
    }
    if len(newNums)==0{
        newNums=nums[:len(nums)]
    }
    sum:=0
    for _,value := range newNums {
        sum+=value
    }
    for seen[sum]{
        sum++
    }
    return sum
}