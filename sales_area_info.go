package untill

import (
	"encoding/xml"
	"net/http"
	"net/url"
)

func (c *Client) NewGetSalesAreasInfoRequest() GetSalesAreasInfoRequest {
	return GetSalesAreasInfoRequest{
		client:      c,
		queryParams: c.NewGetSalesAreasInfoQueryParams(),
		pathParams:  c.NewGetSalesAreasInfoPathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewGetSalesAreasInfoRequestBody(),
	}
}

type GetSalesAreasInfoRequest struct {
	client      *Client
	queryParams *GetSalesAreasInfoQueryParams
	pathParams  *GetSalesAreasInfoPathParams
	method      string
	headers     http.Header
	requestBody GetSalesAreasInfoRequestBody
}

func (p *GetSalesAreasInfoRequest) Action() string {
	return "urn:TPAPIPosIntfU-ITPAPIPOS#GetSalesAreasInfo"
}

func (c *Client) NewGetSalesAreasInfoQueryParams() *GetSalesAreasInfoQueryParams {
	return &GetSalesAreasInfoQueryParams{}
}

type GetSalesAreasInfoQueryParams struct {
}

func (p GetSalesAreasInfoQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetSalesAreasInfoRequest) QueryParams() *GetSalesAreasInfoQueryParams {
	return r.queryParams
}

func (c *Client) NewGetSalesAreasInfoPathParams() *GetSalesAreasInfoPathParams {
	return &GetSalesAreasInfoPathParams{}
}

type GetSalesAreasInfoPathParams struct{}

func (p *GetSalesAreasInfoPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetSalesAreasInfoRequest) PathParams() *GetSalesAreasInfoPathParams {
	return r.pathParams
}

func (r *GetSalesAreasInfoRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetSalesAreasInfoRequest) Method() string {
	return r.method
}

func (s *Client) NewGetSalesAreasInfoRequestBody() GetSalesAreasInfoRequestBody {
	return GetSalesAreasInfoRequestBody{}
}

type GetSalesAreasInfoRequestBody struct {
	XMLName xml.Name `xml:"GetSalesAreasInfo"`

	Credentials Credentials `xml:"Request"`
}

func (r *GetSalesAreasInfoRequest) RequestBody() *GetSalesAreasInfoRequestBody {
	return &r.requestBody
}

func (r *GetSalesAreasInfoRequest) SetRequestBody(body GetSalesAreasInfoRequestBody) {
	r.requestBody = body
}

func (r *GetSalesAreasInfoRequest) NewResponseBody() *GetSalesAreasInfoResponseBody {
	return &GetSalesAreasInfoResponseBody{}
}

type GetSalesAreasInfoResponseBody struct {
	XMLName xml.Name `xml:"GetSalesAreasInfoResponse"`

	ReturnCode    int        `xml:"return>ReturnCode"`
	ReturnMessage string     `xml:"return>ReturnMessage"`
	SalesAreas    SalesAreas `xml:"return>SalesAreas>item"`
}

func (r *GetSalesAreasInfoRequest) URL() *url.URL {
	return r.client.GetEndpointURL("", r.PathParams())
}

func (r *GetSalesAreasInfoRequest) Do() (GetSalesAreasInfoResponseBody, error) {
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

type SalesAreas []SalesArea

type SalesArea struct {
	ID      int    `xml:"SalesAreaId"`
	Number  int    `xml:"SalesAreaNumber"`
	Name    string `xml:"SalesAreaName"`
	PriceID int    `xml:"PriceId"`
	Tables  Tables `xml:"Tables>item"`
	Extra   Extra  `xml:"Extra"`
}

type Tables []Table

type Table struct {
	FromTable int   `xml:"FromTable"`
	ToTable   int   `xml:"ToTable"`
	Extra     Extra `xml:"Extra>item"`
}

// <Extra xsi:type="SOAP-ENC:Array" SOAP-ENC:arrayType="NS3:TExtraInfo[2]">
//   <item xsi:type="NS3:TExtraInfo">
// 	<Key xsi:type="xsd:string">bookkeeping_turnover</Key>
// 	<Value xsi:type="xsd:string">800070</Value>
// 	<Extra xsi:type="SOAP-ENC:Array" SOAP-ENC:arrayType="NS3:TExtraInfo[0]"/>
//   </item>
//   <item xsi:type="NS3:TExtraInfo">
// 	<Key xsi:type="xsd:string">bookkeeping_vat</Key>
// 	<Value xsi:type="xsd:string">190041</Value>
// 	<Extra xsi:type="SOAP-ENC:Array" SOAP-ENC:arrayType="NS3:TExtraInfo[0]"/>
//   </item>
// </Extra>

type Extra []ExtraItem

type ExtraItem struct {
	Key   string `xml:"Key"`
	Value string `xml:"Value"`
}
