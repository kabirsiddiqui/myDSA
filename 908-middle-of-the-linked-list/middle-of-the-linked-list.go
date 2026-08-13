/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    temp:=head
    c:=1
    for temp.Next!=nil{
        temp=temp.Next
        c++
    }
    temp=head
    if c%2!=0{
        c=c/2
        for c>0{
            temp=temp.Next
            c--
        }
        head=temp
        return head
    }else{
        c=c/2
        for c>0{
            temp=temp.Next
            c--
        }
        head=temp
        return head
    }
}