package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

type node interface{}
type charData string
type element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []node
}

var indent = 0

func format(out *os.File, n node) {
	switch n := n.(type) {
	case charData:
		fmt.Fprintf(out, "%*s%s\n", indent*2, "", string(n))
	case element:
		fmt.Fprintf(out, "%*s<%s", indent*2, "", n.Type.Local)
		indent++
		for _, a := range n.Attr {
			fmt.Fprintf(out, ` %s="%s"`, a.Name.Local, a.Value)
		}
		fmt.Fprintf(out, ">\n")
		for _, c := range n.Children {
			format(out, c)
		}
		indent--
		fmt.Fprintf(out, "%*s</%s>\n", indent*2, "", n.Type.Local)
	}
}

func xmlNodes(dec *xml.Decoder) node {
	tok, err := dec.Token()
	if err == io.EOF {
		fmt.Println(err)
		return nil
	} else if err != nil {
		fmt.Fprintf(os.Stderr, "xmlNodes: %v\n", err)
		os.Exit(1)
	}

	// deal different types
	switch tok := tok.(type) {
	case xml.StartElement:
		var e element
		e.Type = tok.Name
		e.Attr = tok.Attr
		var children []node
		for {
			c := xmlNodes(dec)
			if c != nil {
				children = append(children, c)
			} else {
				break
			}
		}
		e.Children = children
		return e
	case xml.CharData:
		c := string(tok.Copy())
		// I'm not sure should I use trim here
		// because the xml.Decoder cut some empty lines as CharData
		c = strings.Trim(c, " \n\r\t")
		if len(c) == 0 {
			return xmlNodes(dec)
		}
		return charData(c)
	case xml.ProcInst:
		return xmlNodes(dec)
	default:
	}
	return nil
}
func main() {
	dec := xml.NewDecoder(os.Stdin)
	n := xmlNodes(dec)
	format(os.Stdout, n)
}
