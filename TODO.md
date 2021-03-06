
# Action Items

## Bugs

+ [ ] I can build up blocks with Macros but mweave:source should assemble the source AND expand macros...

## Next

+ [ ] Decide if `<!--meave:shorthand` evaluation is trigger only in tangle or weave, or both
+ [ ] Build out test coverage
+ [ ] Integrate [shorthand](https://github.com/rsdoiel/shorthand) as mweave's macro system
    + [x] shorthand should be a cli option and execute before mweave.Parse() if selected
    + [ ] Add shorthand related options to mweave (e.g. so we can pre-load macro libraries)
+ [ ] Bump version number to v0.1.1 for Go implementation


## Someday, Maybe

+ [ ] Rewrite mweave in mweave itself, i.e. mweave -tangle -i mwmeave.mweave would generate mweave.go, mweave_test.go, cmd/mweave/mweave.go
+ [ ] Exploring implementing a native Markdown processor in mweave 
    + [ ] Add support to render mweave render documentation and source as Jupyter Notebook
    + [ ] Add support to render PDF documentation
    + [ ] Add support to render ePub documentation
+ [ ] Explore supporting Go implementation of HandlebarsJS, Pango and text/template directly in _mweave_
+ [ ] Rename project to somthing more accurately descriptive (e.g. _mweaver_, _mlit_, _mloom_)


## Completed

+ [x] Rewrite wmeave in Go
    + [x] cmd/mweave/mweave skeleton
    + [x] Makefile and supporting Bash scripts
    + [x] Parse()
    + [x] Render AST to JSON
    + [x] Render AST to XML
    + [x] Weave()
    + [x] Tangle()
+ [x] Fix package.json to use newer dependency due to security issues in stale versions of marked and handlebars.js
+ [x] Figure out if I can depreciate the NodeJS implementation pointed to by npmjs.org.
