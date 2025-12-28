package main

import (
	"fmt"
	"math/rand"
	"sync"
)



func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go randNumbers(ch, &wg)


	wg.Add(1)
	go powNumbers(ch, &wg)
	
	wg.Wait()


}

func randNumbers(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		ch <- rand.Intn(100)
	}
	close(ch)
}




func powNumbers(ch chan int, wg *sync.WaitGroup)  {
	defer wg.Done()
	for v := range ch {
		fmt.Println(v * v)
	}
}