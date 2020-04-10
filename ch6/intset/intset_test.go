// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package intset

import "fmt"

func Example_one() {
	//!+main
	var x, y IntSet
	x.AddAll(1,0,98,56,23)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{0 1 9 23 56 98 144}"

	y.AddAll(9, 15, 54, 0)
	y.Add(42)
	fmt.Println(y.String()) // "{0 9 15 42 54}"

	x.IntersectWith(&y)
	fmt.Println(x.String()) // "{0 9}"

	x.AddAll(1,0,98,56,23)
	fmt.Println(x.String()) // "{0 1 9 23 56 98}"

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{0 1 9 15 23 42 54 56 98}"

	x.DifferenceWith(&y)
	fmt.Println(x.String()) // "{1 23 56 98}"
	fmt.Println(*x.Elems()) // "[1 23 56 98]"
	
	fmt.Println(x.Has(23), x.Has(123)) // "true false"
	//!-main

	// Output:
	// {0 1 9 23 56 98 144}
	// {0 9 15 42 54}
	// {0 9}
	// {0 1 9 23 56 98}
	// {0 1 9 15 23 42 54 56 98}
	// {1 23 56 98}
	// [1 23 56 98]
	// true false
}// 

func Example_two() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)
	x.Add(37)
	x.Add(128)
	x.AddAll(25, 76, 90)
	
	//!+note
	fmt.Println(&x)         // "{1 9 42 144}"
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x)          // "{[4398046511618 0 65536]}"
	//!-note
	fmt.Println(x.Len())	// "9"
	x.Remove(42)
	fmt.Println(&x)			// "{1 9 25 37 76 90 128 144}"
	
	y := x.Copy()
	fmt.Println(y)			// "{1 9 25 37 76 90 128 144}"

	x.Clear()
	fmt.Println(&x)			// "{}"
	// Output:
	// {1 9 25 37 42 76 90 128 144}
	// {1 9 25 37 42 76 90 128 144}
	// {[4535519019522 67112960 65537]}
	// 9
	// {1 9 25 37 76 90 128 144}
	// {1 9 25 37 76 90 128 144}
	// {}
}
