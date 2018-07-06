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

// IntersectionWith sets s to the intersection of s and t
func (s *IntSet) IntersectionWith(t *IntSet) {
	if len(s.words) > len(t.words) {
		s.words = s.words[:len(t.words)]
	}
	for i := 0; i < len(s.words); i++ {
		s.words[i] &= t.words[i]
	}
}

// DifferenceWith sets s to the difference of s and t
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tword
		} else {
			break
		}
	}
}

// SymmetricDifference sets s to the symmetric difference of s and t
func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
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

// AddAll add all elements to s
func (s *IntSet) AddAll(xs ...int) {
	for _, x := range xs {
		s.Add(x)
	}
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
	var t, r IntSet
	var s *IntSet
	r.AddAll(1, 4, 8, 66, 155)
	t.AddAll(2, 4, 8, 9, 66, 67)

	s = r.Copy()
	s.IntersectionWith(&t)
	fmt.Println(s) // {4 8 66}

	s = r.Copy()
	s.DifferenceWith(&t)
	fmt.Println(s) // {1 155}

	s = r.Copy()
	s.SymmetricDifference(&t)
	fmt.Println(s) // {1 2 9 67 155}

}
