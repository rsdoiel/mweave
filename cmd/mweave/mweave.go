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
package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"

	// My packages
	"github.com/rsdoiel/mweave"
	"github.com/rsdoiel/shorthand"

	// Caltech Library packages
	"github.com/caltechlibrary/cli"
)

var (
	description = `
mweave is and experimental literate programming tool. It is 
inspired by Knuth's but much simplier and with the primary
purpose of creating a platform for writing interactive 
function (e.g. Adventure like text games).
`

	examples = `
generate source files from an mweave document

    mweave -i document.mweave -tangle

generate documentation from an mweave document

    mweave -i documemt.mweave -weave

display mweave parse results as XML

    mweave -i document.mweave -xml

display mweave parse results as JSON

    mweave -i document.mweave -xml
`

	// Standard Options
	showHelp             bool
	showLicense          bool
	showVersion          bool
	showExamples         bool
	quiet                bool
	newLine              bool
	generateMarkdownDocs bool
	inputFName           string
	outputFName          string

	// Application Options
	weave           bool
	tangle          bool
	docAsJSON       bool
	docAsXML        bool
	shorthandMacros bool
)

func main() {
	app := cli.NewCli(mweave.Version)

	// Add non-option parameter documentation
	app.AddParams("MWEAVE_FILENAME")

	// Add Help docs
	app.AddHelp("description", []byte(description))
	app.AddHelp("examples", []byte(examples))

	// Standard Options
	app.BoolVar(&showHelp, "h,help", false, "display help")
	app.BoolVar(&showLicense, "l,license", false, "display license")
	app.BoolVar(&showVersion, "v,version", false, "display version")
	app.BoolVar(&showExamples, "examples", false, "display examples")
	app.BoolVar(&quiet, "quiet", false, "suppress error messages")
	app.BoolVar(&newLine, "nl,newline", false, "add a trailing newline")
	app.BoolVar(&generateMarkdownDocs, "generate-markdown-docs", false, "generate Markdown documentation")
	app.StringVar(&inputFName, "i,input", "", "set input filename (the mweave file)")
	app.StringVar(&outputFName, "o,output", "", "set output filename")

	// Application Options
	app.BoolVar(&weave, "w,weave", false, "generate documentations files (e.g. Markdown output)")
	app.BoolVar(&tangle, "t,tangle", false, "generate source code files (e.g. program code)")
	app.BoolVar(&docAsJSON, "json", false, "write mweave doc as JSON")
	app.BoolVar(&docAsXML, "xml", false, "write mweave doc as XML")
	app.BoolVar(&shorthandMacros, "macros", false, "preprocess shorthand macros")

	// Process environment and options
	app.Parse()
	args := app.Args()

	if len(args) > 0 {
		inputFName = args[0]
	}
	if len(args) > 1 {
		outputFName = args[1]
	}

	// Setup IO
	var err error

	app.Eout = os.Stderr

	app.In, err = cli.Open(inputFName, os.Stdin)
	cli.ExitOnError(app.Eout, err, quiet)
	defer cli.CloseFile(inputFName, app.In)

	app.Out, err = cli.Create(outputFName, os.Stdout)
	cli.ExitOnError(app.Eout, err, quiet)
	defer cli.CloseFile(outputFName, app.Out)

	// Process options
	if generateMarkdownDocs {
		app.GenerateMarkdownDocs(app.Out)
		os.Exit(0)
	}
	if showHelp || showExamples {
		if len(args) > 0 {
			fmt.Fprintln(app.Out, app.Help(args...))
		} else {
			app.Usage(app.Out)
		}
		os.Exit(0)
	}
	if showLicense {
		fmt.Fprintln(app.Out, app.License())
		os.Exit(0)
	}
	if showVersion {
		fmt.Fprintln(app.Out, app.Version())
		os.Exit(0)
	}

	// ReadAll of input
	src, err := ioutil.ReadAll(app.In)
	cli.ExitOnError(app.Eout, err, quiet)

	// Pre-process shorthand Macros
	if shorthandMacros {
		macro := shorthand.New()
		src, err = macro.Apply(src, false)
		cli.ExitOnError(app.Eout, err, quiet)
	}

	// Parse input
	doc, err := mweave.Parse(src)
	cli.ExitOnError(app.Eout, err, quiet)

	switch {
	case docAsJSON:
		src, err = json.MarshalIndent(doc, "", "    ")
		cli.ExitOnError(app.Eout, err, quiet)
		fmt.Fprintf(app.Out, "%s", src)
		if newLine {
			fmt.Fprintln(app.Out, "")
		}
	case docAsXML:
		src, err = xml.MarshalIndent(doc, "", "   ")
		cli.ExitOnError(app.Eout, err, quiet)
		fmt.Fprintf(app.Out, "%s", src)
		if newLine {
			fmt.Fprintln(app.Out, "")
		}
	case weave:
		// Render Markdown outputs
		err = doc.Weave(app.Out)
		cli.ExitOnError(app.Eout, err, quiet)
	case tangle:
		err = doc.Tangle()
		cli.ExitOnError(app.Eout, err, quiet)
	default:
		app.Usage(app.Out)
	}
}
