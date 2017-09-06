package main

import (
	"flag"
	"fmt"
	"os"
)

type Arguments struct {
	quiet     bool
	port      int
	host      string
	log       string
	directory string
	cors      bool
	noCache   bool
}

func exitOnError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func parseArguments(args *Arguments) {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s: [-host HOST] [-port PORT] [DIRECTORY]\n"+
			"Serve directory if specified or a list of files specified from stdin if not.\n\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.BoolVar(&args.quiet, "quiet", false, "Quiet mode")
	flag.IntVar(&args.port, "port", 8080, "Port to listen")
	flag.StringVar(&args.host, "host", "localhost", "Host to listen")
	flag.StringVar(&args.log, "log", "", "Optional log file")
	flag.BoolVar(&args.cors, "cors", false, "Elable cross-origin resource sharing")
	flag.BoolVar(&args.noCache, "no-cache", false, "Disable caching")
	flag.Parse()

	if flag.NArg() > 1 {
		flag.Usage()
		os.Exit(1)
	}

	args.directory = flag.Arg(0)
	if args.directory != "" {
		_, err := os.Stat(args.directory)
		exitOnError(err)
	}
}
