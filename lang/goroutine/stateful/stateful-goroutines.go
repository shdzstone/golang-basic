package main

import (
	"log"
	"math/rand"
	"sync/atomic"
	"time"
)

//读操作
type readOp struct {
	key  int
	resp chan int
}

//写操作
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {

	//读写操作计数器
	var readOps uint64 = 0
	var writeOps uint64 = 0

	//读写通道，用以发出读请求和写请求
	reads := make(chan *readOp)
	writes := make(chan *writeOp)

	//goroutine
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				if state[read.key] != 0 {
					log.Println("read key:", read.key, "read val:", state[read.key])
					read.resp <- state[read.key]
				}
			case write := <-writes:
				log.Println("write key:", write.key, "write val:", write.val)
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := &readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				log.Println("readOp send read request with key:", read.key)
				val := <-read.resp
				log.Println("readOp received response value:", val)
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := &writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Let the goroutines work for a second.
	time.Sleep(time.Millisecond * 1000)

	// Finally, capture and report the op counts.
	readOpsFinal := atomic.LoadUint64(&readOps)
	log.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	log.Println("writeOps:", writeOpsFinal)
}
