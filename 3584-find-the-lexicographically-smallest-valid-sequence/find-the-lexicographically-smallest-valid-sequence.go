func validSequence(word1 string, word2 string) []int {
    indices:=[]int{}
    last:=[]int{}
    for i:=0;i<len(word2);i++{
        last=append(last,-1)
    }
    k:=len(word1)-1
    for i:=len(word2)-1;i>=0;i--{
        for k>=0 {
            if string(word2[i])==string(word1[k]){
                last[i]=k
                k--
                break
            }
            k--
        }
    }
    misMatch:=false
    j:=0
    i:=0
    for i<len(word2) && j<len(word1){
        if string(word2[i])==string(word1[j]){
            indices=append(indices,j)
            j++
            i++
        }else if i+1<len(word2) && j+1<=last[i+1] && misMatch==false{
            indices=append(indices,j)
            misMatch=true
            j++
            i++
        }else if misMatch==false && i==len(word2)-1{
            indices=append(indices,j)
            misMatch=true
            i++
            j++
        }else{
            j++
        }
    }
    if i!=len(word2){
        return []int{}
    }
    return indices
}