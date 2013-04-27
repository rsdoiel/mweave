Markdown Weave
==============

# What is mw.js?

An experiment in using Markdown and Donald Knuth's literate programming concept.


## Bootstraping mw.js

I like the idea of writing _mw.js_ using _mw.js_.  To do that I a very simple implementation
of _mv.js_ which I'm calling _mw-bootstrap.js_.  That's the proof concept below. I'm
leveraging Markdown syntax, JavaScript to generate the _vi_ commands to extract
the code. The core to actually generate _mw.js_ itself will be in a markdown file
name [Markdown-Weave.md](Markdown-Weave.md).

Here's code to bootstrap this whole thing-

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
        lines = [],
        line = "",
        check = "",
        outputs = {},
        i = 0,
        markdownFilename = "README.md",
        filename,
        start = 0,
        end = 0;

     if (process.argv.length === 3) {
        markdownFilename = process.argv[2];
     }
     lines = fs.readFileSync(markdownFilename).toString().split("\n");
     for (i = 0; i < lines.length; i += 1) {
        line = lines[i];
        check = line.trim();
        if (i < lines.length - 2 &&
            lines[i + 1].indexOf("```") === 0 &&
            check[0] === '[' && check[check.length - 1] === ')') {
            i += 2;
            start = check.lastIndexOf('(') + 1;
            end = check.lastIndexOf(')');
            filename = line.substr(start, end - start);
            console.log("# Output Filename: " + filename);
            outputs[filename] = {start: i + 1, end: -1};
        } else if (typeof outputs[filename] !== "undefined" &&
            outputs[filename].end < 0 &&
            line.indexOf("```") === 0) {
            outputs[filename].end = i;
            filename = "";
        }
     };
     Object.keys(outputs).forEach(function (ky) {
        console.log("# This vi command to generate the code for " + ky);
        console.log("vi -e -c '" + outputs[ky].start + "," + outputs[ky].end + " wq! " + ky + "' " +
            markdownFilename);
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
    vi -e -c "21,68wq! mw-bootstrap.js" README.md;node mw-bootstrap.js
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


