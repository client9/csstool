# csstool
CSS filters and formatters in golang

[![Build Status](https://travis-ci.org/client9/csstool.svg?branch=master)](https://travis-ci.org/client9/csstool)

## css cut 

Use awesome CSS frameworks without the weight by cutting out unused rules.

`css cut` is similar to [purgecss](https://www.purgecss.com) ([GitHub](https://github.com/FullHuman/purgecss)). It scans your HTML for elements, classes and identifiers and then cuts out any CSS rule that doesn't apply. The results for a [small site](https://www.client9.com/) using a framework like [bootstrap](https://getbootstrap.com) can be dramatic:

|                | Bootstrap | css cut  | savings |
|----------------|-----------|----------|---------|
| uncompressed   |   141k    |   5.6k   |   96%   |
| compressed     |    20k    |   1.8k   |   91%   |


See also [Hugo #4446](https://github.com/gohugoio/hugo/issues/4446#issuecomment-370070252)

### Example

For use with [hugo](https://gohugo.io) using [bootstrap](https://getbootstrap.com):

```bash
# install the binary
go get github.com/client9/csstool/css

# build your site, by default the output is in `public`
hugo

# create new minimized CSS file from bootstrap
curl -s https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/css/bootstrap.min.css | \
    css cut --html 'public/**/*.html' > static/bootstrap-csscut.min.css
```

Of course, you'll need to use the new `bootstrap-csscut.min.css` file in your template source.

Be sure to put the HTML file pattern `'public/**/*.html'` in single quotes.

### Usage

TK - likely to change, feedback welcome

### API

TK - likely to change, feedback welcome

### How It Works

#### The Correct Algorithm

The "correct way" to strip out CSS rules might be:

1. Read in all the CSS files, and extract all the selectors
2. For each HTML file, execute each selector and see if it returns anything
3. Use that data to the emit each CSS file with only the selectors that were used.

There are a few problems:

1. Slow, as you are executing _n_ CSS rules against _m_ HTML files.
2. Need a perfect CSS Level 3 (or 4!) selector library, else you might strip out rules that are actually used.  
3. Need to know which pseudo-classes matter and which ones do not.  For instance `:hover` can be ignored, but `:last_child` needs to be evaluated.

#### The CSSCut Algorithm

Since the Correct Way seems problematic, csscut does the following:

1. Read each HTML file and keep track of elements, classes and ids found.
2. Scan the CSS file and convert a selector into a set of "basic selectors".  If a rule is `h1 h2 h3` then the list of basic selectors is h1, h2, and h3. Classes and identifiers (ids) are preserved, while pseudo elements and attribute selectors are ignored.
3. Then if each of the basic selectors is found in out list in the first step, the original selector is preserved.  This not the rule is tossed out.

As a special case, "universal selectors" are passed through: `*, ::before, ::after, ::root`. Pure attribute selectors are also passed through: `[hidden]`.

In practice this works well for "flat" CSS frameworks such as [bootstrap](https://getbootstrap.com).  For highly specified CSS it might not work as well. 

# css format 

Makes minified CSS readable.

```
css format < bootstrap.min.css
```

# css minify

```
css minify < bootstrap.min.css
```

minify is a shortcut of 'css format' with all options selected to produce the smallest output. It is "conservative" in that it only removes whitespace and does not do any property value rewriting.

# css count

See commonly or rarely used CSS classes and identifiers.

Work in progress

## Credits

* The CSS parsing is done by the most excellent [tdewolff/parse](https://github.com/tdewolff/parse) which powers [tdewolff/minify](https://github.com/tdewolff/minify).
* The [double star](https://www.client9.com/golang-globs-and-the--double-star-glob-operator/) globbing / pattern matching is handled by [mattn/go-zglob](https://github.com/mattn/go-zglob)

