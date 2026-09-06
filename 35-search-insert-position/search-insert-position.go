func searchInsert(nums []int, target int) int {
    for index,value := range nums{
        if value==target{
            return index
        }
    }
    rightIndex:=-1
    for index,value:=range nums{
        if value>target{
            rightIndex=index
            break
        }
    }
    if rightIndex!=-1{
        return rightIndex
    }else{
        return len(nums)
    }
    
}