func winnerSquareGame(n int) bool {
   wp:=make([]bool,n+1)
   wp[0]=false
   for i:=1;i<=n;i++ {
    for x:=1;x*x<=i;x++{
        if wp[i-(x*x)]==false{
            wp[i]=true
        }
    }
    if wp[i]!=true{
        wp[i]=false
    }
   }
   return wp[n]
}
