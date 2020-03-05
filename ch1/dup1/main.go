// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.
//!+

// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("main started\n")
	counts := make(map[string]int)
	fmt.Printf("counts initialization completed also input bufio will start\n")
	input := bufio.NewScanner(os.Stdin)
	fmt.Printf("input bufio is stared\n")
	for input.Scan() {
		fmt.Printf("input is scanned\n")
		if input.Text() == "" {
			break
		}
		counts[input.Text()]++
		fmt.Println("it is counted with count :", counts)

	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		fmt.Printf("every count is printed\n")
		if n > 1 {
			fmt.Println("it is counted with count :", counts)
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

//!-
