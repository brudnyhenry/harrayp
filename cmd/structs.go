package cmd

import "encoding/xml"

// LoginResponse type returned by the array after sending login hash
type LoginResponse struct {
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
	OBJECT struct {
		Text     string `xml:",chardata"`
		Basetype string `xml:"basetype,attr"`
		Name     string `xml:"name,attr"`
		Oid      string `xml:"oid,attr"`
		PROPERTY []struct {
			Text string `xml:",chardata"`
			Name string `xml:"name,attr"`
		} `xml:"PROPERTY"`
	} `xml:"OBJECT"`
}

type HostsResponse struct {
	XMLName xml.Name `xml:"RESPONSE"`
	Text    string   `xml:",chardata"`
	VERSION string   `xml:"VERSION,attr"`
	OBJECT  []struct {
		Basetype string `xml:"basetype,attr"`
		Name     string `xml:"name,attr"`
		Oid      string `xml:"oid,attr"`
		Format   string `xml:"format,attr"`
		PROPERTY []struct {
			Text        string `xml:",chardata"`
			Name        string `xml:"name,attr"`
			Key         string `xml:"key,attr"`
			Type        string `xml:"type,attr"`
			Size        string `xml:"size,attr"`
			Draw        string `xml:"draw,attr"`
			Sort        string `xml:"sort,attr"`
			DisplayName string `xml:"display-name,attr"`
		} `xml:"PROPERTY"`
	} `xml:"OBJECT"`
}

type VolumesResponse struct {
	XMLName xml.Name `xml:"RESPONSE"`
	Text    string   `xml:",chardata"`
	VERSION string   `xml:"VERSION,attr"`
	OBJECT  []struct {
		Basetype string `xml:"basetype,attr"`
		Name     string `xml:"name,attr"`
		Oid      string `xml:"oid,attr"`
		Format   string `xml:"format,attr"`
		PROPERTY []struct {
			Text        string `xml:",chardata"`
			Name        string `xml:"name,attr"`
			Type        string `xml:"type,attr"`
			Size        string `xml:"size,attr"`
			Draw        string `xml:"draw,attr"`
			Sort        string `xml:"sort,attr"`
			DisplayName string `xml:"display-name,attr"`
			Units       string `xml:"units,attr"`
			Key         string `xml:"key,attr"`
			Blocksize   string `xml:"blocksize,attr"`
		} `xml:"PROPERTY"`
	} `xml:"OBJECT"`
}
