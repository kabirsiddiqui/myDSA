func minimumPushes(word string) int {
    if len(word)<=8{
        return len(word)
    }else if len(word)>8 && len(word)<=16{
        return ((len(word)-8)*2)+8
    }else if len(word)>16 && len(word)<=24{
        return ((len(word)-16)*3)+(8*2)+8
    }else{
        return ((len(word)-24)*4)+(8*3)+(8*2)+8
    }
}