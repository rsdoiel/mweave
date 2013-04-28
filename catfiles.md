catfiles.js
===========

A simple example to asynchronous file concatenation where the names and numbers of the discrete files are known in advance. 

# Problem

You want to have a simple way of concatenating file resources stored locally
and "getable" via http.  You don't want to block since this is written in JavaScript running in Node. How do you do this simply?


This is how the solution using it might work.

```JavaScript
    var cat = require("catfiles"),
        concatentated_content = "";
    
    cat.cat(["css/basic.css", "css/browsers.css", "css/custom.css"],
        function (err, buf) {
            if (err) {
                throw err;
            }
            // Here's out concatenated content
            console.log(buf.toString());
        });
```

## Local filesystem only

Implementation might look something like...

```JavaScript
    var fs = require("fs");

    exports.cat = function (filenames, callback) {
        var count = filenames.length,
            output = new Array(count),
            i = 0,
            processed = 0;
         
        function readFile(filename, output_index, callback) {
            fs.readFile(filename, function (err, buf) {
                var error_msg, i = 0, total_size = 0;
                if (err) {
                    error_msg = [filename, ": ", err].join(""): 
                    callback(error_msg, null);
                    return;
                }
                output[output_index] = buf;
                processed += 1;
                if (processed === count) {
                    // Get the length of the concatenated buffer you'll need
                    for (i = 0; i < output.length; i += 1) {
                        total_size += output[i].length;
                    }
                    // then pass new Buffer.concat(list, total_size);
                    // to the callback
                    callback(null, new Buffer.concat(output, total_size));
                }
            });
        }
        
        for (i = 0; i < count; i += 1) {
            readFile(filenames[i], i, callback);
        }
    };
```

## Remote files only

```JavaScript
    var url = require("url"),
        http = require("http"),
        https = require("https");

    exports.cat = function (urls, callback) {
        var count = urls.length,
            output = new Array(count),
            i = 0,
            processed = 0;
         
        function readFile(href, output_index, callback) {
            var parts = url.parse(href),
                ht;
            
            if (parts.protocol === "https") {
                ht = https;
            } else {
                ht = http;
            }
            
            ht.get(href, function (res) {
                var i = 0, total_size = 0;

                output[output_index] = buf;
                processed += 1;
                if (processed === count) {
                    // Get the length of the concatenated buffer you'll need
                    for (i = 0; i < output.length; i += 1) {
                        total_size += output[i].length;
                    }
                    // then pass new Buffer.concat(list, total_size);
                    // to the callback
                    callback(null, new Buffer.concat(output, total_size));
                }
            }).on("error", function (err) {
                var error_msg;
                if (err) {
                    error_msg = [href, ": ", err].join(""): 
                    callback(error_msg, null);
                    return;
                }
            });
        }
        
        for (i = 0; i < count; i += 1) {
            readFile(urls[i], i, callback);
        }
    };
```

## Combining both local and remote files

[catfiles.js](catfiles.js)
```JavaScript
    var url = require("url"),
        http = require("http"),
        https = require("https");

    exports.cat = function (urls, callback) {
        var count = urls.length,
            output = new Array(count),
            i = 0,
            processed = 0;
         
        function readLocalFile(filename, output_index, callback) {
            fs.readFile(filename, function (err, buf) {
                var error_msg, i = 0, total_size = 0;
                if (err) {
                    error_msg = [filename, ": ", err].join(""): 
                    callback(error_msg, null);
                    return;
                }
                output[output_index] = buf;
                processed += 1;
                if (processed === count) {
                    // Get the length of the concatenated buffer you'll need
                    for (i = 0; i < output.length; i += 1) {
                        total_size += output[i].length;
                    }
                    // then pass new Buffer.concat(list, total_size);
                    // to the callback
                    callback(null, new Buffer.concat(output, total_size));
                }
            });
        }

        function readRemoteFile(href, output_index, callback) {
            var parts = url.parse(href),
                ht;
            
            if (parts.protocol === "https") {
                ht = https;
            } else {
                ht = http;
            }
            
            ht.get(href, function (res) {
                var i = 0, total_size = 0;

                output[output_index] = buf;
                processed += 1;
                if (processed === count) {
                    // Get the length of the concatenated buffer you'll need
                    for (i = 0; i < output.length; i += 1) {
                        total_size += output[i].length;
                    }
                    // then pass new Buffer.concat(list, total_size);
                    // to the callback
                    callback(null, new Buffer.concat(output, total_size));
                }
            }).on("error", function (err) {
                var error_msg;
                if (err) {
                    error_msg = [href, ": ", err].join(""): 
                    callback(error_msg, null);
                    return;
                }
            });
        }
        
        for (i = 0; i < count; i += 1) {
            if (urls[i].indexOf("://") > -1) {
                readRemoteFile(urls[i], i, callback);
            } else {
                readLocalFile(urls[i], i, callback);
            }
        }
    };
```
