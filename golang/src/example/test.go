package main

import (
	"github.com/paulrosania/go-charset/charset"
	_ "github.com/paulrosania/go-charset/data"
	"fmt"
	"github.com/labstack/gommon/log"
	"bytes"
	"strings"
	"io/ioutil"
)

func main() {
	//编码
	buf := new(bytes.Buffer)
	w, err := charset.NewWriter("utf-8", buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "你好")
	w.Close()
	fmt.Printf("%q\n", buf.Bytes())

	//解码
	r,err:= charset.NewReader("utf-8",strings.NewReader(string(buf.Bytes())))
	if err!= nil {
		log.Fatal(err)
	}
	result, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", result)

}
