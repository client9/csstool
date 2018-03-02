package csstool

import (
	"bufio"
	"bytes"
	"io"

	"github.com/tdewolff/parse/css"
)

func getTags(tokens []css.Token) []string {
	out := []string{}
	first := true
	for _, tok := range tokens {
		if first && tok.TokenType == css.IdentToken {
			out = append(out, string(tok.Data))
			first = false
		}
		if tok.TokenType == css.WhitespaceToken {
			first = true
		}
	}
	return out
}

type stack []io.Writer

func (s stack) Push(v io.Writer) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, io.Writer) {
	l := len(s)
	return s[:l-1], s[l-1]
}

// CSSFormat contains formatting perferances for CSS
type CSSFormat struct {
	Indent          int
	IndentTab       bool
	AlwaysSemicolon bool
	Matcher         *TagMatch
}

// NewCSSFormat creates an initialized CSSFormat object
func NewCSSFormat(indent int, useTabs bool, tags []string) *CSSFormat {
	if useTabs {
		indent = 1
	}
	return &CSSFormat{
		Indent:    indent,
		IndentTab: useTabs,
		Matcher:   NewTagMatch(tags, true),
	}
}

// Format reformats CSS using a reader to output writer
func (c *CSSFormat) Format(r io.Reader, wraw io.Writer) error {
	var err error
	var w io.Writer
	writers := make(stack, 0)
	wbuf := bufio.NewWriter(wraw)
	w = wbuf
	//writers = writers.Push(wbuf)
	qualified := 0
	ruleCount := 0
	indent := 0
	skipRuleset := false
	rulesetCount := 0

	p := css.NewParser(r, false)
	for err == nil {
		gt, _, data := p.Next()
		switch gt {
		case css.ErrorGrammar:
			wbuf.Flush()
			if err == io.EOF {
				err = nil
			}
			return err

		// a comma-separated list of tags
		// but not the last one .. so h1,h2,h3
		// h1,h2 are here, but h3 is a beginRuleSetGrammar
		case css.QualifiedRuleGrammar:
			tokens := p.Values()
			if c.Matcher.Remove(tokens[0].Data) {
				continue
			}
			if qualified == 0 {
				c.addIndent(w, indent)
			} else {
				c.writeComma(w)
			}
			qualified++
			for _, t := range tokens {
				w.Write(t.Data)
			}
		case css.BeginRulesetGrammar:
			ruleCount = 0
			tokens := p.Values()
			if qualified == 0 {
				if c.Matcher.Remove(tokens[0].Data) {
					indent++
					skipRuleset = true
					continue
				}
				c.addIndent(w, indent)
				for _, t := range tokens {
					w.Write(t.Data)
				}
				c.writeLeftBrace(w)
				indent++
				continue
			}

			qualified = 0
			indent++
			if c.Matcher.Remove(tokens[0].Data) {
				c.writeLeftBrace(w)
				continue
			}
			c.writeComma(w)
			for _, t := range tokens {
				w.Write(t.Data)
			}
			c.writeLeftBrace(w)
		case css.EndRulesetGrammar:
			indent--
			if skipRuleset {
				skipRuleset = false
				continue
			}
			rulesetCount++

			// add semicolon, even if the last rule
			// i.e.  color: #000;}
			if c.AlwaysSemicolon {
				w.Write([]byte{';'})
			}
			c.addNewline(w)
			c.addIndent(w, indent)
			w.Write([]byte{'}'})
			c.addNewline(w)
		case css.BeginAtRuleGrammar:
			ruleCount = 0
			rulesetCount = 0

			// first render the @rule
			// into it's own buffer

			// save existing context
			writers = writers.Push(w)

			w = &bytes.Buffer{}
			c.addIndent(w, indent)
			w.Write(data)
			c.addSpace(w)
			tokens := p.Values()
			for i, tok := range tokens {
				if i > 0 {
					c.addSpace(w)
				}
				w.Write(tok.Data)
			}
			c.writeLeftBrace(w)

			// set up new buffer for content
			writers = writers.Push(w)
			w = &bytes.Buffer{}
			indent++
		case css.EndAtRuleGrammar:
			// have we written anything?
			contents := w.(*bytes.Buffer).Bytes()
			writers, w = writers.Pop()
			header := w.(*bytes.Buffer).Bytes()
			writers, w = writers.Pop()
			indent--
			if len(contents) == 0 {
				// no
				continue
			}
			w.Write(header)
			w.Write(contents)
			c.addIndent(w, indent)
			w.Write([]byte{'}'})
			c.addNewline(w)
		case css.CommentGrammar:
			w.Write(data)
			c.addNewline(w)
		case css.CustomPropertyGrammar:
			c.addIndent(w, indent)
			w.Write(data)
			// do not add space
			w.Write([]byte{':'})
			tokens := p.Values()
			for _, tok := range tokens {
				w.Write(tok.Data)
			}
			c.writeSemicolon(w)
		case css.DeclarationGrammar:
			if skipRuleset {
				continue
			}
			if ruleCount != 0 {
				c.writeSemicolon(w)
			}
			ruleCount++
			c.addIndent(w, indent)
			w.Write(data)
			w.Write([]byte{':'})
			c.addSpace(w)
			tokens := p.Values()
			for _, tok := range tokens {
				if len(tok.Data) == 1 && tok.Data[0] == '!' {
					c.addSpace(w)
					w.Write([]byte{'!'})
				} else {
					w.Write(tok.Data)
				}
			}
		case css.TokenGrammar:
			w.Write(data)
		default:
			panic("Unknown grammar")
		}
	}
	wbuf.Flush()
	return err
}

var (
	spaces = []byte("                  ")
	tabs   = []byte("\t\t\t\t")
)

func (c *CSSFormat) addIndent(w io.Writer, depth int) {
	if c.Indent == 0 || depth == 0 {
		return
	}
	if c.IndentTab {
		w.Write(tabs[:depth])
		return
	}

	w.Write(spaces[:c.Indent*depth])
}
func (c *CSSFormat) addSpace(w io.Writer) {
	if c.Indent == 0 {
		return
	}
	w.Write([]byte{' '})
}

func (c *CSSFormat) addNewline(w io.Writer) {
	if c.Indent == 0 {
		return
	}
	w.Write([]byte{'\n'})
}

func (c *CSSFormat) writeComma(w io.Writer) {
	if c.Indent == 0 {
		w.Write([]byte{','})
		return
	}
	w.Write([]byte{',', ' '})
}

func (c *CSSFormat) writeLeftBrace(w io.Writer) {
	if c.Indent == 0 {
		w.Write([]byte{'{'})
		return
	}
	w.Write([]byte{' ', '{', '\n'})
}

func (c *CSSFormat) writeSemicolon(w io.Writer) {
	if c.Indent == 0 {
		w.Write([]byte{';'})
		return
	}
	w.Write([]byte{';', '\n'})
}
