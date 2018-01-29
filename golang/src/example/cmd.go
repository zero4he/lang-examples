package main

import (
	"runtime"
	"fmt"
	"os"
	"flag"
	"os/signal"
	"time"
)
var (
	//m           = flag.String("m", "GET", "")
	//headers     = flag.String("h", "", "")
)

var usage = `Usage: hey [options...] <url>
Options:
  -n  Number of requests to run. Default is 200.
  -c  Number of requests to run concurrently. Total number of requests cannot
      be smaller than the concurrency level. Default is 50.
  -q  Rate limit, in seconds (QPS).
  -o  Output type. If none provided, a summary is printed.
      "csv" is the only supported alternative. Dumps the response
      metrics in comma-separated values format.
  -m  HTTP method, one of GET, POST, PUT, DELETE, HEAD, OPTIONS.
  -H  Custom HTTP header. You can specify as many as needed by repeating the flag.
      For example, -H "Accept: text/html" -H "Content-Type: application/xml" .
  -t  Timeout for each request in seconds. Default is 20, use 0 for infinite.
  -A  HTTP Accept header.
  -d  HTTP request body.
  -D  HTTP request body from file. For example, /home/user/file.txt or ./file.txt.
  -T  Content-type, defaults to "text/html".
  -a  Basic authentication, username:password.
  -x  HTTP Proxy address as host:port.
  -h2 Enable HTTP/2.
  -host	HTTP Host header.
  -disable-compression  Disable compression.
  -disable-keepalive    Disable keep-alive, prevents re-use of TCP
                        connections between different HTTP requests.
  -disable-redirects    Disable following of HTTP redirects
  -cpus                 Number of used cpu cores.
                        (default for current machine is %d cores)
  -more                 Provides information on DNS lookup, dialup, request and
                        response timings.
`

func main() {

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(usage, runtime.NumCPU()))
	}


	fmt.Println(time.Now())



	method := flag.String("m", "GET", "a string: get,post")
	fmt.Println(method)

	if flag.NArg() < 1 {
		flag.Usage()
		fmt.Fprintf(os.Stderr, "\n")
		os.Exit(1)
	}

	//var m string
	//flag.StringVar(&m, "m", "","")

	flag.Parse()



	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		os.Exit(1)
	}()



}
