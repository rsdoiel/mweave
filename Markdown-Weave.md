# Markdown Weave

Is an experiment in Donald Knuth's literarte programming ideas.  This version assumes the markup syntax
is valid Markdown such as that used on Github.  I've adopted one additional idea which
is a link immediately preceeding a tripple quote block will be considered the filename of the resulting
source code.  If the filename is encountered more than once mw.js should concatenate that with 
previous code associated with the filename.  Since this file is meant to be run by _mw-bootstrap.js_
it will be written so _mw.js_ is in one quote block.  Additional functionality will be folded in as
node modules.  the resulting _mw.js_ should support my original intent allowing you to code some,
write some prose and then code some more accumulating results in the appropriate file.

## Running mw-bootstrap.js on Markdown-Weaver.md

Here's the command which I plan to used to build _mw.js_ -

```Shell
  node mw-bootstrap.js Markdown-Weave.md
```

## mw.js

Ok, so here we go. Let's see if I can implement _mw.js_ from _mw-bootstrap.js_.

[mw.js](mw.js)
```JavaScript
  /**
   * mw.js - Markdown Weave, an exploration in Markdown using Donald
   * Knuth's literate programming concepts.
   * @author R. S. Doiel, <rsdoiel@gmail.com>
   */
  /*jslint indent: 4 */
  /*global exports */
  function Weave () {
    return {
      parse: function () {
        throw "Weave.parse() not implemented.";
      }
    };
  }
  
  if (typeof exports !== "undefined") {
    exports.Weave = Weave;
  }
```

## mw_test.js

Here's some test code for see if mw.js works. I'm using YUI3's test module to test ms.js.

[mw_test.js](mw_test.js)
```JavaScript
  /**
   * mw_test.js - Test code for mw.js which was generated via mw-bootstrap.js.
   * @author R. S. Doiel, <rsdoiel@gmail.com>
   * copyright (c) 2013 all rights reserved
   * Licensed under BSD 2-clause license. See http://opensource.org/licenses/BSD-2-Clause
   */
  var YUI = require("yui"),
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
          source = fs.readFileSync("Markdown-Weave.md"),
          results = weave.parse(source);

          Y.log(results, "debug");
          Y.Assert.isObject(results);
          Y.Assert.isObject(results["mw_test.js"]);
          Y.Assert.isObject(results["mw_test.js"][0]);
          Y.Assert.areSame(results["mw_test.js"][0].start, 28);
          Y.Assert.areSave(results["mw_test.js"][0].end, 62);
      }
    });
    
    Y.Test.Runner.add(testCase);
    Y.Test.Runner.run();
  });
```
