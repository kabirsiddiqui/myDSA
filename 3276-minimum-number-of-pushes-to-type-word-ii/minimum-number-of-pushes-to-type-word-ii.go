func minimumPushes(word string) int {
    
    freq := make(map[byte]int)
	for i := 0; i < len(word); i++ {
		freq[word[i]]++
	}
    seen := make(map[byte]bool)
    chars := []byte(word)
    uniqueChars := []byte{}
    for _,value := range chars {
        if !seen[value]{
            uniqueChars = append(uniqueChars,value)
            seen[value]=true
        }
    }
    sort.Slice(uniqueChars, func(i,j int) bool {
        if freq[uniqueChars[i]] == freq[uniqueChars[j]] {
            return uniqueChars[i] < uniqueChars[j]
        }
        return freq[uniqueChars[i]]>freq[uniqueChars[j]]
    })
    minCost := 0
    cost := map[byte]int{}
    for _,char := range uniqueChars {
        if len(cost)<8{
            cost[char]=1
        }else if len(cost)>=8 && len(cost)<16 {
            cost[char]=2
        }else if len(cost)>=16 && len(cost)<24 {
            cost[char]=3
        }else{
            cost[char]=4
        }
    }
    for _,char := range uniqueChars {
        minCost+=cost[char]*freq[char]
    }
    return minCost
}