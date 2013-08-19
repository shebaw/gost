package main

import (
	"os"
	"flag"
	"fmt"
	"time"
	"net/http"
	"path/filepath"
)

const Version = "0.1.2"

var (
	port int
	host string
	directory string
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s: [-port 8080] [directory]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Serves the directory specified by the first argument " +
												 "which defaults to the current working directory if " +
												 "not specified.\n\n")
		flag.PrintDefaults()
	}

	flag.IntVar(&port, "port", 8080, "Port to listen")
	flag.StringVar(&host, "host", "localhost", "Host to listen")
	flag.Parse()

	directory = flag.Arg(0)
	if len(directory) == 0 {
		directory = "."
	}

	_, err := os.Stat(directory)
	if os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "Specified directory does not exist.\n\n")
		flag.Usage()
		os.Exit(1)
	} else if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
}

func main() {
	directory, err := filepath.Abs(directory)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}

	listen := fmt.Sprintf("%s:%d", host, port)
	fmt.Printf("Static file server running at %s. Ctrl+C to quit.\n", listen)

	http.Handle("/", http.FileServer(http.Dir(directory)))

	err = http.ListenAndServe(listen, logWrapper(http.DefaultServeMux))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func logWrapper(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		fmt.Fprintf(os.Stderr, "[%s] %s %s\n", now.Format(time.RFC850),
								r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
