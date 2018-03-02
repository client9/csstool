# css-count

Return usage frequency of HTML tags and CSS classes

## Usage

Currently just `stdin` and `stdout`

```
curl https://your-hot-site/hot-page.html | css-count

.col-lg-8,.container,.justify-content-between,.justify-content-center,.language-bash,.mb-3,.navbar,.navbar-light,.pb-3,.pl-0,.row,.small,.text-center,article,body,code,div,h1,head,hr,html,li,link,meta,nav,ol,p,pre,script,style,time,title
```

Ok Lots of problems here:

- [x] need globs, in fact [superglobs](https://www.client9.com/golang-globs-and-the--double-star-glob-operator/)
- [ ] need output as list
- [ ] need output as list with count (see what tags are rarely used)
- [ ] move actual computation code into library instead of CLI
- [x] sort output for consistancy and testing
- [ ] need option to extract just elements, classes or ids


