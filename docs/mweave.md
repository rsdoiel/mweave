
# USAGE

	mweave [OPTIONS] MWEAVE_FILENAME

## SYNOPSIS


mweave is and experimental literate programming tool. It is 
inspired by Knuth's but much simplier and with the primary
purpose of creating a platform for writing interactive 
function (e.g. Adventure like text games).


## OPTIONS

```
    -examples                 display examples
    -generate-markdown-docs   generate Markdown documentation
    -h, -help                 display help
    -i, -input                set input filename (the mweave file)
    -json                     write mweave doc as JSON
    -l, -license              display license
    -macros                   preprocess shorthand macros
    -nl, -newline             add a trailing newline
    -o, -output               set output filename
    -quiet                    suppress error messages
    -t, -tangle               generate source code files (e.g. program code)
    -v, -version              display version
    -w, -weave                generate documentations files (e.g. Markdown output)
    -xml                      write mweave doc as XML
```


## EXAMPLES


generate source files from an mweave document

    mweave -i document.mweave -tangle

generate documentation from an mweave document

    mweave -i documemt.mweave -weave

display mweave parse results as XML

    mweave -i document.mweave -xml

display mweave parse results as JSON

    mweave -i document.mweave -xml


mweave v0.1.1
