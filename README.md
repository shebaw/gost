# Gost Static HTTP File Server

> A static HTTP file server written in Golang.

## Install

```
$ go get github.com/vwochnik/gost
```

## Usage

Serve the current directory on port `8080`:

```
$ gost
Static file server running at localhost:8080. Ctrl+C to quit.
```

Server your home directory on port `8888`:

```
gost -port 8888 ~/
Static file server running at localhost:8888. Ctrl+C to quit.
```

See the help:

```
$ gost -h
Usage of gost: [-host HOST] [-port PORT] [DIRECTORY]

Serves the directory specified by the first argument which defaults to the current working directory if not specified.

  -cors
        Elable cross-origin resource sharing
  -host string
        Host to listen (default "localhost")
  -log string
        Optional log file
  -no-cache
        Disable caching
  -port int
        Port to listen (default 8080)
  -quiet
        Quiet mode
```

# Copyright

Copyright (c) 2016 Vincent Wochnik.

License: MIT
