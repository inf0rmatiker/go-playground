package examples

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ReadLines uses a bufio.Scanner on an os.Stdin reader to read a specified number of lines.
func ReadLines() error {

	scanner := bufio.NewScanner(os.Stdin)
	var err error
	var linesToRead int

	fmt.Print("Enter how many lines to read: ")
	if scanner.Scan() {
		linesToRead, err = strconv.Atoi(strings.TrimSpace(scanner.Text()))
		if err != nil {
			return scanner.Err()
		}
	} else {
		return fmt.Errorf("failed to scan for text")
	}
	fmt.Printf("Reading %d lines:\n", linesToRead)

	lines := make([]string, 0, linesToRead)
	for i := range linesToRead {
		fmt.Printf("Enter line %d: ", i+1)
		if scanner.Scan() {
			lines = append(lines, scanner.Text())
		} else {
			return scanner.Err()
		}
	}

	for i, line := range lines {
		fmt.Printf("%d: %s\n", i+1, line)
	}
	return nil
}
