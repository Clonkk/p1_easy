package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

const FoaasVersion = "1.0.0"

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func version(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, FoaasVersion)
}

func main() {
	port := flag.Int("port", 19050, "local http port")

	flag.Parse()
	if flag.NFlag() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	args := flag.Args()
	fmt.Println(args)
	fmt.Println("Running on port: ", *port)

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/version", version)

	portString := fmt.Sprintf(":%v", *port)
	http.ListenAndServe(portString, nil)

}
