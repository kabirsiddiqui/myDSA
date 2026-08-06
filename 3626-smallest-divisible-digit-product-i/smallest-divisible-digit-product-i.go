func smallestNumber(n int, t int) int {
    if digitProd(n)%t==0{
        return n
    }else{
        return smallestNumber(n+1,t)
    }
}
func digitProd(temp int)int{
    prod := 1
    for temp >0{
        digit := temp%10
        prod *=digit
        temp = temp/10
    }
    return prod
}