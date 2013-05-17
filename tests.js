#!/usr/bin/env node
/**
 * tests.js - Run tests to see if _mweave_ is working to some degree.
 */
require("shelljs/global");
var YUI = require("yui").YUI;
// Run mw_test.js first.
if (exec("node mw_test.js").code !== 0) {
    echo("node mw_test.js returned a non-zero status. :-(");
    process.exit(1);
}
// Now run cli.js on Markdown-Weave.md, see if 
// Sandbox is updated correctly.
YUI({
    useSync: true,
    debug: true
}).use("test", function (Y) {
    var TestCase = new Y.Test.Case({
        name: "Test the output of cli.js",
        "Should produce some files.": function () {
            var Assert = Y.Test.Assert;

            Assert.isTrue(test("-f", "mw.js"));
            Assert.isTrue(test("-f", "mw_test.js"));
            Assert.isTrue(test("-f", "tests.js"));
            Assert.isTrue(test("-f", "cli.js"));
            Assert.isTrue(test("-f", "package.json"));
        }
    });

    Y.Test.Runner.add(TestCase);
    Y.Test.Runner.run();
});