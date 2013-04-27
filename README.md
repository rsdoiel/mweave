Markdown Weave
==============

An experiment in using Markdown and Donald Knuth's literate programming concept.

# What is mw.js?

This is an experiment in exporing an of the ideas from  Donald Knuth's 
literate programming concept. This README.md is intended to work as a
bootscript [mw-boostrap.js](mw-boostrap.js).

[mw-bootstrap.js](mw-bootstrap.js)
```JavaScript
    /**
     * mw-bootstrap.js - an experiment in literate style programming in a 
     * markdown file.
     * @author R. S. Doiel, <rsdoiel@gmail.com>
     * copyright (c) 2013 all rights reserved
     * Licensed under the BSD 2-clause license. See http://opensource.org/licenses/BSD-2-Clause
     */
     
     var fs = require("fs");

     lines = fs.readFileSync("README.md").toString().split("\n"),
        outputs = {};

     lines.forEach(function (line, i) {
        var check, target;
        check = trim(line);
        if (i < lines.length - 2 &&
            lines[i + 1].indexOf(```) === 0 &&
            check[0] === '[' && check[check.length - 1] === ')') {
            filename = line.substr(chech.lastIndexOf('(') + 1, -1);
            console.log("Writing " + filename);
        }
     });

```

Above is the bootstrap code.  To "bootstrap" I'm using vi's write lines command to generate the
the first instance of **mw-bootstrap.js**. Then **mw-bootstrap.js** will be used to process
**README.md** and generate subsequent versions. To discover the filename to write to I'm 
looking at the line immediately before the tripple quotes and if there is one item in square
brackets (e.g. [mw-bootstrap.js]) then I assume that is the filename. If there is a blank line
before the tripple quotes then I don't write that quoted block out.


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
      "devDependencies": {
          "yuitest"
      },
      "dependencies": {
          "yui": "3.10.x"
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
```
package.json


