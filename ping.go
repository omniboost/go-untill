package untill

import (
	"encoding/xml"
	"net/http"
	"net/url"
)

func (c *Client) NewPingRequest() PingRequest {
	return PingRequest{
		client:      c,
		queryParams: c.NewPingQueryParams(),
		pathParams:  c.NewPingPathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewPingRequestBody(),
	}
}

type PingRequest struct {
	client      *Client
	queryParams *PingQueryParams
	pathParams  *PingPathParams
	method      string
	headers     http.Header
	requestBody PingRequestBody
}

func (p *PingRequest) Action() string {
	return "urn:TPAPIPosIntfU-ITPAPIPOS#Ping"
}

func (c *Client) NewPingQueryParams() *PingQueryParams {
	return &PingQueryParams{}
}

type PingQueryParams struct {
}

func (p PingQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *PingRequest) QueryParams() *PingQueryParams {
	return r.queryParams
}

func (c *Client) NewPingPathParams() *PingPathParams {
	return &PingPathParams{}
}

type PingPathParams struct{}

func (p *PingPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *PingRequest) PathParams() *PingPathParams {
	return r.pathParams
}

func (r *PingRequest) SetMethod(method string) {
	r.method = method
}

func (r *PingRequest) Method() string {
	return r.method
}

func (s *Client) NewPingRequestBody() PingRequestBody {
	return PingRequestBody{}
}

type PingRequestBody struct {
	XMLName xml.Name `xml:"Ping"`

	Credentials Credentials `xml:"Request"`
}

func (r *PingRequest) RequestBody() *PingRequestBody {
	return &r.requestBody
}

func (r *PingRequest) SetRequestBody(body PingRequestBody) {
	r.requestBody = body
}

func (r *PingRequest) NewResponseBody() *PingResponseBody {
	return &PingResponseBody{}
}

type PingResponseBody struct {
	XMLName xml.Name `xml:"PingResponse"`

	ReturnCode    int    `xml:"return>ReturnCode"`
	ReturnMessage string `xml:"return>ReturnMessage"`
}

func (r *PingRequest) URL() *url.URL {
	return r.client.GetEndpointURL("", r.PathParams())
}

func (r *PingRequest) Do() (PingResponseBody, error) {
	// Set credentials
	r.RequestBody().Credentials = r.client.Credentials()

	// Create http request
	req, err := r.client.NewRequest(nil, r, r.RequestBody())
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
