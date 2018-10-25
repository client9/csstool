package csstool

import (
	"io"
	"log"
	"sort"
	"strings"

	"golang.org/x/net/html"
)

// tags that should be ignore when generating
// special href selectors
var noAttrSelectors = map[string]bool{
	"link":   true,
	"style":  true,
	"script": true,
	"meta":   true,
	"html":   true,
}

// CSSCount is for keeping a running frequency of CSS identifiers
type CSSCount struct {
	counter map[string]int
}

// NewCSSCount returns an initialized CSSCount object
func NewCSSCount() *CSSCount {
	c := CSSCount{}
	c.Reset()
	return &c
}

// Reset return object to initial state
func (c *CSSCount) Reset() {
	c.counter = make(map[string]int)
}

// Counts returns a map of identifers and their frequency counts
//
// NOTE: returns internal object, not a copy
//
func (c *CSSCount) Counts() map[string]int {
	return c.counter
}

// List returns a sort list of identifiers found
func (c *CSSCount) List() []string {
	out := make([]string, 0, len(c.counter))
	for k := range c.counter {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

// Add frequency counts of CSS identifiers from a input reader
func (c *CSSCount) Add(r io.Reader) error {
	z := html.NewTokenizer(r)
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			err := z.Err()
			if err == io.EOF {
				return nil
			}
			return err
		case html.StartTagToken, html.SelfClosingTagToken:
			tnamebytes, hasA := z.TagName()
			tname := string(tnamebytes)
			c.counter[tname]++

			var key, val []byte
			for hasA {
				key, val, hasA = z.TagAttr()
				switch string(key) {
				case "class":
					classes := string(val)
					for _, cname := range strings.Fields(classes) {
						log.Printf("Adding %s", cname)
						c.counter["."+cname]++
						c.counter[tname+"."+cname]++
					}
				case "id":
					c.counter["#"+string(val)]++
				default:
					// tags common in <head> should be ignored
					if noAttrSelectors[tname] {
						continue
					}
					keystr := string(key)
					// special href selectors
					c.counter[tname+"["+keystr+"]"]++
					c.counter["["+keystr+"]"]++
					if tname == "type" {
						c.counter["["+keystr+"="+string(val)+"]"]++
					}
				}
			}
		}
	}
}
