package main

import (
	"fmt"
	"log"
	"net/http"

	_ "net/http/pprof"
)

func fib() func() int {
	a, b := 0, 1

	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	f := fib()
	fmt.Println("start")
	for {
		f()
	}
}
