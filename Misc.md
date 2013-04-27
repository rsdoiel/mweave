
## reboot.sh

Ideally I would like to beable to rebuild the system into useable state
from a simple application of the bootstrap program listed in [README.md](README.md).
This shell script illustrates the commands need to build the system from scratch.

[reboot.sh](reboot.sh)
```Shell
    #!/bin/bash
    echo "This is a shell script executing the commands to bootstrap mw.js"
    echo "Running the vi command to pull mw-bootstrap.js out of README.md"
    vi -e -c "20,67wq! mw-bootstrap.js" README.md
    echo "Running mw-bootstrap.js on Markdown-Weave.md"
    node mw-bootstrap.js Markdown-Weave.md > tmp.sh
    echo "Running the suggested vi commands to make mw.js and mw_test.js"
    . tmp.sh
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
    rm tmp.sh
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


