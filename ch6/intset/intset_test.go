// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package intset

import "fmt"

func Example_one() {
	//!+main
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"

	fmt.Println(x.Has(9), x.Has(123)) // "true false"
	//!-main

	// Output:
	// {1 9 144}
	// {9 42}
	// {1 9 42 144}
	// true false
}

func Example_two() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)
	x.Add(37)
	x.Add(128)
	//x.Remove(37)

	//!+note
	fmt.Println(&x)         // "{1 9 42 144}"
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x)          // "{[4398046511618 0 65536]}"
	//!-note
	fmt.Println(x.Len())
	x.Remove(42)
	fmt.Println(&x)
	//fmt.Printf("%p\n", &x)
	
	//var y IntSet
	y := x.Copy()
	fmt.Println(y)

	x.Clear()
	fmt.Println(&x)
	// Output:
	// {1 9 37 42 128 144}
	// {1 9 37 42 128 144}
	// {[4535485465090 0 65537]}
	// 6
	// {1 9 37 128 144}
	// {1 9 37 128 144}
	// {}
}
