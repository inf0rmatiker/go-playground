package examples

import (
	"fmt"
	"strconv"
)

func StringExample() {

	fmt.Println("=== String Slicing Example ===")
	var s string = "Hello there"
	var b byte = s[3]
	var s2 string = string(s[3])
	var substring string = s[:3]
	fmt.Println(s)
	fmt.Println(b)
	fmt.Println(substring)
	fmt.Println(s2)

	fmt.Println("=== String to Number Conversions ===")

	stringToIntExample()
	intToStringExample()

	fmt.Println("=== String to Slice Example ===")

	runes := []rune(s)
	bytes := []byte(s)
	fmt.Printf("%T runes: %v\n", runes, runes)
	fmt.Printf("%T bytes: %v\n", bytes, bytes)
}

func stringToIntExample() {
	i := 10
	s := strconv.Itoa(i)
	fmt.Println(s)
}

func intToStringExample() {
	s := "65"
	i, err := strconv.Atoi(s)
	fmt.Println(i, err)
}
