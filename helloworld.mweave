
# Hello World

This is an example of "Hello World" implemented with JavaScript and _mweave.js_

[helloworld.js](helloworld.js)
```JavaScript
    console.log("Hello World");
```

Processing this with the command--

```Shell
    node cli.js HelloWorld.md
```

Should yield the script **helloworld.js** containing only

```JavaScript
    console.log("Hello World");
```
## Let us tes this!

We are just going to pop open the results and see if they are what we expect.

[test_helloworld.js](test_helloworld.js)
```JavaScript
    var fs = require("fs");
    fs.readFile("helloworld.js", function (err, buf) {
        var expected, s;
        if (err) {
            console.error("Test filed on helloworld.js: " + err);
            process.exit(1);
        }
        expected = 'console.log("Hello World");';
        s = buf.toString();
        if (s === expected) {
            console.log("Yippee! helloworld.js was rendered correctly!");
        } else {
            console.error("Something went wrong.");
            console.error("      s: [" + s + "]");
            console.error("expected [" + expected + "]");
            process.exit(1);
        }
    });
```

