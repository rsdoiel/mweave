# Markdown Weave

This is an experiment in literarte programming concepts describe by
[Donald Knuth](http://www.literateprogramming.com/).  This implementation assumes the markup
syntax is valid Markdown such as that used on Github.  I have adopted one additional 
idea which is a link immediately preceeding a tripple quote block will be considered the
filename of the resulting source code.  If the filename is encountered more than once mw.js
should concatenate that with previous code associated with the filename. This exception is
this specific file since it is intended to be run by _mw-bootstrap.js_.

## Running mw-bootstrap.js on Markdown-Weaver.md

Here is the command which I plan to used to build _mw.js_ -

To generate _mw.js_ try the following command.

```Shell
  node mw-bootstrap.js Markdown-Weave.md
```

## mw.js

Ok, so here we go. Let us see if I can implement _mw.js_ from _mw-bootstrap.js_. 
mw.js creates an constructor called **Weave()**. **Weave()** generates a JavaScript
object with a parse method. **Weave.parse()** accepts Markdown source code as a string
and generates a new object which has properties with filenames to write an a list (i.e.
Array) of start and end line numbers to use in constructing the target file.

[mw.js](mw.js)
```JavaScript
  /**
   * mw.js - Markdown Weave, an exploration in Markdown using 
   * literate programming concepts.
   * @author R. S. Doiel, <rsdoiel@gmail.com>
   */
  /*jslint indent: 4 */
  /*global exports */
  function Weave () {
    return {
      parse: function (source) {
        var lines = source.split("\n"),
            filename = null,
            outputs = {},
            i = 0,
            j = 0;

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
                if (typeof outputs[filename] === "undefined") {
                    outputs[filename] = [];
                }
                outputs[filename].push({start: i + 1, end: -1});
            } else if (filename !== null && line.indexOf("```") === 0) {
                /* Find the last entry and add the end point */
                j = outputs[filename].length - 1;
                outputs[filename][j].end = i;
                filename = null;
            }
        };
        return outputs;
      }
    };
  }
  
  if (typeof exports !== "undefined") {
    exports.Weave = Weave;
  }
```

## mw_test.js

Here is some test code for see if mw.js works. This code relies on the YUI3 test module.

[mw_test.js](mw_test.js)
```JavaScript
  /**
   * mw_test.js - Test code for mw.js which was generated via mw-bootstrap.js.
   * @author R. S. Doiel, <rsdoiel@gmail.com>
   * copyright (c) 2013 all rights reserved
   * Licensed under BSD 2-clause license. See http://opensource.org/licenses/BSD-2-Clause
   */
  var YUI = require("yui").YUI,
      fs = require("fs"),
      mw = require("./mw");
  
  YUI({
    debug: true,
    useSync: true
  }).use("test", function (Y) {
    var testCase;
    
    testCase = new Y.Test.Case({
      name: "Simple testing for mw.js",
      "Should parse Markdown-Weave.md and yeild an object with what to write to disc": function () {
        var weave = new mw.Weave(),
          source = fs.readFileSync("Markdown-Weave.md").toString(),
          results = weave.parse(source);

          Y.log(results, "debug");
          Y.Assert.isObject(results);
          Y.Assert.isObject(results["mw.js"]);
          Y.Assert.isObject(results["mw.js"][0]);
          Y.Assert.areSame(31, results["mw.js"][0].start);
          Y.Assert.areSame(76, results["mw.js"][0].end);
      }
    });
    
    Y.Test.Runner.add(testCase);
    Y.Test.Runner.run();
  });
```


