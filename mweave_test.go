//
// mweave is an experiment inspired by Knuth's literate program,
// Markdown processors and Fountain screenplay text notation.
//
// @Author R. S. Doiel, <rsdoiel@gmail.com>
//
// Copyright (c) 2018, R. S. Doiel
// All rights not granted herein are expressly reserved by R. S. Doiel.
//
// Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
package mweave

import (
	"testing"
)

func TestAssemble(t *testing.T) {
	text := []byte(`

# Hello World!

<!--mweave:macro set <<NAME>> "George!" -->
<!--mweave:macro bash <<TIME>> -->
echo -n $(date "+%I:%M")
<!--mweave:end -->

This is an example of an embedded document to be extracted by 
[mweave](https://github.com/rsdoiel/mweave).

<!--mweave:source "testout/helloworld.py" 0 -->
` + "```python" + `
    #!/usr/bin/env python3
    print("Hello World!")
	print("Hello, <<NAME>>")
	print("It is <<TIME>>")
` + "```" + `
<!--mweave:end -->
`)

	_, err := Parse(text)
	if err != nil {
		t.Errorf("Parse failed, %s", err)
	}

}
