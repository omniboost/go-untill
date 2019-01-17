package untill

import (
	"encoding/xml"
	"net/http"
	"net/url"
)

func (c *Client) NewGetPaymentsInfoRequest() GetPaymentsInfoRequest {
	return GetPaymentsInfoRequest{
		client:      c,
		queryParams: c.NewGetPaymentsInfoQueryParams(),
		pathParams:  c.NewGetPaymentsInfoPathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewGetPaymentsInfoRequestBody(),
	}
}

type GetPaymentsInfoRequest struct {
	client      *Client
	queryParams *GetPaymentsInfoQueryParams
	pathParams  *GetPaymentsInfoPathParams
	method      string
	headers     http.Header
	requestBody GetPaymentsInfoRequestBody
}

func (p *GetPaymentsInfoRequest) Action() string {
	return "urn:TPAPIPosIntfU-ITPAPIPOS#GetPaymentsInfo"
}

func (c *Client) NewGetPaymentsInfoQueryParams() *GetPaymentsInfoQueryParams {
	return &GetPaymentsInfoQueryParams{}
}

type GetPaymentsInfoQueryParams struct {
}

func (p GetPaymentsInfoQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetPaymentsInfoRequest) QueryParams() *GetPaymentsInfoQueryParams {
	return r.queryParams
}

func (c *Client) NewGetPaymentsInfoPathParams() *GetPaymentsInfoPathParams {
	return &GetPaymentsInfoPathParams{}
}

type GetPaymentsInfoPathParams struct{}

func (p *GetPaymentsInfoPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetPaymentsInfoRequest) PathParams() *GetPaymentsInfoPathParams {
	return r.pathParams
}

func (r *GetPaymentsInfoRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetPaymentsInfoRequest) Method() string {
	return r.method
}

func (s *Client) NewGetPaymentsInfoRequestBody() GetPaymentsInfoRequestBody {
	return GetPaymentsInfoRequestBody{}
}

type GetPaymentsInfoRequestBody struct {
	XMLName xml.Name `xml:"GetPaymentsInfo"`

	Username string `xml:"Request>UserName"`
	Password string `xml:"Request>Password"`
}

func (rb *GetPaymentsInfoRequestBody) SetCredentials(creds Credentials) {
	rb.Username = creds.Username
	rb.Password = creds.Password
}

func (r *GetPaymentsInfoRequest) RequestBody() *GetPaymentsInfoRequestBody {
	return &r.requestBody
}

func (r *GetPaymentsInfoRequest) SetRequestBody(body GetPaymentsInfoRequestBody) {
	r.requestBody = body
}

func (r *GetPaymentsInfoRequest) NewResponseBody() *GetPaymentsInfoResponseBody {
	return &GetPaymentsInfoResponseBody{}
}

type GetPaymentsInfoResponseBody struct {
	XMLName xml.Name `xml:"GetPaymentsInfoResponse"`

	ReturnCode    int           `xml:"return>ReturnCode"`
	ReturnMessage string        `xml:"return>ReturnMessage"`
	Payments      []PaymentInfo `xml:"return>Payments>item"`
}

func (r *GetPaymentsInfoRequest) URL() *url.URL {
	return r.client.GetEndpointURL("", r.PathParams())
}

func (r *GetPaymentsInfoRequest) Do() (GetPaymentsInfoResponseBody, error) {
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

type PaymentInfo struct {
	PaymentID     int    `xml:"PaymentId"`
	PaymentNumber int    `xml:"PaymentNumber"`
	PaymentName   string `xml:"PaymentName"`
	PaymentKind   int    `xml:"PaymentKind"`
	Extra         Extra  `xml:"Extra"`
}
