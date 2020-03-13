// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 43.
//!+

// Cf converts its numeric argument to Celsius and Fahrenheit.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"gopl.io/ch2/tempconv"
)

func main() {
	if len(os.Args[1:]) > 1{
		for _, arg := range os.Args[1:] {
			t, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cf: %v\n", err)
				os.Exit(1)
			}
			f := tempconv.Fahrenheit(t)
			c := tempconv.Celsius(t)
			fmt.Printf("%s = %s, %s = %s\n",
				f, tempconv.FToC(f), c, tempconv.CToF(c))
		}
	} else{
		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Println("enter the value for conversion or press q to quit")
			key, err := reader.ReadString('\n')
			key = key[:len(key)-1]
			if err == nil && key != "q" {
				fmt.Println("the key is : ", key)
				value, err := strconv.Atoi(key)
				fmt.Println("the error found was :", err)
				if err == nil {
					f := tempconv.Fahrenheit(value)
					c := tempconv.Celsius(value)
					fmt.Printf("%s = %s, %s = %s\n\n",
						f, tempconv.FToC(f), c, tempconv.CToF(c))
				} else{
					fmt.Println("could not convert the entered value to integer")
				}
			} else if err != nil {
				fmt.Println("the entered value could not be read. please enter other value of press q to quit")
			} else if key == "q"{
				break
			}
			
		}
	}
	
}

//!-
