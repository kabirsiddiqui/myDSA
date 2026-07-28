func plusOne(digits []int) []int {
    i:=len(digits)-1
    for i>=0{
        digits[i]=digits[i]+1
        if digits[i]>9{
            digits[i]=digits[i]-10
            i--
        }else{
            return digits
        }
    }
    newDig:=make([]int,len(digits)+1)
    newDig[0]=1
    return newDig
}
