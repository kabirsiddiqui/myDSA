func checkDivisibility(n int) bool {
    sum:=0
    prod:=1
    temp:=n
    for temp>0{
        digit:=temp%10
        prod*=digit
        sum+=digit
        temp/=10
    }
    if n%(prod+sum)==0{
        return true
    }else{
        return false
    }
}