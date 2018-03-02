package main

import (
	"flag"
	"log"
	"os"
	"strings"

	"github.com/client9/csstool"
)

var (
	flagTags   = flag.String("tags", "", "html elements or classes to keep or remove")
	flagIndent = flag.Int("indent", 2, "indent spacing")
	flagTab    = flag.Bool("tab", false, "use tabs for indenting")
)

func main() {
	flag.Parse()
	tags := strings.Split(*flagTags, ",")

	cssformat := csstool.NewCSSFormat(*flagIndent, *flagTab, tags)

	err := cssformat.Format(os.Stdin, os.Stdout)
	if err != nil {
		log.Printf("FAIL: %s", err)
	}
}
