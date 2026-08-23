func sumGame(num string) bool {
    mid:=len(num)/2
    left:=num[:mid]
    right:=num[mid:]
    c1:=0
    s1:=0
    c2:=0
    s2:=0
    for _,value := range left{
        n,err:=strconv.Atoi(string(value))
        if err==nil{
            s1+=n
        }else{
            c1++
        }
    }
    for _,value := range right{
        n,err:=strconv.Atoi(string(value))
        if err==nil{
            s2+=n
        }else{
            c2++
        }
    }
    return 2*s1 + 9*c1 != 2*s2 + 9*c2

}