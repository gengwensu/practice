/**
 * Merge k sorted linked lists and return it as one sorted list. 
 * Analyze and describe its complexity.
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func mergeKLists(lists []*ListNode) *ListNode {
    var head,ptr *ListNode
    for i:=len(lists)-1;i>=0;i-- {
        if lists[i]==nil {
            if i<len(lists)-1{
                copy(lists[i:],lists[i+1:])
            }
            lists=lists[:len(lists)-1]
        }
    }
    if len(lists)==0 {
        return head
    }
    
    for len(lists)>0 {
        sort.Sort(ByVal(lists))
        if head==nil{
            head=lists[0]
            ptr=lists[0]
        } else{
            ptr.Next=lists[0]
            ptr=ptr.Next
        }
        lists[0]=lists[0].Next
        ptr.Next=nil
        if lists[0]==nil{
            copy(lists,lists[1:])
            lists=lists[:len(lists)-1]
        }
    }
    return head
}
type ByVal []*ListNode
func (h ByVal) Len() int {return len(h)}
func (h ByVal) Less(i,j int) bool {return h[i].Val<h[j].Val}
func (h ByVal) Swap(i,j int) {h[i],h[j]=h[j],h[i]}