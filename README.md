# csstool
CSS filters and formatters in golang

## csscut 

Use awesome CSS frameworks without the weight by cutting out unused rules.

`csscut` is somewhat like [purgecss](https://www.purgecss.com) ([GitHub](https://github.com/FullHuman/purgecss)). It scans your HTML for elements, classes and identifiers and then cuts out any CSS rule that doesn't apply. The results for a small site using a framework like [bootstrap](https://getbootstrap.com) can be dramatic:

|                | Bootstrap | csscut   | savings |
|----------------|-----------|----------|---------|
| uncompressed   |   141k    |   8.1k   |   94%   |
| compressed     |    20k    |   2.6k   |   87%   |


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

## cssformat 

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

