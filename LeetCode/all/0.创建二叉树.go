package main
import(
	"fmt"
)
type BTnode struct {
	Val   int
	Left  *BTnode
	Right *BTnode
}

func createTree(a []int) *BTnode {
	if len(a) == 0 || len(a)%2 == 0 {
		return nil
	}
	root := &BTnode{
		a[0],
		nil,
		nil,
	}
	stack := []*BTnode{root}
	i := 1
	for len(stack) > 0 && i < len(a){
		node := stack[0]
		stack = stack[1:]
		node.Left = &BTnode{a[i], nil, nil}
		node.Right = &BTnode{a[i+1], nil, nil}
		stack = append(stack, node.Left)
		stack = append(stack, node.Right)
		i = i + 2
	}
	return root
}

func printTree(root *BTnode){
	if root==nil{
		return
	}
	fmt.Println(root.Val)
	printTree(root.Left)
	printTree(root.Right)
}


func main(){
	root := createTree([]int{1,2,3,4,5})
	printTree(root)
}
