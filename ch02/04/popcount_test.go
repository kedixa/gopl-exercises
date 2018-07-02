package popcount

import (
	"testing"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(0x1122334455667788)
	}
}

func BenchmarkPopUsingShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopUsingShift(0x1122334455667788)
	}
}
