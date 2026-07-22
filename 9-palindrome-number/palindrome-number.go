func isPalindrome(x int) bool {
    if x<0{
       return false
    }
    rev := reverse(x)
    if(rev == x){
        return true
    }
    return false

}
func reverse(x int) int{
    var rev int
    for x>0{
        dig := x%10
        rev = (rev*10)+dig
        x=x/10
    }
    return rev
}