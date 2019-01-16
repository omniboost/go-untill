package untill

import (
	"encoding/xml"
	"net/http"
	"strings"
)

type ErrorResponse struct {
	// HTTP response that caused this error
	Response *http.Response `xml:"-"`

	Errors []error
	FaultResponseBody
}

type FaultResponseBody struct {
	XMLName xml.Name `xml:"Fault"`

	FaultCode   string `xml:"faultcode"`
	FaultString string `xml:"faultstring"`
}

func (r ErrorResponse) Error() string {
	str := []string{}
	for _, err := range r.Errors {
		str = append(str, err.Error())
	}

	if r.FaultString != "" {
		str = append(str, r.FaultString)
	}

	return strings.Join(str, ", ")
}
