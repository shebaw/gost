package main

import (
	"os"
	"fmt"
	"flag"
)

type Arguments struct {
	quiet bool
	port int
	host string
	log string
	directory string
}

func exitOnError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func parseArguments(args *Arguments) {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s: [-host HOST] [-port PORT] [DIRECTORY]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Serves the directory specified by the first argument " +
												 "which defaults to the current working directory if " +
												 "not specified.\n\n")
		flag.PrintDefaults()
	}

	flag.BoolVar(&args.quiet, "quiet", false, "Quiet mode")
	flag.IntVar(&args.port, "port", 8080, "Port to listen")
	flag.StringVar(&args.host, "host", "localhost", "Host to listen")
	flag.StringVar(&args.log, "log", "", "Optional log file")
	flag.Parse()

	if flag.NArg() > 1 {
		flag.Usage()
		os.Exit(1)
	}

	args.directory = flag.Arg(0)
	if len(args.directory) == 0 {
		args.directory = "."
	}

	_, err := os.Stat(args.directory)
	exitOnError(err)
}
