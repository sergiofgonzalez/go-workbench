# githubhtml
> practising HTML templates by rendering a page with github issues

To run it do:

```bash
go run . > issues.html
google-chrome issues.html
```

To check how escaping of characters is performed you can do:

```bash
go run . repo:golang/go 3133 10535 > issues2.html
google-chrome issues2.html
```