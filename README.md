
# mweave

_mweave_ is a *literate programming* experiment I started by in 2012. I liked
Donald Knuth's ideas about [literate programming](https://en.wikipedia.org/wiki/Literate_programming)
but didn't enjoy the implementations. This incarnation of my experiment (2018)
I am looking at *literate programming* from the perspective of a _writing tool_ 
for creative projects in [interactive fiction](https://en.wikipedia.org/wiki/Interactive_fiction) 
and [electronic literature](https://en.wikipedia.org/wiki/Electronic_literature). 

## mweave experiment history

### mweave started 2012

This project started out as an experiment to write a document generator written
for NodeJS's in JavaScript. While I thought of it as "literate programming" what
I implemented was really just an inside out document generator.  Code blocks that were
preceeded by a link were scraped and written to a file indicated by the targetted link.
I did not use the concepts of "tangle" and "weave" individually and I didn't
support the arbitrary ordering of code blocks or macros.

My bootstrap was written in Bash, it processed the README.md file using _vi_, _sed_, 
to generate a *mw-bootstrap.js* and then that processed a file, _mw.md_ to implement 
_mw.js_ and _npm_ to assemble dependencies.  In the end my initial experiment failed because 
I failed to use _mw.js_ on a regular basis.  It wasn't compelling.  The version number 
at npmjs.org shows 0.0.2. I've since flagged it as deprecated. I am no longer developing
a NodeJS based implementation.

If you're looking for something practicle two interesting projects capture what I was
thinking about. They are [Jupyter Notebooks](https://jupyter.org/)  and [R Notebooks](https://github.com/rstudio/rmarkdown).
Both have grown out of the [Open Science](https://en.wikipedia.org/wiki/Open_science) and 
[Open Data](https://en.wikipedia.org/wiki/Open_data) communities. Very cool stuff.


### mweave in January, 2018

Today I find myself working in a Research Library and literate programming is again
calling to me.  This experiment builds on the 2012 ideas but now is implemented in Golang.
We'll see if this moves beyond a toy program in the coming years. RSD, 2018-01-05

## The experiment

    Can _mweave_ be a useful tool for writing interactive fiction?  

_mweave_ is a Golang package and command line program. It provides both "tangle" and
"weave" functions.  The _mweave_ command line program integrates macro support by
pre-processing the text through [shorthand](https://rsdoiel.github.io/shorthand) a
very simple label expander. An _mweave_ file is a UTF-8 plain text file with an 
extension of ".mweave" or ".mw".  Documentation is rendered to Markdown file(s) 
and source to the specified source code files. Like the original project _mweave_
is still one command but the explicit options of `-weave` or `-tangle` are now
included so you can generate both markdown and source code files.

At it's core _mweave_ is a pre-processor that looks for _mweave_ directives. Unlike the 2012
version _mweave_ directives are embedded as an HTML/XML style comments.

### Hello World

As an example you can render a [helloworld](demos/helloworld.py) python script from [helloworld.mweave](demos/helloworld.mweave) using `-tangle` and render a [helloworld](demos/helloworld.md) Markdown page by using
`-weave`. Processing that markdown using a Markdown process like [mkpage](https://caltechlibrary.github.io/mkpage)
would give you the final [helloworld](demos/helloworld.html) page.

Here is the example mweave file--

```
    # Hello World!

    This is an example of an embedded document to be extracted by 
    [mweave](https://github.com/rsdoiel/mweave).
    <!--mweave:source "helloworld.py" 0 -->
    ```python
        #!/usr/bin/env python3
        print("Hello World!")
    ```
    <!--mweave:end -->
```

Here are the commands to render [helloworld.md](demos/helloworld.md) and [helloworld.py](demos/helloworld.py)
from our [helloworld.mweave](demos/helloworld.mweave) source.

```shell
    mweave -weave -i helloworld.mweave -o helloworld.md
    mweave -tangle -i hellowolrd.meave
```

Notice that tangle ignores the output file name. That is beceause the source files are set in the 
mweave begin directive.


### How it works

_mweave_ reads in the entire source document and runs through the _shorthand_ macro expander. It then
splits the document into plain text sections and source sections.  Source sections start with 
*mweave:source* and end with a *mweave:end* HTML style comment.  *mweave:source* takes two required 
parameters, the filename (string) and an order value (an integer). The ordering value is used by 
tangle to order blocks of texts associated with the filename. 

Before parsing into source and plain text blocks the [shorthand](https://rsdoiel.github.io/shorthand)
macro processor runs over the code and the resulting text is then parsed into source and plain text.
See [shorthand](https://rsdoiel.github.io/rsdoiel/shorthand/docs) for details about _shorthand_.

This 2nd _mweave_ experiment is still rediculiously simple like 2012. v0.1.0 was I implemented 
in less than a day so I could experiment again with literate programming using the HTML comment
indicator for flag source files to output. v0.1.1 integrated a macro system was another day 
(though the macro engine already existed). I am trying to sort out how simple a tool I can write 
and still support literate programming. My hypothis is that if the tool is simple enough I might 
actually use it and find it more useful and interesting to maintain.

## Requirements

+ Golang version 1.9.2 or better

## Installation

_mweave_ is "go get"-able.

```shell
    go get -u github.com/rsodiel/mweave/...
```


