package main

import (
	"net/http"
	"path/filepath"
	"flag"
	"fmt"
	"log"
)

func main() {
	//第一个docs是用户请的path，第二个docs是server删掉了request path中的前缀docs再映射到系统文件中
	//file, _ := exec.LookPath(os.Args[0])
	//exePath, _ := filepath.Abs(file)//可执行文件路径, filepath.Dir(exePath) 可执行文件所在目录
	dir := flag.String("dir", "www", "file dir")
	port := flag.Int("port", 8900, "file dir")
	flag.Parse()
	//println(*dir)
	absDir, _ := filepath.Abs(*dir);
	fmt.Printf("File dir: %s\n", absDir)
	fmt.Printf("Server port: %d\n", *port)
	fmt.Printf("Server URL:http://localhost:%d/\n",*port)

	http.Handle("/", http.FileServer(http.Dir(absDir)))
	//http.Handle("/docs/", http.StripPrefix("/docs/", http.FileServer(http.Dir(path))))

	//fmt.Sprintf(":%d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}