// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 165.

// Package intset provides a set of integers based on a bit vector.
package intset

import (
	"bytes"
	"fmt"
	"math/bits"
)

//!+intset

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

//!-intset

//!+string

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

//!-string

// Len returns the number of elements in the set
func (s *IntSet) Len() int {// return the number of elements
	var bitCount int
	for _, word := range s.words{
		bitCount += bits.OnesCount64(word)
	}
	return bitCount
}

// Remove removes the specified element from the set
func (s *IntSet) Remove(x int){ // remove x from the set
	word, bit := x/64, uint(x%64)
	s.words[word] = s.words[word] &^ (1<<bit)
}

// Clear clears all the elements of the set
func (s *IntSet) Clear(){ // remove all elements from the set
	s.words = nil
}

// Copy makes a copy of the struct IntSet object 
// and returns the address of new IntSet object
func (s *IntSet) Copy() *IntSet {// return a copy of the set
	var i IntSet
	i.words = make([]uint64, len(s.words))
	copy(i.words, s.words)
	return &i
}

// AddAll is a variadic inputs to Add many values at a time	
func (s *IntSet) AddAll(nums ...int) { 
    for _, num := range nums {
		s.Add(num)
    }
}

// IntersectWith sets s to the intersection of s and t.
func (s *IntSet) IntersectWith(t *IntSet) {
	var minSetLength int
	sWordLen, tWordLen := len(s.words), len(t.words)
	if sWordLen > tWordLen{
		minSetLength = tWordLen
	} else {
		minSetLength = sWordLen
	}
	for i := 0; i < sWordLen; i++ {
		if i < minSetLength {
			s.words[i] &= t.words[i]
		} else {
			s.words[i] = 0
		}
	}
}

// DifferenceWith sets s to the difference of s and t.
func (s *IntSet) DifferenceWith(t *IntSet) {
	// s  t  d
	// 0  0  0
	// 0  1  0
	// 1  0  1
	// 1  1  0

	var minSetLength int
	if sWordLen, tWordLen := len(s.words), len(t.words); sWordLen > tWordLen{
		minSetLength = tWordLen
	} else {
		minSetLength = sWordLen
	}
	for i := 0; i < minSetLength; i++ {
		s.words[i] &= (^t.words[i])
	}
}

// Elems returns the a slice with set elements of the form "[1 2 3]".
func (s *IntSet) Elems() *[]int {
	var set []int
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				set = append(set, 64*i + j)
			}
		}
	}
	return &set
}