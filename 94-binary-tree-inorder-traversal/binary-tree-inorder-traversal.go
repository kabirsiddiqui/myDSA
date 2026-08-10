/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    result:=[]int{}
    return helper(root,result)
}
func helper(root *TreeNode,result []int)[]int{
    if root==nil{
        return result
    }
    result=helper(root.Left,result)
    result=append(result,root.Val)
    result=helper(root.Right,result) 
    return result
}