package main

import (
	"encoding/xml"
	"fmt"
)

// Go offers built-in support for XML and XML-like formats with
// the encoding/xml package.

// Plant will be mapped to XML. Here we use the special features of
// the XML package: the XMLName field dictates the name of the
// XML element representing this struct; id, attr means that if the
// field is an XML attribute rather than a nested element.
type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id int `xml:"id,attr"`
	Name string `xml:"name"`
	Origin []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v", p.Id, p.Name, p.Origin)
}

func main() {
	coffee := &Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	// Emit XML representing the plant, with MarshalIndent used to provide
	// a human readable output.
	out, _ := xml.MarshalIndent(coffee, " ", "  ")
	fmt.Println(string(out))

	// Explicit appending of an XML header.
	fmt.Println(xml.Header + string(out))

	// Unmarshal is used to parse a string of bytes with XML into a data
	// structure. If the XML is malformed or cannot be mapped onto Plant, a
	// descriptive error will be returned.
	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)

	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}

	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants []*Plant `xml:"parent>child>plant"`
	}

	// The parent>child>plant field tag tells the encoder to test
	// all plants under <parent><child>.
    nesting := &Nesting{}
    nesting.Plants = []*Plant{coffee, tomato}

    out, _ = xml.MarshalIndent(nesting, " ", "  ")
    fmt.Println(string(out))
}