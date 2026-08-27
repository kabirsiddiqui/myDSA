func lexGreaterPermutation(s string, target string) string {
    freq:=map[string]int{}
    for _,value:= range s {
        freq[string(value)]++
    }
    answer:=[]byte{}
    for i:=0;i<len(target);i++ {
        if freq[string(target[i])]>=1 {
            answer=append(answer,target[i])
            freq[string(target[i])]--
        }else {
            char:=helper(target[i]+1,freq)
            if char!=0{
                answer=append(answer,char)
                freq[string(char)]--
                return fillRemaining(answer,freq)
            }else{
                return backtracking(answer,target,freq)
            }
            
        }
    }
    if string(answer)==target{
        return backtracking(answer,target,freq)
    }
    return ""
}
func helper(n byte,freq map[string]int) byte {
    for c:=n;c<='z';c++ {
        if freq[string(c)]>0{
            return c
        }
    }
    return 0
}
func fillRemaining(answer []byte, freq map[string]int) string {
    for c := byte('a'); c <= 'z'; c++ {
        for freq[string(c)] > 0 {
            answer = append(answer, c)
            freq[string(c)]--
        }
    }

    return string(answer)
}
func backtracking(answer []byte,target string,freq map[string]int) string {
    for i:=len(answer)-1;i>=0;i-- {
            freq[string(answer[i])]++
            answer=answer[:i]
            char:=helper(target[i]+1,freq)
            if char!=0{
                answer=append(answer,char)
                freq[string(char)]--
                for c:='a';c<='z';c++ {
                    for freq[string(c)]>0 {
                        answer=append(answer,byte(c))
                        freq[string(c)]--
                    }
                }
                return string(answer)
            }
    }
    return string(answer)
}