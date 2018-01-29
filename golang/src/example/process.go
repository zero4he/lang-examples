package main

import (
	"os"
	"fmt"
)

func main() {
	ss,_:=os.FindProcess(3320)

	fmt.Print(ss.Pid)
}
