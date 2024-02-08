package examples

import "fmt"

func MapsExample() {
	fmt.Println("=== Maps Example ===")

	// Maps are declared as: map[keyType]valueType
	var nilMap map[string]int

	fmt.Printf("nilMap: %v\n", nilMap)

	// Map of string -> []string
	arraysMap := map[string][]string{
		"Fred": {"Tuna", "Salmon", "Bass"},
	}
	fmt.Println("arraysMap:", arraysMap)

}
