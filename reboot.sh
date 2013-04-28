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

    if [ -f "cli.js" ] && [ -f "HelloWorld.md" ]; then
        node cli.js HelloWorld.md
        node test_helloworld.js
    fi
