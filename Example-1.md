
# Example 1, write code as a series of blocks

One of the tenants of literate programing is the ability
to write prose then generate code. _mweae_ uses Markdown's
code blocks for that.

[Here is our first block](example-1.js)
```JavaScript
    /**
     * example-1.js
     */
    /*jslint node: true, indent: 4 */
```

In this block I just setup a command header and a _jslint_ directive.

[Now lets get some code going](example-1.js)
```JavaScript
    console.log("Caculating 1 + 1");
    var accumulator = 0;
    accumulator = 1 + 1;
```

Now that we done a caculation let us send that to the console.

[Write the results](example-1.js)
```JavaScript
    console.log("And the answer is", accumulator);
```

Finally let us tell the viewer we are done.

[Last bits of program](example-1.js)
```JavaScript
    console.log("All done!");
    process.exit(0);
```

Running this through _mweave_ should produce a single file called _example-1.js_

