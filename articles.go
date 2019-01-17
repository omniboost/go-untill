package untill

import (
	"encoding/xml"
	"net/http"
	"net/url"
)

func (c *Client) NewGetArticlesRequest() GetArticlesRequest {
	return GetArticlesRequest{
		client:      c,
		queryParams: c.NewGetArticlesQueryParams(),
		pathParams:  c.NewGetArticlesPathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewGetArticlesRequestBody(),
	}
}

type GetArticlesRequest struct {
	client      *Client
	queryParams *GetArticlesQueryParams
	pathParams  *GetArticlesPathParams
	method      string
	headers     http.Header
	requestBody GetArticlesRequestBody
}

func (p *GetArticlesRequest) Action() string {
	return "urn:TPAPIPosIntfU-ITPAPIPOS#GetArticles"
}

func (c *Client) NewGetArticlesQueryParams() *GetArticlesQueryParams {
	return &GetArticlesQueryParams{}
}

type GetArticlesQueryParams struct {
}

func (p GetArticlesQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetArticlesRequest) QueryParams() *GetArticlesQueryParams {
	return r.queryParams
}

func (c *Client) NewGetArticlesPathParams() *GetArticlesPathParams {
	return &GetArticlesPathParams{}
}

type GetArticlesPathParams struct{}

func (p *GetArticlesPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetArticlesRequest) PathParams() *GetArticlesPathParams {
	return r.pathParams
}

func (r *GetArticlesRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetArticlesRequest) Method() string {
	return r.method
}

func (s *Client) NewGetArticlesRequestBody() GetArticlesRequestBody {
	return GetArticlesRequestBody{}
}

type GetArticlesRequestBody struct {
	XMLName xml.Name `xml:"GetArticles"`

	Credentials Credentials `xml:"Request"`
}

func (r *GetArticlesRequest) RequestBody() *GetArticlesRequestBody {
	return &r.requestBody
}

func (r *GetArticlesRequest) SetRequestBody(body GetArticlesRequestBody) {
	r.requestBody = body
}

func (r *GetArticlesRequest) NewResponseBody() *GetArticlesResponseBody {
	return &GetArticlesResponseBody{}
}

type GetArticlesResponseBody struct {
	XMLName xml.Name `xml:"GetArticlesResponse"`

	ReturnCode    int           `xml:"return>ReturnCode"`
	ReturnMessage string        `xml:"return>ReturnMessage"`
	Articles      ArticlesShort `xml:"return>Articles>item"`
}

func (r *GetArticlesRequest) URL() *url.URL {
	return r.client.GetEndpointURL("", r.PathParams())
}

func (r *GetArticlesRequest) Do() (GetArticlesResponseBody, error) {
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

type ArticlesShort []ArticleShort

type ArticleShort struct {
	ArticleID     int    `xml:"ArticleId"`
	ArticleName   string `xml:"ArticleName"`
	ArticleNumber string `xml:"ArticleNumber"`
	SalesAreaID   int    `xml:"SalesAreaId"`
	DepartmentID  int    `xml:"DepartmentId"`
	HQID          string `xml:"HqId"`
	Extra         Extra  `xml:"Extra"`
}
