package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	chRand := make(chan int)
	chPow := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go randNumbers(chRand, &wg)

	wg.Add(1)
	go powNumbers(chRand, chPow, &wg)
	
	for res := range chPow {
		fmt.Println(res)
	}
	
	wg.Wait()
}


func randNumbers(chRand chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		chRand <- rand.Intn(100)
	}
	close(chRand)
}


func powNumbers(chRand chan int, chPow chan int, wg *sync.WaitGroup)  {
	defer wg.Done()
	for v := range chRand {
		chPow <- v * v
	}
	close(chPow)
}