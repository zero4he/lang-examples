package main

import (
	"fmt"
	"reflect"
)

func main() {
	s:="中文"
	fmt.Println(s)
	fmt.Println(reflect.TypeOf(s).Kind())
}
