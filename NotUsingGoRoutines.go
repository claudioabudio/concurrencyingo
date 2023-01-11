package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	doSomething()
	doSomethingElse()
	elapsed := time.Since(start)
	fmt.Printf("Processes took: %s \n", elapsed)
}

func doSomething() {
	time.Sleep(time.Second * 2)
	fmt.Println("Done something...")
}

func doSomethingElse() {
	time.Sleep(time.Second * 2)
	fmt.Println("Done something else...")
}
