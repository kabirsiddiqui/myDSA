func minimumDistance(nums []int) int {
    if len(nums)<3{
        return -1
    }
    pos:=map[int][]int{}
    for index,value := range nums {
        pos[value]=append(pos[value],index)
    }
    min:=math.MaxInt
    for _,value := range pos {
        if len(value)>=3{
            if helper(value)<min{
                min=helper(value)
            }
        }
    }
    if min==math.MaxInt{
        return -1
    }
    return min 
}
func helper(nums []int) int {
    min:=math.MaxInt
    for i:=0;i<=len(nums)-3;i++ {
        j:=i+3
        temp:=nums[i:j]
        dist:=2*(temp[2]-temp[0])
        if dist<min{
            min=dist
        }
    }
    return min
}
