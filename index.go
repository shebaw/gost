package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type indexHandler struct {
	dir http.Dir
}

// yanked out of golang's library
var htmlReplacer = strings.NewReplacer(
	"&", "&amp;",
	"<", "&lt;",
	">", "&gt;",
	// "&#34;" is shorter than "&quot;".
	`"`, "&#34;",
	// "&#39;" is shorter than "&apos;" and apos was not in HTML until HTML5.
	"'", "&#39;",
)

const indexPage = "/index.html"

func (fs *indexHandler) Open(name string) (http.File, error) {
	log.Printf("dir: %s", args.directory)
	log.Printf("name: %s", name)
	if args.directory != "" || name != indexPage {
		return fs.dir.Open(name)
	}
	var b bytes.Buffer
	// build the index page from the standard input instead
	fmt.Fprintf(&b, "<pre>\n")
	// loop through the standard input line by line and add it as
	// an entry
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		name = s.Text()
		log.Println(name)
		// name may contain '?' or '#', which must be escaped to remain
		// part of the URL path, and not indicate the start of a query
		// string or fragment.
		url := url.URL{Path: name}
		fmt.Fprintf(&b, "<a href=\"%s\">%s</a>\n", url.String(), htmlReplacer.Replace(name))
	}
	fmt.Fprintf(&b, "</pre>\n")
	return &memFile{bytes.NewReader(b.Bytes()), time.Now()}, nil
}

// implements http.File
type memFile struct {
	*bytes.Reader
	modTime time.Time
}

// only need for index.html
func (f *memFile) Stat() (os.FileInfo, error) {
	return &memOSFile{indexPage,
		f.Size(),
		0,
		f.modTime,
		false}, nil
}

func (f *memFile) Readdir(count int) ([]os.FileInfo, error) {
	// will not be called, here to complete the http.FileInterace
	return nil, nil
}

func (f *memFile) Close() error {
	return nil
}

// implements os.FileInfo
type memOSFile struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
	isDir   bool
}

func (f *memOSFile) Name() string {
	return f.name
}

func (f *memOSFile) Size() int64 {
	return f.size
}

func (f *memOSFile) Mode() os.FileMode {
	return f.mode
}

func (f *memOSFile) ModTime() time.Time {
	return f.modTime
}

func (f *memOSFile) IsDir() bool {
	return f.isDir
}

func (f *memOSFile) Sys() interface{} {
	return nil
}
