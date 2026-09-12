func simplifiedFractions(n int) []string {
    result:=[]string{}
    seen:=map[float32]bool{}
    for den:=2;den<=n;den++{
        for num:=1;num<den;num++{
            frac:=float32(num)/float32(den)
            if seen[frac]!=true{
                seen[frac]=true
                strfrac:=strconv.Itoa(num) + "/" + strconv.Itoa(den)
                result=append(result,strfrac)
            }
        }
    }
    return result
}