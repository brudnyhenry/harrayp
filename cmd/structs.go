package cmd

import "encoding/xml"

// Response is returned by the array after every request.
type Response struct {
	XMLName xml.Name `xml:"RESPONSE"`
	Text    string   `xml:",chardata"`
	Link    struct {
		Text string `xml:",chardata"`
		Type string `xml:"type,attr"`
		ID   string `xml:"id,attr"`
		Rel  string `xml:"rel,attr"`
		Href string `xml:"href,attr"`
	} `xml:"link"`
	Style struct {
		Text string `xml:",chardata"`
		Type string `xml:"type,attr"`
		ID   string `xml:"id,attr"`
	} `xml:"style"`
	OBJECT []OBJECT `xml:"OBJECT"`
}

// OBJECT struct holds all returned objects (hosts/volumes etc.)
type OBJECT struct {
	XMLName  xml.Name   `xml:"OBJECT"`
	Basetype string     `xml:"basetype,attr"`
	Name     string     `xml:"name,attr"`
	Oid      string     `xml:"oid,attr"`
	Format   string     `xml:"format,attr"`
	PROPERTY []PROPERTY `xml:"PROPERTY"`
}

// PROPERTY holds attributes for a given object (volume size/ host name etc.)
type PROPERTY struct {
	XMLName     xml.Name `xml:"PROPERTY"`
	Text        string   `xml:",chardata"`
	Name        string   `xml:"name,attr"`
	Type        string   `xml:"type,attr"`
	Size        string   `xml:"size,attr"`
	Draw        string   `xml:"draw,attr"`
	Sort        string   `xml:"sort,attr"`
	DisplayName string   `xml:"display-name,attr"`
	Units       string   `xml:"units,attr"`
	Key         string   `xml:"key,attr"`
	Blocksize   string   `xml:"blocksize,attr"`
}
