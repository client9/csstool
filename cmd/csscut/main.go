package main

import (
	"flag"
	"log"
	"os"

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
		log.Fatalf("must specify html pattern")
	}
	if *flagDebug {
		log.Printf("using pattern %q", *flagHTML)
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

	// now get CSS file
	cf := csstool.NewCSSFormat(0, false, csstool.NewTagMatcher(c.List()))
	cf.Debug = *flagDebug
	err = cf.Format(os.Stdin, os.Stdout)
	if err != nil {
		log.Printf("FAIL: %s", err)
	}
}
