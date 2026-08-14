func maximumLengthSubstring(s string) int {
    max:=0
    for i:=0;i<len(s);i++ {
        for j:=i+1;j<=len(s);j++{
            if !helper(s[i:j]) && len(s[i:j])>=max{
                max=len(s[i:j])
            }
        }
    }
    return max
}
func helper(s string)bool{
    seen:=map[string]int{}
    for _,value := range s{
        seen[string(value)]+=1
        if seen[string(value)]>2{
            return true
        }
    }
    return false
}