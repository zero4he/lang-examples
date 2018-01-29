package main

import (
	"fmt"
	"net"
)

func main() {
	mxs, _ := net.LookupMX("baidu.com")
	for i := range mxs {
		fmt.Println(mxs[i].Host)
	}

	cname, err := net.LookupCNAME("baidu.com")
	if err == nil {
		fmt.Println(cname)
	} else {
		fmt.Println(err)
	}
}
