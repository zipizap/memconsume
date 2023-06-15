package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Print(`
This program will consume 1GB of ram memory for 30secs and then exit. Its usefull to test container memory-limits
Usage:
   #to consume 1GB ram for 30s
   ./thisprogram 

   #to consume 2GB ram for 60s
   ./thisprogram -size 2147483648 -secs 60
`)
	var memorySize, secs int
	flag.IntVar(&memorySize, "size", 1*1000*1024*1024, "Memory size in bytes")
	flag.IntVar(&secs, "secs", 30, "Duration in seconds before exiting")
	flag.Parse()

	data := make([]byte, memorySize)
	rand.Seed(time.Now().UnixNano())
	rand.Read(data) // Fill the data slice with random bytes

	fmt.Println("Allocated and filled", len(data), "bytes of memory with random data")
	fmt.Printf("Sleeping %d seconds\n", secs)
	time.Sleep(time.Duration(secs) * time.Second) // Keep the program running for secs seconds
}
