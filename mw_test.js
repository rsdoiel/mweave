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
      "Should parse Markdown-Weave.md and yeild a new object": function () {
        var weave = new mw.Weave(),
          source = fs.readFileSync("Markdown-Weave.md").toString(),
          results = weave.parse(source);

          Y.log(results, "debug");
          Y.Assert.isObject(results);
          Y.Assert.isObject(results["mw.js"]);
          Y.Assert.isObject(results["mw.js"][0]);
          Y.Assert.areSame(46, results["mw.js"][0].start);
          Y.Assert.areSame(96, results["mw.js"][0].end);
      },
      "Should render  a parsed object into a new object.": function () {
          var weave = new mw.Weave(),
            source = fs.readFileSync("Misc.md").toString(),
            obj = weave.parse(source),
            results = weave.render(obj);

          Y.log(obj, "debug");
          Y.log(results, "debug");
          Y.assert(source.length > 0, "Should have some markdown source");
          Y.Assert.isObject(obj["cli.js"]);
          Y.assert(obj["cli.js"].start > 0);
          Y.assert(obj["cli.js"].end > 0);

          Y.Assert.isObect(results);
          Y.Assert.isString(results["cli.js"]);
      }
    });
    
    Y.Test.Runner.add(testCase);
    Y.Test.Runner.run();
  });
