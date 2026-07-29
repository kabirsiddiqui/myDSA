func isValid(s string) bool {
    pairs := map[byte]byte{
        ')':'(',
        '}':'{',
        ']':'[',
    }
    chars:=[]byte(s)
    stack:=[]byte{}
    for _,char := range chars{
        if char=='('||char=='{'||char=='['{
            stack=append(stack,char)
        }else{
            if len(stack)==0{
                return false
            }
            top:=len(stack)-1
            if pairs[char]!=stack[top]{
                return false
            }else{
                stack=stack[:len(stack)-1]
            }
        }
    }
    if len(stack)==0{
        return true
    }
    return false
}