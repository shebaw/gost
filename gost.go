package main

import (
	"runtime"
	"os"
	"fmt"
	"log"
	"net/http"
)

const Version = "0.1.2"

var args Arguments

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	parseArguments(&args)

	if len(args.log) > 0 {
		file, err := os.OpenFile(args.log, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
		exitOnError(err)
		defer file.Close()
		log.SetOutput(file)
	} else if args.quiet {
		file, err := os.Open(os.DevNull)
		exitOnError(err)
		defer file.Close()
		log.SetOutput(file)
		log.SetFlags(0)
	}
}

func main() {
	listen := fmt.Sprintf("%s:%d", args.host, args.port)

	http.Handle("/", http.FileServer(http.Dir(args.directory)))
	handler := buildHttpHandler()

	log.Printf("Static file server running at %s. Ctrl+C to quit.\n", listen)
	err := http.ListenAndServe(listen, handler)
	if err != nil {
		log.Fatalln(err)
	}
}

func buildHttpHandler() http.Handler {
	var handler http.Handler

	handler = http.DefaultServeMux

	if args.cors {
		handler = corsHandler(handler)
	}

	handler = cacheHandler(handler)

	if !args.quiet {
		handler = logHandler(handler)
	}

	return handler
}
