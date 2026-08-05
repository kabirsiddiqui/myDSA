func remainingMethods(n int, k int, invocations [][]int) []int {
    phoneBook := map[int][]int{}
    for i:=0;i<len(invocations);i++{
        phoneBook[invocations[i][0]]=append(phoneBook[invocations[i][0]],invocations[i][1])
    }
    sus := map[int]bool{k:true}
    queue := []int{k}
    for len(queue) > 0 {
    current := queue[0]
    queue = queue[1:]

    for _, neighbor := range phoneBook[current] {
            if !sus[neighbor] {
                sus[neighbor] = true
                queue = append(queue, neighbor)
            }
        }
    }
    allMethods:=[]int{}
    remaining := []int{}
    for i:=0;i<n;i++ {
        allMethods=append(allMethods,i)
    }
    sort.Ints(allMethods)
    if violation(sus,invocations){
        return allMethods
    }
    for i:=0;i<n;i++ {
        if !sus[i]{
            remaining=append(remaining,i)
        }
    }
    sort.Ints(remaining)
    return remaining
    
}
func violation(sus map[int]bool,invocations[][]int) bool{
    for _,value := range invocations {
        if !sus[value[0]] && sus[value[1]]{
            return true
        }
    }
    return false
}
