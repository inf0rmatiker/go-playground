package main

import (
	"fmt"
	"github.com/inf0rmatiker/go-playground/internal/datastructures/maxheap"
	"github.com/inf0rmatiker/go-playground/internal/examples"
)

func main() {

	examples.ArrayExamples()
	examples.SliceExamples()
	examples.StringExample()
	examples.MapsExample()

	heap := maxheap.New()

	heap.Insert(3)
	heap.Insert(5)
	heap.Insert(2)
	heap.Insert(4)
	heap.Print()

	ok, element := heap.Extract()
	if ok {
		fmt.Println("Extracted", element)
	}
	heap.Print()

	ok, element = heap.Extract()
	if ok {
		fmt.Println("Extracted", element)
	}
	heap.Print()

	ok, element = heap.Extract()
	if ok {
		fmt.Println("Extracted", element)
	}
	heap.Print()
}
