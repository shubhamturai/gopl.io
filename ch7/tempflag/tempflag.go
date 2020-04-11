// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 181.

// Tempflag prints the value of its -temp (temperature) flag.
package main

import (
	"flag"
	"fmt"

	"gopl.io/ch7/tempconv"
)

//!+
var c = tempconv.CelsiusFlag("c", 20.0, "the temperature")
var f = tempconv.FahrenheitFlag("f", 20.0, "the temperature")
var k = tempconv.KelvinFlag("k", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*c)
	fmt.Println(*f)
	fmt.Println(*k)
}

//!-
