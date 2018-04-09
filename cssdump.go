package csstool

import (
	"fmt"
	"io"

	"github.com/tdewolff/parse/css"
)

// Dump emits grammar info
//
// really only useful for debugging
func Dump(r io.Reader, w io.Writer) error {
	p := css.NewParser(r, false)
	for {
		gt, tt, data := p.Next()
		switch gt {
		case css.ErrorGrammar:
			return nil
		case css.CommentGrammar, css.EndAtRuleGrammar, css.EndRulesetGrammar:
			fmt.Printf("%s %s %s\n", gt, tt, data)
			continue
		default:
			tokens := p.Values()
			fmt.Printf("%s %s %s\n", gt, tt, data)
			for _, t := range tokens {
				fmt.Printf("    %s %q\n", t, t.Data)
			}
		}
	}
}
