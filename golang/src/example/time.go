package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().Unix()) //1503308563
	//毫秒: 1502966080627 java中的Data.time()返回的就是毫秒
	fmt.Println(time.Now().Nanosecond()) //919484300
	fmt.Println(time.Now().UnixNano()) //1503308563919484300
	t := time.Unix(1502966080627 / 1000, 0)
	fmt.Println(t)
}
