//
// mweave is an experiment inspired by Knuth's literate program,
// Markdown processors and Fountain screenplay text notation.
//
// @Author R. S. Doiel, <rsdoiel@gmail.com>
//
// Copyright (c) 2018, R. S. Doiel
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// * Redistributions of source code must retain the above copyright notice, this
//   list of conditions and the following disclaimer.
//
// * Redistributions in binary form must reproduce the above copyright notice,
//   this list of conditions and the following disclaimer in the documentation
//   and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
package main

import (
	"fmt"
	"os"

	// My packages
	"github.com/rsdoiel/mweave"

	// Caltech Library packages
	"github.com/caltechlibrary/cli"
)

var (
	description = `
mweave is and experimental literate programming tool. It is 
inspired by Knuth's ideas targeting Markdown like markup rather
than TeX.
`

	examples = `
generate source files from an mweave document

    mweave -i document.mweave -tangle

generate documentation from an mweave document

    mweave -i documemt.mweave -weave
`

	// Standard Options
	showHelp             bool
	showLicence          bool
	showVersion          bool
	showExamples         bool
	quitet               bool
	generateMarkdownDocs bool
	inputFName           string
	outputFName          string

	// Application Options
	weave  bool
	tangle bool
)

func main() {
	app := cli.NewCli(mweave.Version)
	fmt.Println(os.Args[0], "not implemented")
}
