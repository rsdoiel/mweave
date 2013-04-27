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
          Y.Assert.areSame(34, results["mw.js"][0].start);
          Y.Assert.areSame(79, results["mw.js"][0].end);
      }
    });
    
    Y.Test.Runner.add(testCase);
    Y.Test.Runner.run();
  });
