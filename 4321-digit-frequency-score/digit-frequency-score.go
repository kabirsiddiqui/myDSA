func digitFrequencyScore(n int) int {
    var num []int
    for n>0 {
        num=append(num,n%10)
        n=n/10
    }
    // freq := map[int]int{}
    // for _,value := range num{
    //     freq[value]++
    // }
    score:=0
    for _,value := range num{
        score+=value
    }
    return score
}