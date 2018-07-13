package popcount

import (
	"fmt"
	"testing"
)

func TestPopCount(t *testing.T) {
	done := make(chan struct{})

	go func() {
		fmt.Println("PopCount(1): ", PopCount(1))
		done <- struct{}{}
	}()
	go func() {
		fmt.Println("PopCount(10): ", PopCount(10))
		done <- struct{}{}
	}()
	<-done
	<-done
}
