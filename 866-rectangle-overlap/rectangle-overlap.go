func isRectangleOverlap(rec1 []int, rec2 []int) bool {
    left:=max(rec1[0],rec2[0])
    right:=min(rec1[2],rec2[2])
    width:=right-left
    if width<1{
        return false
    }
    bottom:=max(rec1[1],rec2[1])
    top:=min(rec1[3],rec2[3])
    height:=top-bottom
    if height<1{
        return false
    }
    return true
}