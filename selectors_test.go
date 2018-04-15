package csstool

import (
	"strings"
	"testing"

	"github.com/tdewolff/parse/css"
)

func TestSelectors(t *testing.T) {
	cases := []struct {
		src  string
		want []string
	}{
		{
			src:  "p",
			want: []string{"p"},
		},
		{
			src:  "a:hover",
			want: []string{"a"},
		},
		{
			src:  "p.foo",
			want: []string{"p.foo"},
		},
		{
			src:  "[ hidden ]",
			want: []string{"[hidden]"},
		},
		{ // special case
			src:  "*",
			want: []string{"*"},
		},
		{
			src:  ":root",
			want: []string{":root"},
		},
		{
			src:  "::after",
			want: []string{"::after"},
		},
		{
			src:  "#foo",
			want: []string{"#foo"},
		},
		{
			src:  "p#foo",
			want: []string{"p#foo"},
		},
		{
			src:  "pre > code",
			want: []string{"pre", "code"},
		},
		{
			src:  "pre code",
			want: []string{"pre", "code"},
		},

		{
			src:  "[type=reset]::-moz-focus-inner",
			want: []string{"[type=reset]"},
		},
		{
			src:  ".no-gutters>.col",
			want: []string{".no-gutters", ".col"},
		},
		{
			src:  ".table thead th",
			want: []string{".table", "thead", "th"},
		},
		{
			src:  ".table tbody+tbody",
			want: []string{".table", "tbody", "tbody"},
		},
		{
			src:  ".table-striped tbody tr:nth-of-type(odd)",
			want: []string{".table-striped", "tbody", "tr"},
		},
		{
			src:  ".table-hover .table-primary:hover>td",
			want: []string{".table-hover", ".table-primary", "td"},
		},
		{
			src:  "a:not([href]):not([tabindex]):focus",
			want: []string{"a"},
		},
		{
			src:  ".table-dark.table-bordered",
			want: []string{".table-dark", ".table-bordered"},
		},
		{
			src:  ".form-control::-ms-input-placeholder",
			want: []string{".form-control"},
		},
		{
			src:  "select.form-control:focus::-ms-value",
			want: []string{"select.form-control"},
		},
		{
			src:  ".btn.focus",
			want: []string{".btn", ".focus"},
		},
		{
			src:  ".btn:not(:disabled):not(.disabled).active",
			want: []string{".btn"},
		},
		{
			src:  ".input-group-lg>.input-group-append>select.btn:not([size]):not([multiple])",
			want: []string{".input-group-lg", ".input-group-append", "select.btn"},
		},
		{
			src:  ".custom-file-input:lang(en)~.custom-file-label::after",
			want: []string{".custom-file-input", ".custom-file-label"},
		},
		{
			src:  ".custom-select.is-valid~.valid-feedback",
			want: []string{".custom-select", ".is-valid", ".valid-feedback"},
		},
		{
			src:  ".btn-group>.btn-group:not(:first-child)>.btn",
			want: []string{".btn-group", ".btn-group", ".btn"},
		},
		{
			src:  ".custom-range::-webkit-slider-thumb:active",
			want: []string{".custom-range"},
		},
		{
			src:  ".navbar-light .navbar-nav .active>.nav-link",
			want: []string{".navbar-light", ".navbar-nav", ".active", ".nav-link"},
		},
		{
			src:  ".row.uniform>*>:first-child",
			want: []string{".row", ".uniform", "*", ":first-child"},
		},
		//
		// HELP
		//
		{
			src:  ".dropdown-menu[x-placement^=bottom]",
			want: []string{".dropdown-menu"},
		},
		{
			src:  "abbr[title]",
			want: []string{"abbr"},
			//want: []string{"abbr[title]"},
		},
		{
			src:  "a.btn.disabled",
			want: []string{"a.btn", ".disabled"},
		},
	}
	for i, c := range cases {
		p := css.NewParser(strings.NewReader(c.src+"{}"), false)
		p.Next()
		got := selectors(p.Values())
		want := c.want
		if len(got) != len(want) {
			t.Errorf("case %d: want %v, got %v", i, want, got)
			continue
		}
		for j := 0; j < len(got); j++ {
			if got[j] != want[j] {
				t.Errorf("case %d: want %q, got %q", i, want[j], got[j])
			}
		}
	}
}
