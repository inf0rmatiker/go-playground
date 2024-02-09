package main

import (
	"github.com/inf0rmatiker/go-playground/internal/datastructures/linkedlist"
	"github.com/inf0rmatiker/go-playground/internal/examples"
)

func main() {

	examples.ArrayExamples()
	examples.SliceExamples()
	examples.StringExample()
	examples.MapsExample()

	//heap := maxheap.New()
	//
	//heap.Insert(3)
	//heap.Insert(5)
	//heap.Insert(2)
	//heap.Insert(4)
	//heap.Print()
	//
	//ok, element := heap.Extract()
	//if ok {
	//	fmt.Println("Extracted", element)
	//}
	//heap.Print()
	//
	//ok, element = heap.Extract()
	//if ok {
	//	fmt.Println("Extracted", element)
	//}
	//heap.Print()
	//
	//ok, element = heap.Extract()
	//if ok {
	//	fmt.Println("Extracted", element)
	//}
	//heap.Print()

	ll := linkedlist.New()
	ll.Append(5)
	ll.Append(6)
	ll.Append(9)
	ll.Append(3)
	ll.Prepend(15)
	ll.Insert(88, 0)
	ll.Insert(99, 0)
	ll.Insert(33, 2)
	ll.Print()
}
