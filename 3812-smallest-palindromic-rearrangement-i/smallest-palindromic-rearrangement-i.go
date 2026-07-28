func smallestPalindrome(s string) string {
    mid:=len(s)/2
    if len(s)%2==0{
        left:=s[:mid]
        chars:=[]byte(left)
        slices.Sort(chars)
		return string(chars) + rev(string(chars))
    }else{
        left:=s[:mid]
        chars:=[]byte(left)
        slices.Sort(chars)
		return string(chars) +string(s[mid])+ rev(string(chars))
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
