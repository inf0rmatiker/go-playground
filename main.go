package main

import (
	"fmt"
	"github.com/inf0rmatiker/go-playground/internal/algorithms/dp"
	"github.com/inf0rmatiker/go-playground/internal/datastructures/stack"
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

	//

	/* Algorithms */

	n := 10
	distinctWays := dp.StaircaseRecursiveSolution(n)
	fmt.Printf("Normal Recursive:\tThere are %d distinct ways to climb a staircase of %d steps\n", distinctWays, n)

	distinctWays = dp.StaircaseDPIterativeBottomUp(n)
	fmt.Printf("Iterative DP:\t\tThere are %d distinct ways to climb a staircase of %d steps\n", distinctWays, n)

	distinctWays = dp.StaircaseDPIterativeBottomUpMinimalMemory(n)
	fmt.Printf("Iterative DP (minimal memory):\t\tThere are %d distinct ways to climb a staircase of %d steps\n", distinctWays, n)

	distinctWays = dp.StaircaseDPRecursiveBottomUp(n)
	fmt.Printf("Recursive DP:\t\tThere are %d distinct ways to climb a staircase of %d steps\n", distinctWays, n)

	//cost := []int{10, 15, 20}
	cost := []int{100, 1, 1, 1, 1, 100, 1, 1, 100, 1}
	minCost := dp.MinCostStaircaseRecursive(cost)
	fmt.Printf("Normal Recursive:\t\tMin cost would be %d\n", minCost)

	minCost = dp.MinCostStaircaseMemoization(cost)
	fmt.Printf("Memoized DP:\t\tMin cost would be %d\n", minCost)

	nums := []int{1, 3, 2, 2, 5}
	maxRobbed := dp.RobHousesRecursive(nums)
	fmt.Printf("Normal Recursive:\t\tMax robbed would be %d\n", maxRobbed)

	nums = []int{1, 3, 2, 2, 5}
	maxRobbed = dp.RobHousesRecursive(nums)
	fmt.Printf("Memoization:\t\tMax robbed would be %d\n", maxRobbed)

	pString := "abc"
	fmt.Println(dp.PalindromicSubstrings(pString))

	fmt.Println(dp.DecodeWaysMemoized("110"))

	fmt.Println(dp.UniquePaths(5, 7))

	s := stack.Stack[int]{}
	s.Put(3)
	s.Print()

}
