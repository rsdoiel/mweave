
## reboot.sh

Ideally I would like to beable to rebuild the system into useable state
from a simple application of the bootstrap program listed in [README.md](README.md).
This shell script illustrates the commands need to build the system from scratch.

[reboot.sh](reboot.sh)
```Shell
    #!/bin/bash

    #
    # Generate JavaScript source files.
    #
    echo "This is a shell script executing the commands to bootstrap mw.js"
    echo "Running the vi command to pull mw-bootstrap.js out of README.md"
    vi -e -c "20,67wq! mw-bootstrap.js" README.md
    echo "Running mw-bootstrap.js on Markdown-Weave.md"
    node mw-bootstrap.js Markdown-Weave.md > tmp.sh
    echo "Running the suggested vi commands to make mw.js and mw_test.js"
    . tmp.sh
    echo "Create cli.js for command line program."
    node mw-bootstrap.js Misc.md > tmp.sh
    . tmp.sh
    rm tmp.sh

    #
    # Setup and run some testing.
    #
    if [ -f "mw.js" ];then
        echo "Found mw.js"
    else
        echo "Missing mw.js, something went wrong."
        exit 1
    fi
    if [ -f "mw_test.js" ]; then
        node mw_test.js
    else
        echo "Missing mw_test.js, something went wrong."
        exit 1
    fi
```

## Node packing of Markdown-Weave


[package.json](package.json)
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
          "opt": "0.1.x",
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


## Biulding command-line tool

The command line tool provides the bindings to file IO.

[cli.js](cli.js)
```JavaScript
    /**
     * cli.js - this is the command line tool for mweave command. It includes
     * binding mw.js to the file system.
     * @author R. S. Doiel, <rsdoiel@gmail.com>
     * copyright (c) 2013 all rights reserved
     */

    var fs = require("fs"),
        path = require("path"),
        opt = require("opt").create(),
        mw = require("./mw"),
        markdownFilename = "";

    opt.optionHelp("USAGE mweave MARKDOWN_FILENAME",
        "SYNOPSIS: Process the markdown file listed on the command line and render any" +
        "source files defined in it.",
        "OPTIONS",
        " copyright (c) 2013 all rights reserved\n" +
        " Released under the BSD 2-clause license\n" + 
        " See : http://opensource.org/licenses/bsd-license.php\n");

     
    opt.consume();
    opt.option(["-i", "--input"], function (param) {
        if (param) {
            markdownFilename = param.trim();
        }
        opt.consume(param);
    }, "Set the Markdown file to process.");
    opt.option(["-h", "--help"], function (param) {
        opt.usage();
    }, "Generate this help page.");

    var argv = opt.optionWith(process.argv);
   
    console.log("DEBUG argv:", argv);
    if (argv[2] !== undefined) {
        markdownFilename = argv[2];
    }

    console.log("DEBUG processing:", markdownFilename);
    fs.readFile(markdownFilename, function (err, buf) {
        var obj,
            source,
            weave = new mw.Weave();

        if (err) {
            opt.usage(err, 1);
        }
        source = buf.toString();
        obj = weave.parse(source);
        console.log("DEBUG obj", obj);// DEBUG
        results = weave.parse(source, obj);
        console.log("DEBUG results:", results);// DEBUG
    });
```
