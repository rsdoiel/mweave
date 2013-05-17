# Markdown Weave

## Rational

This is an experiment in [literarte programming](http://en.wikipedia.org/wiki/Literate_programming)
using Markdown as the markup language.  While it does not attempt all the ideas put forward
by [Donald Knuth](http://www.literateprogramming.com/) it proceed from the point of view that
that creating a document that is human readable helps clarify thinking on what needs to be executed
by the computer. While the core of computers have not changed (i.e. they still have volitile memory,
non-volitile memory and one or more computational units) the tools, languages and manner in which
they are applied has grown. Today there are many more programming languages. The distinction between
instructional languages and "professional" languages is largely gone. Another change is the 
breadth of tools available for building applications. Many editors today have built in linters,
color coding of the programming text and live views of the results. Additionally the notion
of typesetting has morphed.  When Knuth first wrote is tools, his [TeX](http://en.wikipedia.org/wiki/TeX)
was the masterful land capable of generating primary documentation via paper printouts.  Today
were are seen computer documentation spread via the web rather then reams of paper. Still the idea
of writing prose and code has strong attraction. 

## Weave/Tangle reconsidered

One of the challenges I repeatily run into is transmitting to my colleagues my understanding when
creating web APIs and other systems.  It is more then simple code documentation I would get form
a system like javadoc, doxygen, or yuidoc. Those tools are very helpful when you already know
how things work (e.g. looking up a class method, looking inside how a method was written). What
is missing is the narrative in how the system came about and the tutorials need to get familiar
with the system.  I have adopted the common "best practices" of including a README.md, an INSTALLATION.md
and assortment of example usages and usually several additional Markdown documents explaining 
things.  This is helpful when the program or API is first deployed.  Keeping them currrent is
difficult if I am the only coding. Expecting others to keep them current is not very reasonable.
Not only would they need to update the source code they need to find all the places I might have
talked about it in my prose. So why not stick the code in my prose?  This is done all the time
on github with documentation written in Markdown.  The only peice missing is generating the
program code from the markdown documents.  _mweave.js_ is an attempt to explore that. It is
far simpler then learning the markup of [cweb](http://www.literateprogramming.com).  It can
be applied equally to languages commonly shared on Github. It should be a small tool
in the toolchain that can be leverage to get us a little further down the road to maintainable
source code.


My experiment, _mweave.js_, differs from Literate Programming definition in a couple of ways

* No change to the syntax of Markdown as practiced on Github
* No Macro language
* Order of explanation DOES reflect the output of the files you generate
* The rendered document includes the code examples from "code" blocks that are immediately precede by links
* The render program includes any embedded commments as they were in the code blocks from the source file
* _mweave.js_ is intended not as a complete system but another helper tool like _jslint_ or _markdown_
* The code generated should look like the code in the original document just out-dented by 4 spaces
* I should be able to use this approach to bootstrap a better _mweave.js_ tool


## Running mw-bootstrap.js on Markdown-Weaver.md

You need to generate _mw-bootstrap.js_.  You can do this with four Unix commands ---

```Shell
    npm install shelljs
    vi -e -c "20,81wq! mw-bootstrap.js" README.md
    sed -e "s/    //" -i mw-bootstrap.js 
    chmod 770 mw-bootstrap.js
    ./mw-bootstrap.js Markdown-Weave.md
```


## mw.js

Ok, so here we go. Let us see if I can implement _mw.js_ from _mw-bootstrap.js_ processing
this file.   mw.js creates an constructor called **Weave()**. **Weave()** generates a JavaScript
object with a parse method and render method. **Weave.parse()** accepts Markdown source code as a string
and generates a new object which has properties with filenames to write an a list (i.e.
Array) of start and end line numbers to use in constructing the target file. **Weave.render()**
takes the original Markdown source code along with the parse results and renders a new object
containing properties corresponding with the extracted source code. The object will need further
processing to be written out to disc.

[mw.js](mw.js)
```JavaScript
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
```

### mw_test.js

Here is some test code for see if mw.js works. This code relies on the YUI3 test module.

[mw_test.js](mw_test.js)
```JavaScript
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
```

### design choices

So why stop just before rendering text to disc? Because it may be helpful to use _mweave.js_ with outer 
browser based tools (e.g. CodeMirror, Ace). Additionally NodeJS (where this will likely run) resents
an event module for I/O and leveraging that in a wrapper of this library (e.g. _cli.js_) makes the most
sense to me at this stage.


## Biulding command-line tool

The command line tool provides the bindings to file IO and processing of command line options.

[cli.js](cli.js)
```JavaScript
    #!/usr/bin/env node
    /**
     * cli.js - this is the command line tool for mweave command. It includes
     * binding mw.js to the file system.
     * @author R. S. Doiel, <rsdoiel@gmail.com>
     * copyright (c) 2013 all rights reserved
     */

    var VERSION = "0.0.2", 
        fs = require("fs"),
        path = require("path"),
        handlebars = require("handlebars"),
        marked = require("marked"),
        opt = require("opt").create(),
        mw = require("./mw"),
        markdownFilename = "",
        documentDirectory = "",
        handlebarsTemplate = "",
        jsonFilename = "",
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
    
    opt.option(["-j", "--json"], function (param) {
        if (param) {
            jsonFilename = param.trim();
        }
        opt.consume(param);
    }, 'Use JSON file for additional content when rendering a template. (e.g. {"title":"My Webpage"})');
    opt.option(["-t", "--template"], function (param) {
        if (param) {
            handlebarsTemplate = param.trim();
        }
        opt.consume(param);
    }, "Use the handlebars template when rendering HTML.");
    opt.option(["-o", "--output"], function (param) {
        renderHTML = true;
        if (param) {
            htmlFilename = param.trim();
        }
        opt.consume(param);
    }, "Render HTML from Markdown document as filename");

    opt.option(["-v", "--version"], function (param) {
        console.log("Version ", VERSION);
        process.exit(0);
    }, "Show the version number");

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
            template_source,
            html,
            data = {
                title: markdownFilename,
                content: null
            },
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
            if (handlebarsTemplate !== "") {
                if (jsonFilename !== "") {
                    data = JSON.parse(fs.readFileSync(jsonFilename).toString());
                }
                template_source = fs.readFileSync(handlebarsTemplate).toString();
                template = handlebars.compile(template_source);
                data.content = html;
                html = template(data);
            }
            if (htmlFilename !== "") {
                console.log("Writing", path.join(documentDirectory, htmlFilename));
                fs.writeFile(path.join(documentDirectory, htmlFilename), html);
            } else {
                process.stdout.write(html);
            }
        }
    });
```

# Misc support scripts

## Node packaging of _mweave.js_

[package.json](package.json)
```JavaScript
    {
        "name": "mweave",
        "version": "0.0.2",
        "description": "This is an experiment in using Markdown and some concepts from literate programming.",
        "main": "mw.js",
        "bin": {
            "mweave": "./cli.js"
        },
        "scripts": {
            "build": "node build.js",
            "test": "node tests.js"
        },
        "optionalDependencies": {
            "yui": "3.10.x",
            "yuitest": "0.7.x"
        },
        "dependencies": {
            "handlebars": "1.0.x",
            "marked": "0.2.x",
            "opt": "0.1.x",
            "shelljs": "0.1.x"
        },
        "repository": {
            "type": "git",
            "url": "git@github.com:rsdoiel/mweave.git"
        },
        "keywords": [
          "markdown",
          "weave",
          "javascript"
        ],
        "engines": {
            "node": "0.10.x",
            "npm": "1.2.x"
        },
        "files": [
            "README.md",
            "Markdown-Weave.md",
            "Example-1.md",
            "HelloWorld.md"
        ],
        "author": "R. S. Doiel",
        "license": "BSD",
        "readmeFilename": "README.md"
    }
```

## Run Some Tests

Software should have a sense of if it is working. This is done by tests.
We can use the bootstrap process to see if the basics of _mweave_ is working.

[tests.js](tests.js)
```JavaScript
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
```

