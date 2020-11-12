package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func SayGreetings(greeting string, times int) {
	for i := 0; i < times; i++ {
		log.Println(greeting)
		d := time.Second * time.Duration(rand.Intn(5)) / 2
		time.Sleep(d)
	}
	// Notify a task is finished.
	wg.Done() // <=> wg.Add(-1)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)
	wg.Add(2) // register two tasks.
	go SayGreetings("hi!", 10)
	go SayGreetings("hello!", 10)
	wg.Wait() // block until all tasks are finished.
	log.Println("All goroutine finished")
}
