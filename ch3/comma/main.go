// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	//n := len(s)
	var p int
	var buf bytes.Buffer
	n := len(s)
	if strings.ContainsRune(s, '.') {
		p = strings.IndexRune(s, '.')
		for i := 0; i < n; i++ {
			fmt.Fprint(&buf,string(s[i]))
			if (p-i) % 3 == 1 && i != p-1 && i <= p || (i-p) % 3 == 0 && i-p != n-p-1 && i > p {
				buf.WriteString(",")
			}
		}
		
	} else {
		for i := 0; i < n; i++ {
			fmt.Fprint(&buf,string(s[i]))
			if (n-i) % 3 == 1 && i != n-1 {
				buf.WriteString(",")
			}
		}
	}
	return buf.String()
}

/*
//!+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

//!-
*/
