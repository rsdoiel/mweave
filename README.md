
# mweave

_mweave_ is a literate programming experiment I started by in 2012. I liked
Donald Knuth's ideas about [literate programming](https://en.wikipedia.org/wiki/Literate_programming)
but didn't wish to write in TeX.  Today I mostly write documents 
in an extended version of Markdown or in the case of screenplay's Fountain.
While TeX (and particularly LaTeX) remain really cool in my thinking I just find
it easier to type and proof read in a simpler form of markup.

## History

### Where mweave started 2012

This project started out as an experiment to write a document generator written
for NodeJS's in JavaScript. While I thought of it as "literate programming" what
I implemented was really just an inside out document generator.  Code blocks that were
preceeded by a link were scraped and written to a file indicated by the targetted link.
I did not use the concepts of "tangle" and "weave" individually and I didn't
support the arbitrary ordering of code blocks when generating files.

My bootstrap was written in a Bash script that would process the README.md file
using _vi_, _sed_, to generate a *mw-bootstrap.js* and then _npm_ to build it. 
In the end my initial experiment failed because I failed to use it. 

I stopped the experiment sometime in 2013. It was not a success as I did not find
it compelling enough to use in practice. It was a fun thing to write and think about.
Eventually I moved on from NodeJS/JavaScript to other development languages and projects.
Of course cool things like [Jupyter Notebooks](https://jupyter.org/) happened as well.

### mweave in January, 2018

Today I find myself working in a Research Library and literate programming is again
calling to me.  This experiment builds on the 2012 ideas but now is implemented in Golang.
We'll see if this moves beyond a toy program in the coming years. RSD, 2018-01-05

## The experiment

_mweave_ is a Golang package and command line program. It provides both "tangle" and
"weave" functions that are expected in literate programming tools today. An _mweave_
file is presumed to be a UTF-8 plain text file with an extension of ".mweave" or ".mw".
Documentation is rendered to Markdown file(s) and source to the specified source code files.
It is still one command but by default both documentation and source are generated.
At it's core _mweave_ is a pre-processor that looks for _mweave_ directives. Unlike the 2012
version _mweave_ directives are embedded in HTML/XML style comments.

```
    <!--mweave:begin "HelloWorld.md" 0 -->
    # Hello World!

    This is an example of an embedded document to be extracted by 
    [mweave](https://github.com/rsdoiel/mweave).
    <!--mweave:end -->
```

_mweave_ reads in the entire source document splitting it up based on the directives it encounters.
There are two directives - *mweave:begin* and *mweave:end*.  *mweave:begin* takes two required 
parameters, the filename (string) and an order value (an integer). The ordering value is used by
tangle to order blocks of texts associated with the filename. *mweave:end* takes no parameter.
Currently I have no other directives planned and do not plan on implemented any sort of macro 
system. The _mweave_ directives are expected to be the first non-space elemnt on the line. 
They take up a whole line. 

This experiment is still rediculiously simple like 2012. Something that can be implemented in an 
morning or two hacking. I am trying to sort out how simple a tool I can write and still support 
literate programming. My hypothis is that if the tool is simple enough I might actually use it and 
find it more useful and interesting to maintain.

## Requirements

+ Golang version 1.9.2 or better

_mweave_ is "go get"-able.

```shell
    go get -u github.com/rsodiel/mweave/...
```


