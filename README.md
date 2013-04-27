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
     
     var fs = require("fs"),
        lines = fs.readFileSync("README.md").toString().split("\n"),
        line = "",
        check = "",
        outputs = {},
        i = 0,
        filename,
        start = 0,
        end = 0;

     for (i = 0; i < lines.length; i += 1) {
        line = lines[i];
        check = line.trim();
        if (i < lines.length - 2 &&
            lines[i + 1].indexOf("```") === 0 &&
            check[0] === '[' && check[check.length - 1] === ')') {
            i += 1;
            start = check.lastIndexOf('(') + 1;
            end = check.lastIndexOf(')');
            filename = line.substr(start, end - start);
            console.log("Filename: " + filename);
            outputs[filename] = {start: i + 1, end: -1};
        } 
        if (typeof outputs[filename] !== "undefined" &&
            outputs[filename].end < 0 &&
            line.indexOf("```") === 0) {
            outputs[filename].end = i - 1;
        }
     };
     Object.keys(outputs).forEach(function (ky) {
        console.log("vi -e -c '" + outputs[ky].start + "," + outputs[ky].end + " wq! " + ky);
     });
```

Above is the bootstrap code.  To "bootstrap" I'm using _ex_'s write lines command to generate the
the first instance of **mw-bootstrap.js**. Then **mw-bootstrap.js** will be used to process
**README.md** and generate subsequent versions. To discover the filename to write to I'm 
looking at the line immediately before the tripple quotes and if there is a line then I assume
the link target is the desired filename.  If there is a blank line before the tripple quotes then
I don't write that quoted block out.

Here's the _vi_ command to generate **mw-bootstrap.js** the first time.

```Shell
    vi -e -c "14,53wq! mw-bootstrap.js" README.md;node mw-bootstrap.js
```


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


