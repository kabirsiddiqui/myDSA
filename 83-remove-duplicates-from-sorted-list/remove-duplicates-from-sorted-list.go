/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    temp:=head
    for temp!=nil && temp.Next!=nil{
        current:=temp.Next
        for current!=nil && temp.Val==current.Val {
            current=current.Next
        }
        temp.Next=current
        temp=temp.Next
    }
    return head
}