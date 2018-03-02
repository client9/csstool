package csstool

import (
	"io"
	"sort"
	"strings"

	"golang.org/x/net/html"
)

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
		if tt == html.ErrorToken {
			err := z.Err()
			if err == io.EOF {
				return nil
			}
			return err
		}
		if tt == html.StartTagToken {
			tname, hasA := z.TagName()
			c.counter[string(tname)]++
			var key, val []byte
			for hasA {
				key, val, hasA = z.TagAttr()
				switch string(key) {
				case "class":
					classes := string(val)
					for _, cname := range strings.Fields(classes) {
						c.counter["."+cname]++
					}
				case "id":
					c.counter["#"+string(val)]++
				}
			}
		}
	}
}
