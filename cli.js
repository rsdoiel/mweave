#!/usr/bin/env node
/**
 * cli.js - this is the command line tool for mweave command. It includes
 * binding mw.js to the file system.
 * @author R. S. Doiel, <rsdoiel@gmail.com>
 * copyright (c) 2013 all rights reserved
 */

var fs = require("fs"),
    path = require("path"),
    marked = require("marked"),
    opt = require("opt").create(),
    mw = require("./mw"),
    markdownFilename = "",
    documentDirectory = "",
    renderHTML = false;

opt.optionHelp("USAGE mweave MARKDOWN_FILENAME",
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
}, "Set the Markdown filename to process.");

opt.option(["-d", "--directory"], function (param) {
    if (param) {
        documentDirectory = param.trim();
    }
    opt.consume(param);
}, "Set the document directory to write to.");

opt.option(["-o", "--output"], function (param) {
    renderHTML = true;
    if (param) {
        htmlFilename = param.trim();
    }
    opt.consume(param);
}, "Render HTML from Markdown document as filename");

opt.option(["-h", "--help"], function (param) {
    opt.usage();
}, "Generate this help page.");

var argv = opt.optionWith(process.argv);
   
if (argv[2] !== undefined && markdownFilename === "") {
    markdownFilename = argv[2];
}

if (argv[3] !== undefined && htmlFilename === "") {
    htmlFilename = argv[3];
}

fs.readFile(markdownFilename, function (err, buf) {
    var obj,
        source,
        html,
        weave = new mw.Weave();

    if (err) {
        opt.usage(err, 1);
    }
    source = buf.toString();
    obj = weave.parse(source);
    results = weave.render(source, obj);

    Object.keys(results).forEach(function (filename) {
        console.log("Writing", path.join(documentDirectory, filename));
        fs.writeFile(path.join(documentDirectory, filename), results[filename]);
    });
    if (renderHTML === true) {
        marked.setOptions({
            gfm: true,
            tables: true,
            breaks: false,
            pedantic: false,
            sanitize: true,
            smartLists: true,
            langPrefix: 'language-',
            highlight: function(code, lang) {
                if (lang === 'js') {
                    return highlighter.javascript(code);
                }
                return code;
            }
        });
        html = marked(source);
        if (htmlFilename !== "") {
            console.log("Writing", path.join(documentDirectory, htmlFilename));
            fs.writeFile(path.join(documentDirectory, htmlFilename), html);
        } else {
            process.stdout.write(html);
        }
    }
});
