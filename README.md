
# mweave

_mweave_ is a *literate programming* experiment I started by in 2012. I liked
Donald Knuth's ideas about [literate programming](https://en.wikipedia.org/wiki/Literate_programming)
but didn't enjoy the implementations I had available. I was also looking at it from
a programming project perspective.  In this incarnation of my experiment I am looking 
at *literate programming* from the perspective of a _writing tool_ for creative
projects such as [interactive fiction](https://en.wikipedia.org/wiki/Interactive_fiction) 
and [electronic literature](https://en.wikipedia.org/wiki/Electronic_literature) more
generally.

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
at npmjs.org shows 0.0.2. I've sense flagged it as deprecated.

While it was a fun thing to write and think about I moved on to other projects
Of course cool things like [Jupyter Notebooks](https://jupyter.org/) 
happened while this experiment languished. _wmeave_ is inspired in part by watching
scientists use Jupyter Notebooks in Open Science.


### mweave in January, 2018

Today I find myself working in a Research Library and literate programming is again
calling to me.  This experiment builds on the 2012 ideas but now is implemented in Golang.
We'll see if this moves beyond a toy program in the coming years. RSD, 2018-01-05

## The experiment

_mweave_ is a Golang package and command line program. It provides both "tangle" and
"weave" functions that are expected in literate programming tools today. An _mweave_
file is presumed to be a UTF-8 plain text file with an extension of ".mweave" or ".mw".
Documentation is rendered to Markdown file(s) and source to the specified source code files.
It is still one command but by picking the `-weave` or `-tangle` options you can generate
both documentation (e.g. Markdown output) and source code file(s).
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

_mweave_ reads in the entire source document splitting it up based on the directives it encounters.
There are three directives - *mweave:source*, *mweave:macro* and *mweave:end*.  *mweave:source* 
takes two required parameters, the filename (string) and an order value (an integer). The ordering 
value is used by tangle to order blocks of texts associated with the filename. *mweave:macro* 
takes two parameters one is a label string and the other is a macro (shorthand) op. The last is 
*mweave:end* which takes no parameter.  Currently I have no other directives planned. I am 
considering integrating [shorthand](https://github.com/rsdoiel/shorthand) as the macro system 
sometime in the future.  The _mweave_ directives are expected to be the first non-space element 
on the line.  They take up a whole line. 

This 2nd _mweave_ experiment is still rediculiously simple like 2012. v0.1.0 was I implemented 
in less than a day so I could experiment again with literate programming. v0.1.1, integrating
a macro system was another day (though the macro engine already existed). I am trying to sort out how 
simple a tool I can write and still support literate programming. My hypothis is that if the 
tool is simple enough I might actually use it and find it more useful and interesting to maintain.

### The macro system in v0.1.1

_mweave_ (v0.1.1) now has en experimental macro system. The macro system is 
[shorthand](https://rsdoiel.github.io/shorthand). Shorthand provides additional
functionality like including other files, running shell processes, running Markdown
processors, evaluting other other shorthand files. The experiment will not see
if the two mweave directives plus shorthand gives us a useful platform for doing literate
development.


## Requirements

+ Golang version 1.9.2 or better

## Installation

_mweave_ is "go get"-able.

```shell
    go get -u github.com/rsodiel/mweave/...
```


