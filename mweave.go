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
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
	"text/scanner"
)

const (
	// Version of package
	Version = `v0.1.1-dev`
	DocType = `mweave`

	// Constants used to identify type
	PlainText = iota
	EmptyBlock
	SourceBegin
	SourceEnd
)

type Document struct {
	XMLName  xml.Name   `json:"-"`
	DocType  string     `xml:"type,attr,omitempty" json:"doc_type,omitempty"`
	Version  string     `xml:"version,attr,omitempty" json:"version,omitempty"`
	Elements []*Element `xml:"elements>element,omitempty" json:"elements,omitempty"`
}

type Element struct {
	XMLName    xml.Name   `json:"-"`
	Type       int        `xml:"type,attr,omitempty" json:"type,omitempty"`
	LineNo     int        `xml:"line_no,attr,omitempty" json:"line_no,omitempty"`
	Attributes []xml.Attr `xml:",any,attr" json:"attr,omitempty"`
	Value      string     `xml:",chardata" json:"value,omitempty"`
	rawValue   []byte
}

func (elem *Element) MarshalJSON() ([]byte, error) {
	m := map[string]string{}
	m["type"] = strconv.Itoa(elem.Type)
	switch elem.Type {
	case PlainText:
		m["_type"] = "text/plain"
	case SourceBegin:
		m["_type"] = "directive/begin"
	case SourceEnd:
		m["_type"] = "directive/end"
	case EmptyBlock:
		m["_type"] = "text/empty"
	default:
		m["_type"] = "Unknown"
	}
	m["line_no"] = strconv.Itoa(elem.LineNo)
	if len(elem.Value) > 0 {
		m["value"] = elem.Value
	}
	for _, elem := range elem.Attributes {
		m[elem.Name.Local] = elem.Value
	}
	return json.Marshal(m)
}

func parseAttributes(src []byte, attributes []string) ([]xml.Attr, error) {
	var (
		attrs []xml.Attr
		s     scanner.Scanner
	)
	// Trim off <!--mweave:begin and -->
	src = bytes.TrimPrefix(src, []byte("<!--mweave:begin "))
	src = bytes.TrimSuffix(src, []byte(" -->"))

	s.Init(bytes.NewReader(src))
	i := 0
	for tok := s.Scan(); tok != scanner.EOF && i < len(attributes); tok = s.Scan() {
		if i < len(attributes) {
			attrs = append(attrs, xml.Attr{Name: xml.Name{Local: attributes[i]}, Value: s.TokenText()})
		} else {
			attrs = append(attrs, xml.Attr{Name: xml.Name{Local: "Unknown"}, Value: s.TokenText()})
		}
		i++
	}
	if i > len(attributes) {
		return attrs, fmt.Errorf("unexpected attribute, %+v", attrs[i:])
	}
	return attrs, nil
}

func Parse(src []byte) (*Document, error) {
	var err error
	doc := new(Document)
	doc.DocType = "mweave"
	doc.Version = Version
	//NOTE: This is a naive implementation based on analysing individual lines.
	lines := bytes.Split(src, []byte("\n"))
	for i, line := range lines {
		switch {
		case bytes.HasPrefix(line, []byte("<!--mweave:begin")):
			// Save our directive
			elem := new(Element)
			elem.Type = SourceBegin
			elem.LineNo = i
			elem.rawValue = bytes.TrimSpace(line)
			elem.Value = "mweave:begin"
			elem.Attributes, err = parseAttributes(elem.rawValue, []string{"filename", "index"})
			if err != nil {
				return doc, err
			}
			doc.Elements = append(doc.Elements, elem)
		case bytes.HasPrefix(line, []byte("<!--mweave:end")):
			// Save our directive
			elem := new(Element)
			elem.Type = SourceEnd
			elem.LineNo = i
			elem.rawValue = bytes.TrimSpace(line)
			elem.Value = "mweave:end"
			elem.Attributes, err = parseAttributes(elem.rawValue, []string{})
			doc.Elements = append(doc.Elements, elem)
			if err != nil {
				return doc, err
			}
		default:
			//Create a the next PlainText Element
			elem := new(Element)
			if len(bytes.TrimSpace(line)) == 0 {
				elem.Type = EmptyBlock
			} else {
				elem.Type = PlainText
			}
			elem.LineNo = i
			elem.rawValue = line
			elem.Value = fmt.Sprintf("%s\n", elem.rawValue)
			doc.Elements = append(doc.Elements, elem)
		}
	}
	return doc, nil
}

// Weave will transform the weave document into a plain text document.
func (doc *Document) Weave(out io.Writer) error {
	if len(doc.Elements) == 0 {
		return fmt.Errorf("nothing to weave")
	}
	for _, elem := range doc.Elements {
		if elem.Type == PlainText || elem.Type == EmptyBlock {
			fmt.Fprint(out, elem.Value)
		}
	}
	return nil
}

func assemble(m map[string]string) []byte {
	var (
		keys      []string
		buf       []byte
		skip      bool
		shiftLeft bool
	)
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	//last := len(keys) - 1
	for _, key := range keys {
		if val, ok := m[key]; ok == true {
			for _, s := range strings.Split(val, "\n") {
				if shiftLeft == false && strings.HasPrefix(s, "```") {
					shiftLeft = true
					skip = true
				} else if shiftLeft == true && strings.HasPrefix(s, "```") {
					shiftLeft = false
					skip = true
				}
				if shiftLeft {
					s = strings.TrimPrefix(s, "    ")
				}
				if skip {
					skip = false
				} else {
					buf = append(buf, []byte(s+"\n")...)
				}
			}
		}
	}
	return buf
}

func getAttribute(attrs []xml.Attr, key string) (string, bool) {
	for _, attr := range attrs {
		if attr.Name.Local == key {
			return strings.Trim(attr.Value, "\""), true
		}
	}
	return "", false
}

// Tangle processes a Document stuct and renders source code
// identified with <!--mweave:begin --> directives.
func (doc *Document) Tangle() error {
	var (
		fName string
		index string
		ok    bool
		tdocs map[string]map[string]string
	)

	// collect all the tangled docs
	tdocs = make(map[string]map[string]string)
	fName = ""
	index = ""

	// collect the socs to rangle out
	for _, elem := range doc.Elements {
		switch elem.Type {
		case SourceBegin:
			fName, ok = getAttribute(elem.Attributes, "filename")
			if ok == false {
				return fmt.Errorf("missing doc name for mweave:begin at line %d", elem.LineNo)
			}
			index, ok = getAttribute(elem.Attributes, "index")
			if ok == false {
				return fmt.Errorf("missing doc index for mweave:begin at line %d", elem.LineNo)
			}
			// NOTE: we need to left pad index with zero since we're
			// going to need to sort the string eventually.
			if i, err := strconv.Atoi(index); err == nil {
				index = fmt.Sprintf("%010d", i)
			} else {
				return fmt.Errorf("was expecting an integer value for index, got %q", index)
			}
		case SourceEnd:
			fName = ""
			index = ""
		default:
			if fName != "" {
				if parts, ok := tdocs[fName]; ok == true {
					if src, ok := parts[index]; ok == true {
						parts[index] = src + elem.Value
					} else {
						parts[index] = elem.Value
					}
					tdocs[fName] = parts
				} else {
					parts := make(map[string]string)
					parts[index] = elem.Value
					tdocs[fName] = parts
				}
			}
		}
	}

	if len(tdocs) == 0 {
		return fmt.Errorf("nothing to tangle")
	}
	// assemble the tangled docs
	for fName, parts := range tdocs {
		err := ioutil.WriteFile(fName, assemble(parts), 0664)
		if err != nil {
			return fmt.Errorf("error writing %s, %s", fName, err)
		}
	}
	return nil
}
