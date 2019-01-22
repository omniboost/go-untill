package untill

import (
	"encoding/xml"
	"net/http"
	"net/url"
)

func (c *Client) NewGetClientsRequest() GetClientsRequest {
	return GetClientsRequest{
		client:      c,
		queryParams: c.NewGetClientsQueryParams(),
		pathParams:  c.NewGetClientsPathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewGetClientsRequestBody(),
	}
}

type GetClientsRequest struct {
	client      *Client
	queryParams *GetClientsQueryParams
	pathParams  *GetClientsPathParams
	method      string
	headers     http.Header
	requestBody GetClientsRequestBody
}

func (p *GetClientsRequest) Action() string {
	return "urn:TPAPIPosIntfU-ITPAPIPOS#GetClients"
}

func (c *Client) NewGetClientsQueryParams() *GetClientsQueryParams {
	return &GetClientsQueryParams{}
}

type GetClientsQueryParams struct {
}

func (p GetClientsQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetClientsRequest) QueryParams() *GetClientsQueryParams {
	return r.queryParams
}

func (c *Client) NewGetClientsPathParams() *GetClientsPathParams {
	return &GetClientsPathParams{}
}

type GetClientsPathParams struct{}

func (p *GetClientsPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetClientsRequest) PathParams() *GetClientsPathParams {
	return r.pathParams
}

func (r *GetClientsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetClientsRequest) Method() string {
	return r.method
}

func (s *Client) NewGetClientsRequestBody() GetClientsRequestBody {
	return GetClientsRequestBody{}
}

type GetClientsRequestBody struct {
	XMLName xml.Name `xml:"GetClients"`

	Username string `xml:"Request>UserName"`
	Password string `xml:"Request>Password"`
}

func (rb *GetClientsRequestBody) SetCredentials(creds Credentials) {
	rb.Username = creds.Username
	rb.Password = creds.Password
}

func (r *GetClientsRequest) RequestBody() *GetClientsRequestBody {
	return &r.requestBody
}

func (r *GetClientsRequest) SetRequestBody(body GetClientsRequestBody) {
	r.requestBody = body
}

func (r *GetClientsRequest) NewResponseBody() *GetClientsResponseBody {
	return &GetClientsResponseBody{}
}

type GetClientsResponseBody struct {
	XMLName xml.Name `xml:"GetClientsResponse"`

	ReturnCode    int    `xml:"return>ReturnCode"`
	ReturnMessage string `xml:"return>ReturnMessage"`
}

func (r *GetClientsRequest) URL() *url.URL {
	return r.client.GetEndpointURL("", r.PathParams())
}

func (r *GetClientsRequest) Do() (GetClientsResponseBody, error) {
	// Set credentials
	r.RequestBody().SetCredentials(r.client.Credentials())

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
