// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 6.
//!+

// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	s, sep := "", ""
	start := time.Now()
	//for _, arg := range os.Args[1:] {
	for i, arg := range os.Args[:] {
		//s += arg + sep
		s += sep + arg
		sep = " "
		fmt.Println(i, " ", arg)
	}
	elapsed := time.Since(start)
	fmt.Println(s)
	fmt.Println(elapsed)
}

//!-
