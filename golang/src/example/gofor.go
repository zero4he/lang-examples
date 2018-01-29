package main

import (
	"fmt"
	"runtime"
)

func main() {
	for i := 0; i < 1000; i++ {
		go func() {
			for {
				fmt.Println("go2 ")
			}
		}()
	}
	fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())
	go func() {
		for {
			fmt.Println("go ")
		}
	}()
	fmt.Println("----")

}



