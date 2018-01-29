package main

import "fmt"

func main() {
	ch := make(chan int)
	quit := make(chan int)

	defer close(ch)
	defer close(quit)

	go func() {
		quit <- 0
		ch <- 0
	}()

	//select 获取到任意一个值后会结束
	select {
	case c := <-ch:
		fmt.Println("-----> ", c)

	case <-quit:
	//q:=quit , 这里因为没有用到q变量,所以可以直接写成
		fmt.Println("quit")
	}
}