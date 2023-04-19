package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(`
This program will consume 1GB of ram memory for 30secs and then exit. Its usefull to test container memory-limits
Usage to consume 2GB ram: ./thisprogram -size 2147483648
`)
	var memorySize int
	flag.IntVar(&memorySize, "size", 1*1000*1024*1024, "Memory size in bytes")
	flag.Parse()

	data := make([]byte, memorySize)
	rand.Seed(time.Now().UnixNano())
	rand.Read(data) // Fill the data slice with random bytes

	fmt.Println("Allocated and filled", len(data), "bytes of memory with random data")
	time.Sleep(30 * time.Second) // Keep the program running for 30 seconds
}
