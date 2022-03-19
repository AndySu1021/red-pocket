package main

import "fmt"

type Block struct {
	Val  int
	Next *Block
}

func main() {
	var arr [4]*Block
	for i := 0; i < 4; i++ {
		arr[i] = &Block{
			Val:  i,
			Next: nil,
		}
	}

	moveOnto(&arr, 3, 1)
	moveOnto(&arr, 2, 1)

	for i := 0; i < 4; i++ {
		fmt.Printf("%d:", i)
		curr := arr[i]
		for curr != nil {
			fmt.Printf(" %d", curr.Val)
			curr = curr.Next
		}
		fmt.Println()
	}
}

func moveOnto(arr *[4]*Block, a int, b int) {
	curr := arr[a]
	prev := curr
	for curr != nil {
		if curr.Val != a {
			arr[curr.Val] = curr
			prev.Next = nil
		}
		prev = curr
		curr = curr.Next
	}

	curr = arr[b]
	prev = curr
	for curr != nil {
		if curr.Val != b {
			arr[curr.Val] = curr
			prev.Next = nil
		}
		prev = curr
		curr = curr.Next
	}

	arr[b].Next = arr[a]
	arr[a] = nil
}
