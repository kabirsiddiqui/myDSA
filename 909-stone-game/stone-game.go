func stoneGame(piles []int) bool {
    cache := make([][]int, len(piles))
    for i := range cache {
        cache[i] = make([]int, len(piles))
        for j := range cache[i] {
            cache[i][j] = math.MinInt
        }
    }
    if score(0,len(piles)-1,piles,cache)>=0{
        return true
    }else{
        return false
    }
    return true
}
func score(i,j int,piles []int,cache [][]int) int {
    if i==j{
        return piles[i]
    }
    if cache[i][j]==math.MinInt{
        cache[i][j]=max(piles[i]-score(i+1,j,piles,cache),piles[j]-score(i,j-1,piles,cache))
        return max(piles[i]-score(i+1,j,piles,cache),piles[j]-score(i,j-1,piles,cache))
    }else{
        return cache[i][j]
    }
    
}