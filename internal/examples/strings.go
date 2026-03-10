package examples

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
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

// Sort a string by its characters

// Implementation of sort.Interface. Need three functions:
// 1. Less(i int, j int) bool
// 2. Len() int
// 3. Swap(i int, j int)
type sortedRunes []rune

func (sr sortedRunes) Less(i, j int) bool {
	return sr[i] < sr[j]
}

func (sr sortedRunes) Len() int {
	return len(sr)
}

func (sr sortedRunes) Swap(i, j int) {
	sr[i], sr[j] = sr[j], sr[i]
}

// Takes a string, s, and returns a new string with all the characters sorted alphabetically.
func SortString(s string) string {
	sr := sortedRunes(s)
	sort.Sort(sr)
	return string(sr)
}

func Encode(strs []string) string {
	var builder strings.Builder
	for _, str := range strs {
		fmt.Fprintf(&builder, "%d;%s", len([]rune(str)), str)
	}
	return builder.String()
}

func Decode(encoded string) []string {
	// If the encoded string was empty, then so should be the returned slice
	if len(encoded) == 0 {
		return []string{}
	}

	// 1. Parse a number, num, until we get to a ';' rune.
	// 2. Read forward num runes, building a current rune slice to be added to the results.
	currNum := []rune{}
	num := 0
	decodedStrings := []string{}
	runeSlice := []rune(encoded)
	for i := 0; i < len(runeSlice); i++ {
		r := runeSlice[i]
		if r == ';' {
			// Done building current number:
			// Convert it to an int, num, then take the next num runes starting at i+1 index.
			num, _ = strconv.Atoi(string(currNum))
			decodedStrings = append(decodedStrings, string(runeSlice[i+1:i+1+num]))
			i += num

			currNum = []rune{}
		} else {
			currNum = append(currNum, r)
		}
	}
	return decodedStrings
}
