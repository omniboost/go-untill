package untill

import (
	"encoding/xml"
)

type Envelope struct {
	Namespaces []xml.Attr `xml:"-"`

	Header Header `xml:"-"`
	Body   Body   `xml:"-"`
}

func (env Envelope) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name = xml.Name{Local: "soapenv:Envelope"}

	type headerWrapper struct {
		Header
	}

	type bodyWrapper struct {
		Body
	}

	type wrapper struct {
		Header headerWrapper `xml:"soapenv:Header"`
		Body   bodyWrapper   `xml:"soapenv:Body"`
	}

	for _, ns := range env.Namespaces {
		start.Attr = append(start.Attr, ns)
	}

	w := wrapper{
		Header: headerWrapper{env.Header},
		Body:   bodyWrapper{env.Body},
	}

	return e.EncodeElement(w, start)
}

func (env *Envelope) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type headerWrapper struct {
		Header `xml:",any"`
	}

	type bodyWrapper struct {
		Body `xml:",any"`
	}

	type wrapper struct {
		Header headerWrapper `xml:"Header"`
		Body   bodyWrapper   `xml:"Body"`
	}

	w := wrapper{
		Header: headerWrapper{env.Header},
		Body:   bodyWrapper{env.Body},
	}

	err := d.DecodeElement(&w, &start)
	if err != nil {
		return err
	}

	env.Header = w.Header.Header
	env.Body = w.Body.Body
	return nil
}

type Body interface{}

type Header interface{}
