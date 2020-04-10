// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 173.

// Bytecounter demonstrates an implementation of io.Writer that counts bytes.
package main

import (
	"fmt"
	"bufio"
)

//!+bytecounter
// ByteCounter counts the number of bytes
type ByteCounter int

// WordCounter count words
type WordCounter int

// LineCounter count Lines
type LineCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

type scanFunc func(p []byte, EOF bool) (advance int, token []byte, err error)

func scanBytes(p []byte, fn scanFunc) (cnt int) {
    for true {
        advance, token, _ := fn(p, true)
        if len(token) == 0 {
            break
        }
        p = p[advance:]
        cnt++
    }
    return cnt
}

func (c *WordCounter) Write(p []byte) (int, error) {
    cnt := scanBytes(p, bufio.ScanWords)
    *c += WordCounter(cnt)
    return cnt, nil
}

func (c WordCounter) String() string {
    return fmt.Sprintf("contains %d words", c)
}

func (c *LineCounter) Write(p []byte) (int, error) {
    cnt := scanBytes(p, bufio.ScanLines)
    *c += LineCounter(cnt)
    return cnt, nil
}

func (c LineCounter) String() string {
    return fmt.Sprintf("contains %d lines", c)
}

//!-bytecounter

func main() {
	//!+main
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c) // "5", = len("hello")

	c = 0 // reset the counter
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // "12", = len("hello, Dolly")
	//!-main
	var w WordCounter
	w.Write([]byte("hello world"))
    fmt.Println(w)

    fmt.Fprintf(&w, "This is an sentence.")
    fmt.Println(w)

    w = 0
    fmt.Fprintf(&w, "This")
    fmt.Println(w)

	var l LineCounter
	l.Write([]byte("hello world\nwelcome"))
    fmt.Println(l)

    fmt.Fprintf(&l, `This is another line`)
    fmt.Println(l)

    l = 0
    fmt.Fprintf(&l, "This is another\nline")
    fmt.Println(l)

    fmt.Fprintf(&l, "This is one line")
    fmt.Println(l)
}
