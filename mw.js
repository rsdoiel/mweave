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
