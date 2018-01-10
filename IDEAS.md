
Converting from NodeJS to Goland:

+ extension would be .mw or .mweave
+ initial go port to target generating Markdown and source code files from mweave document
    + someday maybe render ePub and PDF?
    + someday maybe render .TeX and .LaTeX?
+ option for code block should support filename and fragment number
+ option to specify "paging" to different files (e.g. for generating slides)
+ option to generate "front matter" for markdown and source code
+ option to handle "front matter" like fountain
+ option to render AST of mweave for external precessing
+ explicit support for weave,tangle and macros via Shorthand


Some feature ideas:

+ Use the pound sign in a URL to indicate block order. This would allow code to be written out of sequence as well as allow literate text to wrap blocks of code going to the same file.
+ Embed a comment in the render code file(s) that indicate where they came from in the Markdown document
    - e.g. /\*mweave 2,14 myMardownFile.md \*/

+ Add source map support
+ Add support for name of file to be rendered to end block quote rather than be immediately before it.
+ Figure out a way to have multiple blocks accumulate but output only the last link to the final version

# Reference

+ [Source Maps background](http://www.html5rocks.com/en/tutorials/developertools/sourcemaps/) - a littled dated, but background never the less
+ [Github repo](https://github.com/ryanseddon/source-map/wiki/Source-maps%3A-languages,-tools-and-other-info) about source maps
+ [Mozilla Source Map project](https://github.com/mozilla/source-map/)
+ [UglifyJS2 and source maps](https://github.com/mishoo/UglifyJS2)

