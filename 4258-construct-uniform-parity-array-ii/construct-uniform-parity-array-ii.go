func uniformArray(nums1 []int) bool {
    if allOdd(nums1) || allEven(nums1){
        return true
    }
    smallestOdd:=99999
    for _,value:=range nums1{
        if value%2!=0 && value<smallestOdd{
            smallestOdd=value
        }
    }
    for _,value:=range nums1{
        if value%2==0 && value<smallestOdd {
            return false
        }
    }
    return true
        
}
func allOdd(nums []int)bool{
    for _,value:= range nums{
        if value%2==0{
            return false
        }
    }
    return true
}
func allEven(nums []int)bool{
    for _,value:=range nums{
        if value%2!=0{
            return false
        }
    }
    return true
}