package main

import (
	"bytes"
	"fmt"
)

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

// Len returns the number of elements in s
func (s *IntSet) Len() int {
	n := 0
	var bitCount = func(u uint64) (x int) {
		for u != 0 {
			x++
			u = u & (u - 1)
		}
		return
	}
	for _, w := range s.words {
		n += bitCount(w)
	}
	return n
}

// Remove remove the xth element in s
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if word < len(s.words) {
		s.words[word] &= ^(1 << bit)
	}
}

// Clear clear all elements in s
func (s *IntSet) Clear() {
	s.words = s.words[:0]
}

// Copy return deep copy of s
func (s *IntSet) Copy() *IntSet {
	var ss IntSet
	ss.words = make([]uint64, len(s.words))
	for i, w := range s.words {
		ss.words[i] = w
	}
	return &ss
}

func main() {
	s := IntSet{}
	for _, w := range []int{1, 3, 5, 8, 10} {
		s.Add(w)
	}
	fmt.Println(&s) // {1 3 5 8 10}
	s.Remove(8)
	fmt.Println(&s) // {1 3 5 10}

	t := s.Copy()
	fmt.Println(t) // {1 3 5 10}

	t.Remove(10)
	fmt.Println(&s) // {1 3 5 10}
	fmt.Println(t)  // {1 3 5}

	fmt.Println(t.Has(10), s.Has(10)) // false true

	t.Clear()
	fmt.Println(t) // {}
}
