package csstool

import (
	"strings"
	"testing"

	"github.com/tdewolff/parse/css"
)

func TestGetSelectors(t *testing.T) {

	cases := []struct {
		src  string
		want string
	}{
		{
			"p",
			"p",
		},
		{
			"a:hover",
			"a",
		},
		{
			"p.foo",
			"p.foo",
		},
		{
			".foo",
			".foo",
		},
		{
			// multiple class selector
			// only check first class
			"p.foo.bar",
			"p.foo",
		},
		{
			".foo.bar",
			".foo",
		},
		{
			"[ hidden ]",
			"[hidden]",
		},
		{ // special case
			"*",
			"*",
		},
		{
			":root",
			":root",
		},
		{
			"::after",
			"::after",
		},
		{
			"#foo",
			"#foo",
		},
		{
			"p#foo",
			"p#foo",
		},
	}

	for i, c := range cases {
		p := css.NewParser(strings.NewReader(c.src+"{}"), false)
		p.Next()
		got := string(primarySelector(p.Values()))
		if got != c.want {
			t.Errorf("case %d: want %q got %q", i, c.want, got)
		}
	}
}
