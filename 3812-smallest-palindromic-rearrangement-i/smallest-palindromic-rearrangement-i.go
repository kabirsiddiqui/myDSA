func smallestPalindrome(s string) string {
    mid:=len(s)/2
    if len(s)%2==0{
        left:=s[:mid]
        chars:=[]byte(left)
        slices.Sort(chars)
        og:=string(chars)
        slices.Reverse(chars)
		return og + string(chars)
    }else{
        left:=s[:mid]
        chars:=[]byte(left)
        slices.Sort(chars)
        og:=string(chars)
        slices.Reverse(chars)
		return og +string(s[mid])+ string(chars)
    }
    
}

