Markdown Weave
==============

# What is mw.js?

An experiment in using Markdown and Donald Knuth's literate programming concept.


## Bootstraping mw.js

I like the idea of writing _mw.js_ using _mw.js_.  To do that I wrote a very simple implementation
of _mv.js_ which I'm calling _mw-bootstrap.js_.  I'm leveraging Markdown syntax via JavaScript to 
generate the _vi_ commands to extract the code. The source for _mw.js_ will be generate by running
_mw-bootstrap.js_ on [Markdown-Weave.md](Markdown-Weave.md).

Here is code to bootstrap this whole thing ---

[mw-bootstrap.js](mw-bootstrap.js)
```JavaScript
    #!/usr/bin/env node
    /**
     * mw-bootstrap.js - an experiment in literate style programming in a 
     * markdown file.
     * @author R. S. Doiel, <rsdoiel@gmail.com>
     * copyright (c) 2013 all rights reserved
     * Licensed under the BSD 2-clause license. See http://opensource.org/licenses/BSD-2-Clause
     */
    require("shelljs/global"); 
    var fs = require("fs"),
        lines = [],
        line = "",
        check = "",
        outputs = {},
        i = 0,
        markdownFilename = "Markdown-Weave.md",
        filename,
        start = 0,
        end = 0;
    
    function exportLines(outFilename, start, end, lines) {
        var i = 0;
        console.log("# Output Filename: " + outFilename);
        for (i = start; i < lines.length  && i < end; i += 1) {
            lines[i] = lines[i].replace(/\t/g, "    ").replace(/^    /, "");
        }
        fs.writeFile(outFilename, lines.slice(start, end).join("\n"), function (err) {
            if (err) {
                console.error(err);
                process.exit(1);
            }
        });
    }
    
    if (process.argv.length === 3) {
        markdownFilename = process.argv[2];
    }
    lines = fs.readFileSync(markdownFilename).toString().split(/\n|\r\n/);
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
            console.log("# Found Filename: " + filename);
            outputs[filename] = {start: i, end: -1};
        } else if (typeof outputs[filename] !== "undefined" &&
                outputs[filename].end < 0 &&
                line.indexOf("```") === 0) {
            outputs[filename].end = i;
            filename = "";
        }
    };
    Object.keys(outputs).forEach(function (outFilename) {
        exportLines(outFilename, outputs[outFilename].start, outputs[outFilename].end, lines);
    });
```

You can bootstrap with a few Unix commands (_vi_, _sed_, _chmod_, and _node_).

[bootstrap.sh](bootstrap.sh)
```shell
    #!/bin/bash
    npm install shelljs
    vi -e -c "20,79wq! mw-bootstrap.js" README.md
    sed -e "s/    //" -i mw-bootstrap.js
    chmod 770 mw-bootstrap.js
    ./mw-bootstrap.js
    npm install
    npm test
```


# Further reading

* [Markdown-Weave.md](Markdown-Weave.md) is the source to _mw.js_ and _cli.js_ (i.e. command line tool). Since this is process by _mw-bootstrap.js_ it is not quiet literate yet.
* [HelloWorld.md](HelloWorld.md) - A simple hello world example.
* [Example-1.md](Example-1.md) - An example of writing multiple code blocks that form a single JavaScript file.


