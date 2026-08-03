func distributeCandies(candies int, num_people int) []int {
    arr:=make([]int,num_people)
    i:=0
    n:=1
    for candies>0 {
        if i>len(arr)-1{
            i=0
        }
        if n>=candies{
            arr[i]+=candies
            return arr
        }
        arr[i]+=n
        candies-=n
        i++
        n++
    }
    return arr
}