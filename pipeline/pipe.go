package pipeline

import "fmt"

func Gen(done <-chan struct{}, nums ...int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for _, v := range nums {
			select {
			case ch <- v:
			case <-done:
				return
			}
		}
	}()
	return ch
}

func Sq(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := range in {
			select {
			case out <- i * i:
			case <-done:
				return
			}
		}
	}()
	return out
}

func Consumer() {
	done := make(chan struct{})

	g := Gen(done, 2, 3, 4, 5, 6, 7)
	c1 := Sq(done, g)
	c2 := Sq(done, g)

	fmt.Println(<-c1)
	fmt.Println(<-c2)
	close(done)
	fmt.Println(<-c2)
}
