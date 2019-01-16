package untill

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/shopspring/decimal"
)

func (c *Client) NewGetTurnoverReportRequest() GetTurnoverReportRequest {
	return GetTurnoverReportRequest{
		client:      c,
		queryParams: c.NewGetTurnoverReportQueryParams(),
		pathParams:  c.NewGetTurnoverReportPathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewGetTurnoverReportRequestBody(),
	}
}

type GetTurnoverReportRequest struct {
	client      *Client
	queryParams *GetTurnoverReportQueryParams
	pathParams  *GetTurnoverReportPathParams
	method      string
	headers     http.Header
	requestBody GetTurnoverReportRequestBody
}

func (p *GetTurnoverReportRequest) Action() string {
	return "urn:TPAPIPosIntfU-ITPAPIPOS#GetTurnoverReport"
}

func (c *Client) NewGetTurnoverReportQueryParams() *GetTurnoverReportQueryParams {
	return &GetTurnoverReportQueryParams{}
}

type GetTurnoverReportQueryParams struct {
}

func (p GetTurnoverReportQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetTurnoverReportRequest) QueryParams() *GetTurnoverReportQueryParams {
	return r.queryParams
}

func (c *Client) NewGetTurnoverReportPathParams() *GetTurnoverReportPathParams {
	return &GetTurnoverReportPathParams{}
}

type GetTurnoverReportPathParams struct{}

func (p *GetTurnoverReportPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetTurnoverReportRequest) PathParams() *GetTurnoverReportPathParams {
	return r.pathParams
}

func (r *GetTurnoverReportRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetTurnoverReportRequest) Method() string {
	return r.method
}

func (s *Client) NewGetTurnoverReportRequestBody() GetTurnoverReportRequestBody {
	return GetTurnoverReportRequestBody{}
}

type GetTurnoverReportRequestBody struct {
	XMLName xml.Name `xml:"GetTurnoverReport"`

	Username    string   `xml:"Request>UserName"`
	Password    string   `xml:"Request>Password"`
	From        DateTime `xml:"Request>From"`
	Till        DateTime `xml:"Request>Till"`
	SalesAreaID int      `xml:"Request>SalesAreaId"`
	Extra       []Extra  `xml":Request>Extra`
}

func (rb *GetTurnoverReportRequestBody) SetCredentials(creds Credentials) {
	rb.Username = creds.Username
	rb.Password = creds.Password
}

func (r *GetTurnoverReportRequest) RequestBody() *GetTurnoverReportRequestBody {
	return &r.requestBody
}

func (r *GetTurnoverReportRequest) SetRequestBody(body GetTurnoverReportRequestBody) {
	r.requestBody = body
}

func (r *GetTurnoverReportRequest) NewResponseBody() *GetTurnoverReportResponseBody {
	return &GetTurnoverReportResponseBody{}
}

type GetTurnoverReportResponseBody struct {
	XMLName xml.Name `xml:"GetTurnoverReportResponse"`

	ReturnCode    int           `xml:"return>ReturnCode"`
	ReturnMessage string        `xml:"return>ReturnMessage"`
	Bills         TurnoverBills `xml:"return>Data>item"`
	Extra         Extra         `xml:"return>Extra"`
}

func (r *GetTurnoverReportRequest) URL() *url.URL {
	return r.client.GetEndpointURL("", r.PathParams())
}

func (r *GetTurnoverReportRequest) Do() (GetTurnoverReportResponseBody, error) {
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

type TurnoverBills []TurnoverBill

type TurnoverBill struct {
	ID            int               `xml:"Id"`
	OpenDateTime  DateTime          `xml:"OpenDateTime"`
	CloseDateTime DateTime          `xml:"CloseDateTime"`
	ComputerName  string            `xml:"ComputerName"`
	BillNumber    string            `xml:"BillNumber"`
	BillSuffix    string            `xml:"BillSuffix"`
	TableNumber   int               `xml:"TableNumber"`
	TablePart     string            `xml:"TablePart"`
	Covers        decimal.Decimal   `xml:"Covers"`
	UserID        int               `xml:"UserId"`
	SalesAreadID  int               `xml:"SalesAreaId"`
	Items         TurnoverBillItems `xml:"Items>item"`
}

type TurnoverBillItems []TurnoverBillItem

type TurnoverBillItem struct {
	ArticleID    int             `xml:"ArticleId"`
	DateTime     DateTime        `xml:"DateTime"`
	ItemNumber   int             `xml:"ItemNumber"`
	ComputerName string          `xml:"ComputerName"`
	Quantity     int             `xml:"Quantity"`
	Price        decimal.Decimal `xml:"Price"`
	UserID       int             `xml:"UserId"`
	HQID         string          `xml:"HqId"`
	Extra        Extra           `xml:"Extra"`
}
