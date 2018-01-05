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
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
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
	showLicense          bool
	showVersion          bool
	showExamples         bool
	quiet                bool
	newLine              bool
	generateMarkdownDocs bool
	inputFName           string
	outputFName          string

	// Application Options
	weave   bool
	tangle  bool
	astJSON bool
	astXML  bool
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
	app.BoolVar(&generateMarkdownDocs, "generate-markdown-docs", false, "generate Markdown documentation")
	app.StringVar(&inputFName, "i,input", "", "set input filename (the mweave file)")
	app.StringVar(&outputFName, "o,output", "", "set output filename")

	// Application Options
	app.BoolVar(&weave, "w,weave", false, "generate documentations files (e.g. Markdown output)")
	app.BoolVar(&tangle, "t,tangle", false, "generate source code files (e.g. program code)")
	app.BoolVar(&astJSON, "ast,json", false, "write out the AST of parsing the mweave file as JSON")
	app.BoolVar(&astXML, "xml", false, "write out the AST of parsing the mweave file as JSON")

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
	// Parse input
	mwDoc, err := mweave.Parse(src)
	cli.ExitOnError(app.Eout, err, quiet)

	if astJSON {
		src, err = json.MarshalIndent(mwDoc, "", "    ")
		cli.ExitOnError(app.Eout, err, quiet)
		fmt.Fprintf(app.Out, "%s", src)
	}

	if astXML {
		src, err = xml.MarshalIndent(mwDoc, "", "   ")
		cli.ExitOnError(app.Eout, err, quiet)
		fmt.Fprintf(app.Out, "%s", src)
	}

	if weave {
		err = mwDoc.Weave(app.Eout)
		if err != nil {
			fmt.Fprintf(app.Eout, "%s", err)
		}
	}

	if tangle {
		err = mwDoc.Tangle(app.Eout)
		if err != nil {
			fmt.Fprintf(app.Eout, "%s", err)
		}
	}

	if newLine {
		fmt.Fprintln(app.Out, "")
	}
}
