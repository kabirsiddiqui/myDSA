func maxProduct(n int) int {
    var digits []int
    temp:=n
    for temp>0{
        digit:=temp%10
        digits = append(digits,digit)
        temp = temp/10
    }
    max:=0
    for i:=0;i<len(digits);i++{
        for j:=i+1;j<len(digits);j++{
            if digits[i]*digits[j]>max{
                max=digits[i]*digits[j]
            }
        }
    }
    return max

}