# csstool
CSS filters and formatters in golang

For use in [hugo](https://gohugo.io):

```bash
go get github.com/client9/csstool/cmd/cssdrop
hugo 
cssdrop -html 'public/**/*.html' < static/boostrap.min.css > static/bootstrap-dropped.min.css
```

Be sure to put html file pattern `'public/**/*.html'` in single quotes.

