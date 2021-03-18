package main
import(
	"fmt"
)
type LinkNode struct{
	Val int
	Next *LinkNode
}

func createNode(a []int) *LinkNode{
	head :=&LinkNode{
		0,
		nil,
	}
	prev := head
	for i:=0;i<len(a);i++{
		node := &LinkNode{
			a[i],
			nil,
		}
		prev.Next = node
		prev = node
	}
	return head.Next
}

func printLinkNode(node *LinkNode){
	for node!=nil{
		fmt.Println(node.Val)
		node = node.Next
	}
}
func main(){
	node := createNode([]int{1,2,3,4,5})
	printLinkNode(node)
}
