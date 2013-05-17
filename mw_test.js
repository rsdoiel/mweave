/**
 * mw_test.js - Test code for mw.js which was generated via mw-bootstrap.js.
 * @author R. S. Doiel, <rsdoiel@gmail.com>
 * copyright (c) 2013 all rights reserved
 * Licensed under BSD 2-clause license. See http://opensource.org/licenses/BSD-2-Clause
 */
/*jslint node: true, indent: 4 */
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

            Y.Assert.isObject(results);
            Y.Assert.isObject(results["mw.js"]);
            Y.Assert.isObject(results["mw.js"][0]);
            // Remember array of lines cound from zero. End is inclusive.
            Y.Assert.areSame(79, results["mw.js"][0].start);
            Y.Assert.areSame(159, results["mw.js"][0].end);

            // Now try running on HelloWorld.md
            source = fs.readFileSync("HelloWorld.md").toString();
            results = weave.parse(source);
            Y.Assert.areSame(8, results["helloworld.js"][0].start);
            Y.Assert.areSame(8, results["helloworld.js"][0].end);
        },
        "Should render  a parsed object into a new object.": function () {
            var weave = new mw.Weave(),
                source = fs.readFileSync("Markdown-Weave.md").toString(),
                obj = weave.parse(source),
                results = weave.render(source, obj);

            Y.assert(source.length > 0, "Should have some markdown source");
            Y.Assert.isObject(obj["cli.js"]);
            Y.assert(obj["cli.js"][0].start > 0);
            Y.assert(obj["cli.js"][0].end > 0);

            Y.Assert.isObject(results);
            Y.Assert.isString(results["cli.js"]);

            // Now test our simple HelloWorld.md
            source = fs.readFileSync("HelloWorld.md").toString();
            obj = weave.parse(source);
            results = weave.render(source, obj);
            Y.Assert.isString(results["helloworld.js"]);
            Y.Assert.areEqual('console.log("Hello World");', results["helloworld.js"]);  
        }
    });

    Y.Test.Runner.add(testCase);
    Y.Test.Runner.run();
 });