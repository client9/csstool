package csstool

import (
	"github.com/tdewolff/parse/css"
)

func selectors(tokens []css.Token) []string {
	out := []string{}
	buf := []byte{}
	hasClass := false
	halt := false
	skipTillDelim := false
	for _, t := range tokens {
		if halt {
			break
		}
		if skipTillDelim && (t.TokenType != css.WhitespaceToken && t.TokenType != css.DelimToken) {
			continue
		}
		if skipTillDelim && t.TokenType == css.DelimToken && t.Data[0] == '.' {
			continue
		}
		skipTillDelim = false

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
			skipTillDelim = true
		case css.LeftBracketToken:
			// we only handle case of raw
			// [ foo ]
			// not
			// a[foo]
			if len(buf) == 0 {
				buf = append(buf, t.Data...)
				continue
			}
			halt = true
		case css.RightBracketToken:
			buf = append(buf, t.Data...)
			halt = true
		case css.WhitespaceToken:
			// really a delimiter
			out = append(out, string(buf))
			buf = []byte{}
			hasClass = false
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
					out = append(out, string(buf))
					buf = []byte{}
				}
				hasClass = true
				buf = append(buf, t.Data...)
				continue
			}
			if t.Data[0] == '=' {
				buf = append(buf, t.Data...)
				continue
			}
			if t.Data[0] == '*' {
				buf = append(buf, t.Data...)
			}
			if len(buf) > 0 {
				out = append(out, string(buf))
				buf = []byte{}
			}
			hasClass = false
		default:
			buf = append(buf, t.Data...)
		}
	}
	if len(buf) > 0 {
		out = append(out, string(buf))
	}
	return out
}

// gives the complete identier from tokens
func completeSelector(tokens []css.Token) []byte {
	buf := []byte{}
	for _, t := range tokens {
		buf = append(buf, t.Data...)
	}
	return buf
}
