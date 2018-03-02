package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mattn/go-zglob"

	"github.com/client9/csstool"
)

var (
	flagHTML  = flag.String("html", "", "pattern for finding HTML files")
	flagDebug = flag.Bool("debug", false, "enable debug logging")
)

func main() {
	flag.Parse()

	c := csstool.NewCSSCount()
	if *flagHTML == "" {
		if *flagDebug {
			log.Printf("reading from stdin")
		}
		err := c.Add(os.Stdin)
		if err != nil {
			log.Fatalf("FAIL: %s", err)
		}
	} else {
		if *flagDebug {
			log.Printf("Using pattern %q", *flagHTML)
		}
		files, err := zglob.Glob(*flagHTML)
		if err != nil {
			log.Fatalf("FAIL: %s", err)
		}
		for _, f := range files {
			log.Printf("reading %s", f)
			r, err := os.Open(f)
			if err != nil {
				log.Fatalf("FAIL: %s", err)
			}
			err = c.Add(r)
			if err != nil {
				log.Fatalf("FAIL: %s", err)
			}
			r.Close()
		}
	}
	fmt.Printf("%s\n", strings.Join(c.List(), ","))
}
