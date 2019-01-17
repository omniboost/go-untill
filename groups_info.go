package untill

import (
	"encoding/xml"
	"net/http"
	"net/url"
)

func (c *Client) NewGetGroupsInfoRequest() GetGroupsInfoRequest {
	return GetGroupsInfoRequest{
		client:      c,
		queryParams: c.NewGetGroupsInfoQueryParams(),
		pathParams:  c.NewGetGroupsInfoPathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewGetGroupsInfoRequestBody(),
	}
}

type GetGroupsInfoRequest struct {
	client      *Client
	queryParams *GetGroupsInfoQueryParams
	pathParams  *GetGroupsInfoPathParams
	method      string
	headers     http.Header
	requestBody GetGroupsInfoRequestBody
}

func (p *GetGroupsInfoRequest) Action() string {
	return "urn:TPAPIPosIntfU-ITPAPIPOS#GetGroupsInfo"
}

func (c *Client) NewGetGroupsInfoQueryParams() *GetGroupsInfoQueryParams {
	return &GetGroupsInfoQueryParams{}
}

type GetGroupsInfoQueryParams struct {
}

func (p GetGroupsInfoQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetGroupsInfoRequest) QueryParams() *GetGroupsInfoQueryParams {
	return r.queryParams
}

func (c *Client) NewGetGroupsInfoPathParams() *GetGroupsInfoPathParams {
	return &GetGroupsInfoPathParams{}
}

type GetGroupsInfoPathParams struct{}

func (p *GetGroupsInfoPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetGroupsInfoRequest) PathParams() *GetGroupsInfoPathParams {
	return r.pathParams
}

func (r *GetGroupsInfoRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetGroupsInfoRequest) Method() string {
	return r.method
}

func (s *Client) NewGetGroupsInfoRequestBody() GetGroupsInfoRequestBody {
	return GetGroupsInfoRequestBody{}
}

type GetGroupsInfoRequestBody struct {
	XMLName xml.Name `xml:"GetGroupsInfo"`

	Credentials Credentials `xml:"Request"`
}

func (r *GetGroupsInfoRequest) RequestBody() *GetGroupsInfoRequestBody {
	return &r.requestBody
}

func (r *GetGroupsInfoRequest) SetRequestBody(body GetGroupsInfoRequestBody) {
	r.requestBody = body
}

func (r *GetGroupsInfoRequest) NewResponseBody() *GetGroupsInfoResponseBody {
	return &GetGroupsInfoResponseBody{}
}

type GetGroupsInfoResponseBody struct {
	XMLName xml.Name `xml:"GetGroupsInfoResponse"`

	ReturnCode    int         `xml:"return>ReturnCode"`
	ReturnMessage string      `xml:"return>ReturnMessage"`
	Groups        []GroupInfo `xml:"return>Groups>item"`
}

func (r *GetGroupsInfoRequest) URL() *url.URL {
	return r.client.GetEndpointURL("", r.PathParams())
}

func (r *GetGroupsInfoRequest) Do() (GetGroupsInfoResponseBody, error) {
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

type GroupInfo struct {
	GroupID    int    `xml:"GroupId"`
	GroupName  string `xml:"GroupName"`
	CategoryID int    `xml:"CategoryId"`
	HQID       string `xml:"HqId"`
	Extra      Extra  `xml:"Extra>item"`

	// <GroupId xsi:type="xsd:long">5000000048</GroupId>
	// <GroupName xsi:type="xsd:string">Dranken Hoog</GroupName>
	// <CategoryId xsi:type="xsd:long">5000000044</CategoryId>
	// <HqId xsi:type="xsd:string">Dranken Hoog</HqId>
	//   </item>
	// </Extra>
}
