package untill

import (
	"encoding/xml"
	"net/http"
	"net/url"
)

func (c *Client) NewGetDepartmentsInfoRequest() GetDepartmentsInfoRequest {
	return GetDepartmentsInfoRequest{
		client:      c,
		queryParams: c.NewGetDepartmentsInfoQueryParams(),
		pathParams:  c.NewGetDepartmentsInfoPathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewGetDepartmentsInfoRequestBody(),
	}
}

type GetDepartmentsInfoRequest struct {
	client      *Client
	queryParams *GetDepartmentsInfoQueryParams
	pathParams  *GetDepartmentsInfoPathParams
	method      string
	headers     http.Header
	requestBody GetDepartmentsInfoRequestBody
}

func (p *GetDepartmentsInfoRequest) Action() string {
	return "urn:TPAPIPosIntfU-ITPAPIPOS#GetDepartmentsInfo"
}

func (c *Client) NewGetDepartmentsInfoQueryParams() *GetDepartmentsInfoQueryParams {
	return &GetDepartmentsInfoQueryParams{}
}

type GetDepartmentsInfoQueryParams struct {
}

func (p GetDepartmentsInfoQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetDepartmentsInfoRequest) QueryParams() *GetDepartmentsInfoQueryParams {
	return r.queryParams
}

func (c *Client) NewGetDepartmentsInfoPathParams() *GetDepartmentsInfoPathParams {
	return &GetDepartmentsInfoPathParams{}
}

type GetDepartmentsInfoPathParams struct{}

func (p *GetDepartmentsInfoPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetDepartmentsInfoRequest) PathParams() *GetDepartmentsInfoPathParams {
	return r.pathParams
}

func (r *GetDepartmentsInfoRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetDepartmentsInfoRequest) Method() string {
	return r.method
}

func (s *Client) NewGetDepartmentsInfoRequestBody() GetDepartmentsInfoRequestBody {
	return GetDepartmentsInfoRequestBody{}
}

type GetDepartmentsInfoRequestBody struct {
	XMLName xml.Name `xml:"GetDepartmentsInfo"`

	Credentials Credentials `xml:"Request"`
}

func (r *GetDepartmentsInfoRequest) RequestBody() *GetDepartmentsInfoRequestBody {
	return &r.requestBody
}

func (r *GetDepartmentsInfoRequest) SetRequestBody(body GetDepartmentsInfoRequestBody) {
	r.requestBody = body
}

func (r *GetDepartmentsInfoRequest) NewResponseBody() *GetDepartmentsInfoResponseBody {
	return &GetDepartmentsInfoResponseBody{}
}

type GetDepartmentsInfoResponseBody struct {
	XMLName xml.Name `xml:"GetDepartmentsInfoResponse"`

	ReturnCode    int               `xml:"return>ReturnCode"`
	ReturnMessage string            `xml:"return>ReturnMessage"`
	Departments   []DepartmentsInfo `xml:"return>Departments>item"`
}

func (r *GetDepartmentsInfoRequest) URL() *url.URL {
	return r.client.GetEndpointURL("", r.PathParams())
}

func (r *GetDepartmentsInfoRequest) Do() (GetDepartmentsInfoResponseBody, error) {
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

type DepartmentsInfo struct {
	DepartmentID     int    `xml:"DepartmentId"`
	DepartmentNumber int    `xml:"DepartmentNumber"`
	DepartmentName   string `xml:"DepartmentName"`
	Available        []int  `xml:"Available"`
	Supplement       int    `xml:"Supplement"`
	Condiment        int    `xml:"Condiment"`
	GroupID          int    `xml:"GroupId"`
	SpecialArticlese []int  `xml:"SpecialArticles"`
	HQID             string `xml:"HqId"`
	Extra            Extra  `xml:"Extra"`
}
