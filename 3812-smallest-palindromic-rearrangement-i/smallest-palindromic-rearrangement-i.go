func smallestPalindrome(s string) string {
    mid:=len(s)/2
    if len(s)%2==0{
        left:=s[:mid]
        return lexico(left)+rev(lexico(left))
    }else{
        left:=s[:mid]
        return lexico(left)+string(s[mid])+rev(lexico(left))
    }
    
}
func rev(s string) string{
    chars:=[]byte(s)
    left:=0
    right:=len(chars)-1
    for left<right{
        chars[left],chars[right]=chars[right],chars[left]
        left++
        right--
    }
    return string(chars)
}
func lexico(s string) string{
    chars:=[]byte(s)
    sort.Slice(chars,func(i,j int) bool{
        return chars[i]<chars[j]
    })
    return string(chars)
}
