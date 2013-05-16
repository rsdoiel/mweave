#!/usr/bin/env node
/**
 * mw-bootstrap.js - an experiment in literate style programming in a 
 * markdown file.
 * @author R. S. Doiel, <rsdoiel@gmail.com>
 * copyright (c) 2013 all rights reserved
 * Licensed under the BSD 2-clause license. See http://opensource.org/licenses/BSD-2-Clause
 */
 require("shelljs/global"); 
 var fs = require("fs"),
    lines = [],
    line = "",
    check = "",
    outputs = {},
    i = 0,
    markdownFilename = "README.md",
    filename,
    start = 0,
    end = 0;

 function exportLines(lines, outFilename, start, end) {
     fs.writeFile(outFilename, function (err) {
         if (err) {
             console.error(err);
             process.exit(1);
         }
         sed("-i", /    /,"", outFilename);
     });
 }

 if (process.argv.length === 3) {
    markdownFilename = process.argv[2];
 }
 lines = fs.readFileSync(markdownFilename).toString().split(/\n|\r\n/);
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
        outputs[filename] = {start: i + 1, end: -1};
    } else if (typeof outputs[filename] !== "undefined" &&
        outputs[filename].end < 0 &&
        line.indexOf("```") === 0) {
        outputs[filename].end = i;
        filename = "";
    }
 };
 Object.keys(outputs).forEach(function (ky) {
     exportLines(lines, ky, outputs[ky].start, outputs[ky].end);
    /*
     console.log("# This vi command to generate the code for " + ky);
     console.log("vi -e -c '" + outputs[ky].start + "," + outputs[ky].end + " wq! " + ky + "' " +
        markdownFilename);
     console.log('sed -e "s/    //" -i ' + ky);
     */
 });
