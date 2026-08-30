func minimumDeletions(nums []int) int {
    if len(nums)==1{
        return 1
    }
    if len(nums)==2{
        return 2
    }
    min:=9999
    max:=-9999
    maxIndex:=0
    minIndex:=0
    for index,value := range nums {
        if value>max {
            max=value
            maxIndex=index
        }
        if value < min {
			min = value
			minIndex = index
		}
    }
    deleteFront:=0
    deleteBack:=0
    if maxIndex>minIndex{
        deleteFront=maxIndex+1
        deleteBack=len(nums)-minIndex
    }else{
        deleteFront=minIndex+1
        deleteBack=len(nums)-maxIndex
    }
    deleteBoth:=0
    for i:=0;i<len(nums);i++{
        if i==minIndex || i==maxIndex{
            deleteBoth+=i+1
            break
        }
    }
    c:=1
    for j:=len(nums)-1;j>=0;j--{
        if j==minIndex || j==maxIndex{
            deleteBoth+=c
            break
        }
        c++
    }
    minimum:=min3(deleteFront,deleteBack,deleteBoth)
    return minimum
}
func min3(a, b, c int) int {
    if a < b && a < c {
        return a
    }
    if b < c {
        return b
    }
    return c
}