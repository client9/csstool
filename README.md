# csstool
CSS filters and formatters in golang

[![Build Status](https://travis-ci.org/client9/csstool.svg?branch=master)](https://travis-ci.org/client9/csstool)

## csscut 

Use awesome CSS frameworks without the weight by cutting out unused rules.

`csscut` is somewhat like [purgecss](https://www.purgecss.com) ([GitHub](https://github.com/FullHuman/purgecss)). It scans your HTML for elements, classes and identifiers and then cuts out any CSS rule that doesn't apply. The results for a small site using a framework like [bootstrap](https://getbootstrap.com) can be dramatic:

|                | Bootstrap | csscut   | savings |
|----------------|-----------|----------|---------|
| uncompressed   |   141k    |   8.1k   |   94%   |
| compressed     |    20k    |   2.6k   |   87%   |


See also [Hugo #4446](https://github.com/gohugoio/hugo/issues/4446#issuecomment-370070252)

### Example

For use with [hugo](https://gohugo.io) using [bootstrap](https://getbootstrap.com):

```bash
# install the binary
go get github.com/client9/csstool/cmd/csscut

# build your site, by default the output is in `public`
hugo

# create new minimized CSS file from bootstrap
curl -s https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/css/bootstrap.min.css | \
    csscut -html 'public/**/*.html' > static/bootstrap-csscut.min.css
```

Of course, you'll need to use the new `bootstrap-csscut.min.css` file in your template source.

Be sure to put the HTML file pattern `'public/**/*.html'` in single quotes.

### The Correct Algorithm

The "correct way" is to 

1. Read in all the CSS files, and extract all the selectors
2. For each HTML file, execute each selector and see if it returns anything
3. Use that data to the emit each CSS file with only the selectors that were used.

There are a few problems:

1. Slow, as you are executing _n_ CSS rules against _m_ HTML files.
2. Need a perfect CSS Level 3 (or 4!) selector library, else you might strip out rules that are actually used.  
3. Need to know which pseudo-classes matter and which one's don't.  For instance `:hover` can be ignored, but `:last_child` needs to be evaluated.

### CSSCut Algorithm

1. Read each HTML file and keep track of elements, classes and ids found.  This is can be in a fast and simple way.
2. Scan the CSS file and using the _primary selector_ and previous step, keep or remove the selector and it's rule.   Combinators and pseudo-elements are ignored in making the decision. This can be in streaming mode.

Ok what does that mean?  If the `li` element is found in the html then the CSS selector is also allowed `li:first-of-type` and `li li` even if you don't have a nested list.  If `img` is found, then the CSS rule `img ~ p` is also allowed, even if you don't have paragraph as a sibling of an image.

As a special case, "universal selectors" are passed through: `*, ::before, ::after`. Pure attribute selectors are also passed through: `[hidden]`.

In practice this works well for "flat" CSS frameworks such as bootstrap.  For highly specified CSS it might not work as well. 

# cssformat 

Makes minified CSS readable.

```
go get github.com/client9/csstool/cmd/cssformat

cssformat < bootstrap.min.css
```

## csscount

See commonly or rarely used CSS classes and identifiers.

Work in progress

## Credits

* The CSS parsing is done by the most excellent [tdewolff/parse](https://github.com/tdewolff/parse) which powers [tdewolff/minify](https://github.com/tdewolff/minify).
* The [double star](https://www.client9.com/golang-globs-and-the--double-star-glob-operator/) globbing / pattern matching is handled by [mattn/go-zglob](https://github.com/mattn/go-zglob)

