/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func mergeKLists(lists []*ListNode) *ListNode {
    if lists == nil {
        return nil
    }
    
    if len(lists) < 2 {
        return lists[0]
    }
    
    rln := lists[0] //result ptr to ListNode
    for i := range lists {
        if i>0{
            rln = mergeTwo(rln, lists[i])
        }
    }
    return rln
}

func mergeTwo(l1,l2 *ListNode) *ListNode {
    var rl,ol *ListNode
    if rl,ol=l1,l2; rl.Val > l2.Val { //rl return list, ol the other list 
        rl=l2
        ol=l1
    } //pick the smaller one as return list
    fmt.Printf("rl = %d picked, ol is %d\n",rl.Val, ol.Val)
    
    prEle := rl //previous rl element
    for re := rl;re!=nil;re=re.Next {
        fmt.Printf("prEle = %d, re = %d\n",prEle.Val, re.Val)
        if re.Val > ol.Val {
            prEle.Next = ol
            poEle := ol
            fmt.Printf(" chaining ol = %d\n",ol.Val)
            for oe := ol; oe != nil; oe=oe.Next {
                fmt.Printf("poEle = %d, ro = %d\n",poEle.Val,oe.Val)
                if oe.Val > re.Val {
                    poEle.Next = re
                    ol = oe
                    fmt.Printf("chain back poEle = %d, ro = %d\n",poEle.Val,oe.Val)
                    break
                }
                poEle = oe
            }
        }
        prEle = re
    }
    return rl
    
}