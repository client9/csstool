package csstool

import (
	"bytes"
	"strings"
	"testing"
)

var testcases = []struct {
	css  string
	tags []string
	want string
}{
	{
		css:  "h1,h2,h3{color:#000}",
		tags: []string{"h1"},
		want: "h1{color:#000}",
	},
	{
		css:  "h1,h2,h3{color:#000}",
		tags: []string{"h2"},
		want: "h2{color:#000}",
	},
	{
		css:  `h1,h2,h3{color:#000}`,
		tags: []string{"h3"},
		want: "h3{color:#000}",
	},
	{
		css:  "h1,h2,h3{color:#000}",
		tags: []string{"h4"},
		want: "",
	},
	{
		css:  "h1,h2,h3{color:#000}",
		tags: []string{"h1", "h2"},
		want: "h1,h2{color:#000}",
	},
	{
		css:  "h1,h2,h3{color:#000}",
		tags: []string{"h1", "h3"},
		want: "h1,h3{color:#000}",
	},
	{
		css:  "h1,h2,h3{color:#000}",
		tags: []string{"h2", "h3"},
		want: "h2,h3{color:#000}",
	},
	{
		// test comments
		css:  "/* start */h1,h2,h3{color:#000}/* end */",
		tags: []string{"h1"},
		want: "/* start */h1{color:#000}/* end */",
	},
	{
		// test no ending semicolon
		css:  "h1,h2,h3{color:#000;background-color:#fff}",
		tags: []string{"h1"},
		want: "h1{color:#000;background-color:#fff}",
	},
	{
		// special case, "*", and "::tags" should always be kept
		css:  "*,::after,::before{box-sizing:border-box}",
		tags: []string{"h1"},
		want: "*,::after,::before{box-sizing:border-box}",
	},
	{
		// raw attribute selectors are preserved
		css:  "[hidden]{display:none!important}",
		tags: []string{"h1"},
		want: "[hidden]{display:none!important}",
	},
	{
		// if keep <a> then also keep a:something
		css:  "a:hover{color:#0056b3}",
		tags: []string{"a"},
		want: "a:hover{color:#0056b3}",
	},
	{
		// another special case with ":"
		css:  ".row:after{clean:both}",
		tags: []string{".row"},
		want: ".row:after{clean:both}",
	},
	{
		// specifier ">"
		css:  ".row>*{float:left}",
		tags: []string{".row"},
		want: ".row>*{float:left}",
	},
	{
		// "." specifier after first char
		css:  ".row.uniform>*>:first-child{margin-top:0}",
		tags: []string{".row"},
		want: ".row.uniform>*>:first-child{margin-top:0}",
	},
	{
		// [] specifier after pos 0
		css:  `input[type="submit"]{-moz-appearance:none}`,
		tags: []string{"input"},
		want: `input[type="submit"]{-moz-appearance:none}`,
	},
	{
		// make sure rules work inside @
		css:  "@media(min-width:1200px){.container{max-width:1140px}}",
		tags: []string{".container"},
		want: "@media(min-width:1200px){.container{max-width:1140px}}",
	},
	{
		// if @ query contains nothing, then @ should be removed
		css:  "@media(min-width:1200px){.container{max-width:1140px}}",
		tags: []string{"h1"},
		want: "",
	},
	{
		// test rendering of custom rules, e.g. --blue
		css:  ":root{--blue:#007bff;}",
		tags: []string{"h1"},
		want: ":root{--blue:#007bff;}",
	},
	{
		// rendering of @ rules
		css:  "@import url(font-awesome.min.css);",
		tags: []string{"h1"},
		want: "@import url(font-awesome.min.css);",
	},
}

func TestCut(t *testing.T) {
	for i, tcase := range testcases {
		cf := NewCSSFormat(0, false, tcase.tags)
		in := strings.NewReader(tcase.css)
		out := bytes.Buffer{}
		err := cf.Format(in, &out)
		if err != nil {
			t.Errorf("case %d failed with error: %s", i, err)
		}
		outString := out.String()
		if tcase.want != outString {
			t.Errorf("case %d failed: want %q got %q", i, tcase.want, outString)
		}
	}
}
