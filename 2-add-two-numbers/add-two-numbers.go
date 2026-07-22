/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


type LinkedList struct{
    head *ListNode
}



func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    list := LinkedList{}
    carry :=0
    for l1!=nil || l2!=nil || carry!=0{
        var digit1 int
        if l1 != nil {
            digit1 = l1.Val
        }

        var digit2 int
        if l2 != nil {
            digit2 = l2.Val
        }
        sum := digit1+digit2+carry
        digit := sum%10
        carry = sum/10
        newNode := &ListNode{Val:digit,Next:nil} 
        list.appendBack(newNode)
        if l1!=nil{
            l1=l1.Next
        }   
        if l2!=nil{
            l2=l2.Next
        }
    }
    return list.head
}

func (l *LinkedList) appendBack(newNode *ListNode){
    temp := l.head
    if temp==nil{
        l.head=newNode
        return
    }
    for temp != nil{
        if temp.Next==nil{
            temp.Next=newNode
            return
        }
        temp=temp.Next
    }
}