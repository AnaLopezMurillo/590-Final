package main

import (
	"fmt"
	"math"
	"math/rand/v2"
	"os"
	"time"
)

func main() {
	inCh := make(chan int, 5)
	outCh := make(chan int, 5)

	go prodOdd(inCh)
	go prodEven(inCh)
	go cons1(<-inCh, outCh, inCh)
	go cons2(<-inCh, outCh, inCh)
	go finalFilter(<-outCh, outCh)

	keep_running()

}

func keep_running() {
	// neverending loop
	for {
		time.Sleep(time.Second)
	}
}

func prodOdd(inCh chan int) {
	for i := 1; i <= 29; i += 2 {
		inCh <- i
		time.Sleep(time.Millisecond * time.Duration(rand.IntN(1500)))
	}
}

func prodEven(inCh chan int) {
	for i := 2; i <= 30; i += 2 {
		inCh <- i
		time.Sleep(time.Millisecond * time.Duration(rand.IntN(1500)))
	}
}

func cons1(i int, outCh chan int, inCh chan int) {
	result := math.Pow(float64(i), 2) // square the res
	outCh <- int(result)              // give to out channel
	time.Sleep(time.Millisecond * time.Duration(rand.IntN(3)))

	cons1(<-inCh, outCh, inCh)

}

func cons2(i int, outCh chan int, inCh chan int) {
	result := math.Pow(float64(i), 2) // square the res
	outCh <- int(result)              // give to out channel
	time.Sleep(time.Millisecond * time.Duration(rand.IntN(3)))

	cons2(<-inCh, outCh, inCh)

}

func finalFilter(i int, outCh chan int) {
	fmt.Println(i)
	for {
		new := <-outCh
		if new > i {
			fmt.Println(new)
			i = new
		}
		// exit if we finally reach 30^2 (no num greater than this)
		if new == 900 {
			os.Exit(0)
		}
	}
}
