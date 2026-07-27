func maxProduct(nums []int) int {
    var index_i int
    var index_j int
    max:=0
    for i:=0;i<len(nums);i++ {
        for j:=0;j<len(nums);j++{
            if i!=j{
                if nums[i]*nums[j]>max{
                    max=nums[i]*nums[j]
                    index_i=i
                    index_j=j
                }
            }
        }
    }
    return (nums[index_i]-1)*(nums[index_j]-1)
}