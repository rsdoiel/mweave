/**
 * mw.js - Markdown Weave, an exploration in Markdown using 
 * literate programming concepts.
 * @author R. S. Doiel, <rsdoiel@gmail.com>
 */
/*jslint indent: 4 */
/*global exports */
function Weave() {
    return {
        parse: function (source) {
            var lines = source.split("\n"),
                filename = null,
                outputs = {},
                start_cut = 0,
                end_cut= 0,
                i = 0,
                j = 0;

            for (i = 0; i < lines.length; i += 1) {
                line = lines[i];
                check = line.trim();
                if (i < lines.length - 2 &&
                        lines[i + 1].indexOf("```") === 0 &&
                        check[0] === '[' && check[check.length - 1] === ')') {
                    // Now skip ahead to lines of actual code.
                    i += 2;
                    start_cut = check.lastIndexOf('(') + 1;
                    end_cut = check.lastIndexOf(')');
                    filename = line.substr(start_cut, end_cut - start_cut);
                    if (typeof outputs[filename] === "undefined") {
                        outputs[filename] = [];
                    }
                    // I am storing line numbers, not index into lines.
                    // Start and End points are inclusive.
                    outputs[filename].push({start: i + 1, end: -1});
                } else if (filename !== null && line.indexOf("```") === 0) {
                    /* Find the last entry and add the end point */
                    j = outputs[filename].length - 1;
                    outputs[filename][j].end = i;
                    filename = null;
                }
            }
            return outputs;
        },
        render: function (source, parsed) {
            var lines = source.split("\n"),
                filenames = Object.keys(parsed),
                outputs = {};

            function catSource(points) {
                var output = [];
                points.forEach(function (point) {
                    var i, start, end, outdent = 4;
                    // Convert from line numbers to array index
                    start = point.start - 1;
                    end = point.end - 1;
                    // end is inclusive.
                    for (i = start; i <= end && i < lines.length; i += 1) {
                        outdent = 0;
                        if (lines[i].indexOf("    ") === 0) {
                            outdent = 4;
                        } else if (lines[start].indexOf("\t") === 0) {
                            outdent = 1;
                        }
                        output.push(lines[i].substr(outdent));
                    }
                });
                return output.join("\n");
            }

            filenames.forEach(function (filename) {
                outputs[filename] = catSource(parsed[filename]);
            });
            return outputs;
        }
    };
}
  
if (typeof exports !== "undefined") {
    exports.Weave = Weave;
}
