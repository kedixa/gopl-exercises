package measure

import (
	"testing"
)

func BenchmarkMethodInefficient(b *testing.B) {
	b.StopTimer()
	// make testing data
	s := "abcdefghijklmnopqrstuvwxyz"
	var args []string
	for i := 0; i < 30; i++ {
		args = append(args, s)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		MethodInefficient(args)
	}
}

func BenchmarkMethodJoin(b *testing.B) {
	b.StopTimer()
	// make testing data
	s := "abcdefghijklmnopqrstuvwxyz"
	var args []string
	for i := 0; i < 30; i++ {
		args = append(args, s)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		MethodJoin(args)
	}
}
