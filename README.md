Markdown Weave
==============

An experiment in using Markdown and Donald Knuth's literate programming concept.

# What is mw.js?

This is an experiment in exporing an of the ideas from  Donald Knuth's 
literate programming concept. This README.md is intended to work as a
bootscript [mw-boostrap.js](mw-boostrap.js).

```JavaScript
    /**
     * mw-bootstrap.js - an experiment in literate style programming in a 
     * markdown file.
     * @author R. S. Doiel, <rsdoiel@gmail.com>
     * copyright (c) 2013 all rights reserved
     * Licensed under the BSD 2-clause license. See http://opensource.org/licenses/BSD-2-Clause
     */
     
     var fs = require("fs");

     lines = fs.readFileSync("README.md").toString().split("\n");

     lines.forEach(function (line) {
     });

```mw-bootstrap.js

Above is the bootstrap code.  To "bootstrap" I'm using vi's write lines command to generate the
first pass at mw.js. Then **mw-bootstrap.js** will be used to process **mw.md** and generate 
**mw.js** and **mw_test.js**.

To generate **mw-bootstrap.js** I'm typing README.md in vi. I'm then using the write command
to write the lines in the block describing **mw-bootstrap.js**

## mw.js

mw.js is intended to be the actual Markdown Weave experiment. You should be able to run
commands like

```shell
    node mw.js ExampleMW-1.md
```

and render all the code described in ExampleMW-1.md independant of the text.  This is accomplished
through overloading the "```" (i.e. open tripple quote) operator in Markdown. In github flavored
markdown the text following an opening tripple quote descripts the specific language flavor be
quoted (e.g. javascript, shell, python). Markdown Weave is using the closing tripple quote
for a filename to output to. If closing tripple quotes have no filename nothing is written
to disc. Each additional time a previously referenced filename is encountered the contents
are appended and all files collected are written to disc when the Markdown text has completed
evaluation.  This allows you to have blocks of markdown text discussing sections of the source
code while still generating the whole source code file in the end. Eventually a source map will
be generated so you know where in your markdown source to fix when you encounter an error.




## testing

I am using YUI3's test module. Like _mw.js_ I'm 

## Node package.json

```JavaScript
    {
      "name": "markdown-weave",
      "version": "0.0.0",
      "description": "This is an experiment in using Markdown and some concepts from Donald Knuth's literate programming.",
      "main": "mw.js",
      "scripts": {
        "test": "mw_test.js"
      },
      "repository": {
        "type": "git",
        "url": "git@github.com:rsdoiel/markdown-weave.git"
      },
      "keywords": [
        "markdown",
        "weave"
      ],
      "author": "R. S. Doiel",
      "license": "BSD",
      "readmeFilename": "README.md"
    }
```package.json

