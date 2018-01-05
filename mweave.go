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
package mweave

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"strconv"
	"text/scanner"
)

const (
	// Version of package
	Version = `v0.1.0-dev`
	DocType = `mweave`

	// Constants used to identify type
	PlainText = iota
	EmptyBlock
	DirectiveBegin
	DirectiveEnd
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
	case DirectiveBegin:
		m["_type"] = "directive/begin"
	case DirectiveEnd:
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
			elem.Type = DirectiveBegin
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
			elem.Type = DirectiveEnd
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

func (doc *Document) Weave(eout io.Writer) error {
	return fmt.Errorf("Weave() not implemented")
}

func (doc *Document) Tangle(eout io.Writer) error {
	return fmt.Errorf("Tangle() not implemented")
}
