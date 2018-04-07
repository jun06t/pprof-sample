package main

import (
	"fmt"
	"log"
	"net/http"

	_ "net/http/pprof"
)

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	fmt.Println("start")
	for {
		fmt.Println(fib(25))
		fmt.Println(fib(26))
		fmt.Println(fib(27))
		fmt.Println(fib(28))
		fmt.Println(fib(29))
		fmt.Println(fib(30))
		fmt.Println(fib(31))
	}
}
