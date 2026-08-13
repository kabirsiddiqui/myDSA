func canConstruct(ransomNote string, magazine string) bool {
    freq:=map[string]int{}
    for _,value := range magazine {
        freq[string(value)]+=1
    }
    for _,value := range ransomNote {
        if freq[string(value)]>0{
            freq[string(value)]--
            continue
        }else{
            return false
        }
    }
    return true
}