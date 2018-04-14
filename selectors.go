package csstool

import (
	"github.com/tdewolff/parse/css"
)

func primarySelector(tokens []css.Token) []byte {
	buf := []byte{}
	hasClass := false
	for _, t := range tokens {
		switch t.TokenType {
		case css.IdentToken, css.HashToken:
			buf = append(buf, t.Data...)
			continue
		case css.ColonToken:
			// three cases,
			// foo:hover
			// ::after
			// :root
			if len(buf) == 0 || (len(buf) == 1 && buf[0] == ':') {
				buf = append(buf, t.Data...)
				continue
			}
			return buf
		case css.LeftBracketToken:
			// we only handle case of raw
			// [ foo ]
			// not
			// a[foo]
			if len(buf) == 0 {
				buf = append(buf, t.Data...)
				continue
			}
			return buf
		case css.RightBracketToken:
			buf = append(buf, t.Data...)
			return buf
		case css.WhitespaceToken:
			// really a delimiter
			return buf
		case css.DelimToken:
			if len(t.Data) != 1 {
				panic("got delim with len > 1")
			}
			if t.Data[0] == '.' {
				// only allow one class:
				// .foo.bar --> .foo
				// if .foo doesn't exist then .foo.bar won't either
				// it could be .foo exists but .bar doesn't but
				// are not optimizing that
				if hasClass {
					return buf
				}
				hasClass = true
				buf = append(buf, t.Data...)
				continue
			}
			if t.Data[0] == '*' {
				buf = append(buf, t.Data...)
				continue
			}
			return buf
		default:
			buf = append(buf, t.Data...)
		}
	}
	return buf
}

// gives the complete identier from tokens
func completeSelector(tokens []css.Token) []byte {
	buf := []byte{}
	for _, t := range tokens {
		buf = append(buf, t.Data...)
	}
	return buf
}
