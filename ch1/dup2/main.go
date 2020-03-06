// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			var repeat bool = false
			for line, n := range counts {
				if n > 1 {
					//fmt.Printf("count map is :%v", counts)
					fmt.Printf("%d\t%s\n", n, line)
					if repeat == false {
						repeat = true
						fmt.Println("file is :", arg)
					}
				}
			}
			f.Close()
		}
	}
	/*
		for line, n := range counts {
			if n > 1 {
				fmt.Printf("count map is :%v", counts)
				fmt.Printf("%d\t%s\n", n, line)
			}
		}*/

}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if input.Text() == "" {
			break
		}
		counts[input.Text()]++
		//fmt.Printf("count map is :%v", counts)
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-
