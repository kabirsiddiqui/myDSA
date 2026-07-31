func digitFrequencyScore(n int) int {
    // var num []int
    // for n>0 {
    //     num=append(num,n%10)
    //     n=n/10
    // }
    score:=0
    // for _,value := range num{
    //     score+=value
    // }
    var digit int
    for n>0{
        digit=n%10
        score+=digit
        n=n/10
    }
    return score
}