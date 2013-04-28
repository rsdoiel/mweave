#!/usr/bin/env node
/**
 * cli.js - this is the command line tool for mweave command. It includes
 * binding mw.js to the file system.
 * @author R. S. Doiel, <rsdoiel@gmail.com>
 * copyright (c) 2013 all rights reserved
 */

var fs = require("fs"),
    opt = require("opt").create(),
    mw = require("./mw"),
    markdownFilename = "";

opt.optionHelp("USAGE mweave -i MARKDOWN_FILENAME",
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
   
if (argv[2] !== undefined && markdownFilename === "") {
    markdownFilename = argv[2];
}

fs.readFile(markdownFilename, function (err, buf) {
    var obj,
        source,
        weave = new mw.Weave();

    if (err) {
        opt.usage(err, 1);
    }
    source = buf.toString();
    obj = weave.parse(source);
    results = weave.render(source, obj);

    Object.keys(results).forEach(function (filename) {
        console.log("Writing", filename);
        fs.writeFile(filename, results[filename]);
    });
});
