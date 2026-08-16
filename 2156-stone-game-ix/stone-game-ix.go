func stoneGameIX(stones []int) bool {
    c:=[3]int{}
    for _,value:= range stones{
        c[value%3]++
    }
    if c[0]%2==0{
        return c[1]>0 && c[2]>0
    }
    diff:=c[1]-c[2]
    if diff<0{
        diff=-diff
    }
    return diff>2
}