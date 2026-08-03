func stoneGameIII(stoneValue []int) string {
    i:=0
    cache := make([]int,len(stoneValue))
    for index,_ := range cache {
        cache[index]=math.MinInt
    }
    if score(i,stoneValue,cache)>0{
        return "Alice"
    }else if score(i,stoneValue,cache)<0{
        return "Bob"
    }else{
        return "Tie"
    }
}

func score(i int,arr []int,cache []int) int {
    if i>=len(arr) {
        return 0
    }
    if cache[i]!=math.MinInt {
        return cache[i]
    }
    total :=0
    best := math.MinInt
    for k:=0;k<3 && i+k<len(arr);k++{
        total += arr[i+k]
        playerScore := total-score(i+k+1,arr,cache)
        if playerScore>best{
            best=playerScore
        }
    }
    cache[i]=best
    return cache[i]
}