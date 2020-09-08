package untill

import (
	"encoding/xml"
	"net/http"
	"net/url"
)

func (c *Client) NewGetDiscountReasonsRequest() GetDiscountReasonsRequest {
	return GetDiscountReasonsRequest{
		client:      c,
		queryParams: c.NewGetDiscountReasonsQueryParams(),
		pathParams:  c.NewGetDiscountReasonsPathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewGetDiscountReasonsRequestBody(),
	}
}

type GetDiscountReasonsRequest struct {
	client      *Client
	queryParams *GetDiscountReasonsQueryParams
	pathParams  *GetDiscountReasonsPathParams
	method      string
	headers     http.Header
	requestBody GetDiscountReasonsRequestBody
}

func (p *GetDiscountReasonsRequest) Action() string {
	return "urn:TPAPIPosIntfU-ITPAPIPOS#GetDiscountReasons"
}

func (c *Client) NewGetDiscountReasonsQueryParams() *GetDiscountReasonsQueryParams {
	return &GetDiscountReasonsQueryParams{}
}

type GetDiscountReasonsQueryParams struct {
}

func (p GetDiscountReasonsQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetDiscountReasonsRequest) QueryParams() *GetDiscountReasonsQueryParams {
	return r.queryParams
}

func (c *Client) NewGetDiscountReasonsPathParams() *GetDiscountReasonsPathParams {
	return &GetDiscountReasonsPathParams{}
}

type GetDiscountReasonsPathParams struct{}

func (p *GetDiscountReasonsPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetDiscountReasonsRequest) PathParams() *GetDiscountReasonsPathParams {
	return r.pathParams
}

func (r *GetDiscountReasonsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetDiscountReasonsRequest) Method() string {
	return r.method
}

func (s *Client) NewGetDiscountReasonsRequestBody() GetDiscountReasonsRequestBody {
	return GetDiscountReasonsRequestBody{}
}

type GetDiscountReasonsRequestBody struct {
	XMLName xml.Name `xml:"GetDiscountReasons"`

	Username string  `xml:"Request>UserName"`
	Password string  `xml:"Request>Password"`
	Extra    []Extra `xml:"Request>Extra>item"`
}

func (rb *GetDiscountReasonsRequestBody) SetCredentials(creds Credentials) {
	rb.Username = creds.Username
	rb.Password = creds.Password
}

func (r *GetDiscountReasonsRequest) RequestBody() *GetDiscountReasonsRequestBody {
	return &r.requestBody
}

func (r *GetDiscountReasonsRequest) SetRequestBody(body GetDiscountReasonsRequestBody) {
	r.requestBody = body
}

func (r *GetDiscountReasonsRequest) NewResponseBody() *GetDiscountReasonsResponseBody {
	return &GetDiscountReasonsResponseBody{}
}

type GetDiscountReasonsResponseBody struct {
	XMLName xml.Name `xml:"GetDiscountReasonsResponse"`

	ReturnCode      int             `xml:"return>ReturnCode"`
	ReturnMessage   string          `xml:"return>ReturnMessage"`
	Extra           Extra           `xml:"return>Extra"`
	DiscountReasons DiscountReasons `xml:"return>Items>item"`
}

func (r *GetDiscountReasonsRequest) URL() *url.URL {
	return r.client.GetEndpointURL("", r.PathParams())
}

func (r *GetDiscountReasonsRequest) Do() (GetDiscountReasonsResponseBody, error) {
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

type DiscountReasons []DiscountReason

type DiscountReason struct {
	ID     int    `xml:"Id"`
	Name   string `xml:"Name"`
	Number int    `xml:"Number"`
	Extra  Extra  `xml:"Extra"`
}
