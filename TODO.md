
# Action Items

## Bugs

## Next

+ [ ] Rewrite wmeave in Go, explore supporting Go implementation of HandlebarsJS, Pango and text/template
    + [x] cmd/mweave/mweave skeleton
    + [x] Makefile and supporting Bash scripts
    + [x] Parse()
    + [x] Render AST to JSON
    + [x] Render AST to XML
    + [x] Weave()
    + [ ] Tangle()
+ [ ] Bump version number to v0.1.x for Go implementation, leave NodeJS implementation at v0.0.x
+ [ ] Figure out if I can depreciate the NodeJS implementation pointed to by npmjs.org.


## Someday, Maybe

+ [ ] Rewrite mweave as itself, i.e. mweave -tangle -i mwmeave.mweave would generate mweave.go, mweave_test.go, cmd/mweave/mweave.go
+ [ ] Add [shorthand](https://github.com/rsdoiel/shorthand) as macro system
+ [ ] Add support to render mweave render documentation and source as Jupyter Notebook
+ [ ] Add support to render PDF documentation
+ [ ] Add support to render ePub documentation
+ [ ] Add a native Markdown processor to mweave (so outputs could be source and HTML rather than source and Markdown)
+ [ ] Rename project to somthing more accurately descriptive (e.g. _mweaver_, _mlit_, _mloom_)


## Completed

+ [x] Fix package.json to use newer dependency due to security issues in stale versions of marked and handlebars.js
