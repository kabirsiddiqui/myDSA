func maximumLengthSubstring(s string) int {
    freq:=map[byte]int{}
    left:=0
    right:=0
    maxLength:=0
    for right<len(s){
        freq[s[right]]++
        for freq[s[right]]>2{
            freq[s[left]]--
            left++
        }
        currentLength:=right-left+1
        if currentLength>maxLength{
            maxLength=currentLength
        }
        right++
    }
    return maxLength
}