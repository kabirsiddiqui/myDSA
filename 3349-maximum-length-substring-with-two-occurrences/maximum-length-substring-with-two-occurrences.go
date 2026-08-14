func maximumLengthSubstring(s string) int {
    freq:=map[string]int{}
    left:=0
    right:=0
    maxLength:=0
    for right<len(s){
        freq[string(s[right])]+=1
        for freq[string(s[right])]>2{
            freq[string(s[left])]--
            left++
        }
        currentLength:=right-left+1
        if currentLength>=maxLength{
            maxLength=currentLength
        }
        right++
    }
    return maxLength
}