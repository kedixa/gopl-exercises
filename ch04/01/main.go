package main

import (
	"crypto/sha256"
	"fmt"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// count 1s in binary byte x
func popCount(x byte) int {
	return int(pc[x])
}

// count the different bits of two [32]byte
func diffBits(s1, s2 [32]byte) int {
	s := 0
	for i := 0; i < 32; i++ {
		s += popCount(s1[i] ^ s2[i])
	}
	return s
}
func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\ndiff bits: %d\n", c1, c2, diffBits(c1, c2))
}
