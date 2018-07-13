package main

const m = 10000000

func main() {
	done := make(chan bool)
	// two chans without buffer
	ch1 := make(chan int)
	ch2 := make(chan int)
	// start communication
	go func() {
		for i := 0; i < m; i++ {
			ch2 <- i
			<-ch1
		}
		close(ch2)
		done <- true
	}()
	go func() {
		for i := range ch2 {
			ch1 <- i
		}
		done <- true
	}()
	<-done
	<-done
}
