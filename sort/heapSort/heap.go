package heapSort

import (
	"github.com/Golang-Commons/algorithmCommons/structure/heap"
)

func sort(arr []int) []int {
	h := heap.NewMin()
	for i := 0; i < len(arr); i++ {
		h.Insert(heap.Int(arr[i]))
	}

	for i := 0; i < len(arr); i++ {
		arr[i] = int(h.Extract().(heap.Int))
	}

	return arr
}
