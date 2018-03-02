# cssformat

Lightly reformats minimized CSS code to be readable

## Usage

Bare bones CSS reformater. For now it only support `stdin` and `stdout`.   You can change the indent to whatever number of spaces you want but beyond that it's unlikely to be your CSS style guide enforcement tool.

```css
css-format < bootstrap.min.css 
```

It does have a feature to strip away or preserve specified tags or classes.   This is really more for testing `css-purge` but you might find a use case.

- [x] Output no spaces and newlines
- [ ] Tab output instead of spaces
- [x] [globstar](https://www.client9.com/golang-globs-and-the--double-star-glob-operator/) support
- [ ] It would be nice to specify sorting of rules (i.e. `background-color` comes before `color` etc) so a given 
- [ ] In place reformat
- [ ] debug logging to see what is dropped or not

