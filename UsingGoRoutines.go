package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	start := time.Now()
	wg.Add(2)
	go doSomething1()
	go doSomethingElse1()
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("Processes took: %s \n", elapsed)
}

func doSomething1() {
	time.Sleep(time.Second * 2)
	fmt.Println("Done something...")
	wg.Done()
}

func doSomethingElse1() {
	time.Sleep(time.Second * 2)
	fmt.Println("Done something else...")
	wg.Done()
}
