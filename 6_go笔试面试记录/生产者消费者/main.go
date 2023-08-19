package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func producer(c chan<- int, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	defer close(c)
	for {
		time.Sleep(1 * time.Second)
		c <- rand.Intn(100) + 1
	}
}

func consumer(c <-chan int, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	for {
		num := <-c
		fmt.Println(num)
	}
}

func main() {
	var wg sync.WaitGroup
	cc := make(chan int, 5)
	defer close(cc)

	go consumer(cc, &wg)
	go consumer(cc, &wg)

	go producer(cc, &wg)
	go producer(cc, &wg)
	go producer(cc, &wg)

	wg.Wait()
}
