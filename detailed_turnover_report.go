package untill

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/shopspring/decimal"
)

func (c *Client) NewGetDetailedTurnoverReportRequest() GetDetailedTurnoverReportRequest {
	return GetDetailedTurnoverReportRequest{
		client:      c,
		queryParams: c.NewGetDetailedTurnoverReportQueryParams(),
		pathParams:  c.NewGetDetailedTurnoverReportPathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewGetDetailedTurnoverReportRequestBody(),
	}
}

type GetDetailedTurnoverReportRequest struct {
	client      *Client
	queryParams *GetDetailedTurnoverReportQueryParams
	pathParams  *GetDetailedTurnoverReportPathParams
	method      string
	headers     http.Header
	requestBody GetDetailedTurnoverReportRequestBody
}

func (p *GetDetailedTurnoverReportRequest) Action() string {
	return "urn:TPAPIPosIntfU-ITPAPIPOS#GetDetailedTurnoverReport"
}

func (c *Client) NewGetDetailedTurnoverReportQueryParams() *GetDetailedTurnoverReportQueryParams {
	return &GetDetailedTurnoverReportQueryParams{}
}

type GetDetailedTurnoverReportQueryParams struct {
}

func (p GetDetailedTurnoverReportQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetDetailedTurnoverReportRequest) QueryParams() *GetDetailedTurnoverReportQueryParams {
	return r.queryParams
}

func (c *Client) NewGetDetailedTurnoverReportPathParams() *GetDetailedTurnoverReportPathParams {
	return &GetDetailedTurnoverReportPathParams{}
}

type GetDetailedTurnoverReportPathParams struct{}

func (p *GetDetailedTurnoverReportPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetDetailedTurnoverReportRequest) PathParams() *GetDetailedTurnoverReportPathParams {
	return r.pathParams
}

func (r *GetDetailedTurnoverReportRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetDetailedTurnoverReportRequest) Method() string {
	return r.method
}

func (s *Client) NewGetDetailedTurnoverReportRequestBody() GetDetailedTurnoverReportRequestBody {
	return GetDetailedTurnoverReportRequestBody{}
}

type GetDetailedTurnoverReportRequestBody struct {
	XMLName xml.Name `xml:"GetDetailedTurnoverReport"`

	Username    string   `xml:"Request>UserName"`
	Password    string   `xml:"Request>Password"`
	From        DateTime `xml:"Request>From"`
	Till        DateTime `xml:"Request>Till"`
	SalesAreaID int      `xml:"Request>SalesAreaId"`
	Extra       []Extra  `xml:"Request>Extra>item"`
}

func (rb *GetDetailedTurnoverReportRequestBody) SetCredentials(creds Credentials) {
	rb.Username = creds.Username
	rb.Password = creds.Password
}

func (r *GetDetailedTurnoverReportRequest) RequestBody() *GetDetailedTurnoverReportRequestBody {
	return &r.requestBody
}

func (r *GetDetailedTurnoverReportRequest) SetRequestBody(body GetDetailedTurnoverReportRequestBody) {
	r.requestBody = body
}

func (r *GetDetailedTurnoverReportRequest) NewResponseBody() *GetDetailedTurnoverReportResponseBody {
	return &GetDetailedTurnoverReportResponseBody{}
}

type GetDetailedTurnoverReportResponseBody struct {
	XMLName xml.Name `xml:"GetDetailedTurnoverReportResponse"`

	ReturnCode    int                  `xml:"return>ReturnCode"`
	ReturnMessage string               `xml:"return>ReturnMessage"`
	Transactions  TurnoverTransactions `xml:"return>Data>item"`
	Extra         []Extra              `xml:"return>Extra"`
}

func (r *GetDetailedTurnoverReportRequest) URL() *url.URL {
	return r.client.GetEndpointURL("", r.PathParams())
}

func (r *GetDetailedTurnoverReportRequest) Do() (GetDetailedTurnoverReportResponseBody, error) {
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

type TurnoverTransactions []TurnoverTransaction

type TurnoverTransaction struct {
	ID              int               `xml:"Id"`
	SalesArea       int               `xml:"SalesArea"`
	OpenDateTime    DateTime          `xml:"OpenDateTime"`
	CloseDateTime   DateTime          `xml:"CloseDateTime"`
	TranNumber      string            `xml:"TranNumber"`
	TableNumber     int               `xml:"TableNumber"`
	TablePart       string            `xml:"TablePart"`
	Covers          decimal.Decimal   `xml:"Covers"`
	UserID          int               `xml:"UserId"`
	DiscountOnTotal decimal.Decimal   `xml:"DiscountOnTotal"`
	ServiceCharge   decimal.Decimal   `xml:"ServiceCharge"`
	ClientName      string            `xml:"ClientName"`
	ClientID        int               `xml:"ClientId"`
	Orders          TransactionOrders `xml:"Orders>item"`
	Bills           TransactionBills  `xml:"Bills>item"`
}

type TransactionOrders []TransactionOrder

type TransactionOrder struct {
	ID           int                   `xml:"Id"`
	DateTime     DateTime              `xml:"DateTime"`
	ComputerName string                `xml:"ComputerName"`
	OrderNumber  string                `xml:"OrderNumber"`
	UserID       int                   `xml:"UserId"`
	SalesAreaID  int                   `xml:"SalesAreaId"`
	Items        TransactionOrderItems `xml:"Items>item"`
}

type TransactionOrderItems []TransactionOrderItem

type TransactionOrderItem struct {
	ID          int             `xml:"Id"`
	ArticleID   int             `xml:"ArticleId"`
	ItemNumber  int             `xml:"ItemNumber"`
	Kind        int             `xml:"Kind"`
	Quantity    int             `xml:"Quantity"`
	SinglePrice decimal.Decimal `xml:"SinglePrice"`
	Price       decimal.Decimal `xml:"Price"`
	Discount    decimal.Decimal `xml:"Discount"`
	Vat         decimal.Decimal `xml:"Vat"`
	VatPercent  int             `xml:"VatPercent"`
	Text        string          `xml:"Text"`
	HQID        string          `xml:"HqId"`
	Extra       Extra           `xml:"Extra>item"`
}

type TransactionBills []TransactionBill

type TransactionBill struct {
	ID           int             `xml:"Id"`
	DateTime     DateTime        `xml:"DateTime"`
	RealDateTime DateTime        `xml:"RealDateTime"`
	ComputerName string          `xml:"ComputerName"`
	BillNumber   string          `xml:"BillNumber"`
	UserID       int             `xml:"UserId"`
	SalesAreaID  int             `xml:"SalesAreaId"`
	Tip          decimal.Decimal `xml:"Tip"`
	Payments     Payments        `xml:"Payments>item"`
	ClientID     int             `xml:"ClientId"`
}

type Payments []Payment

type Payment struct {
	ID        int             `xml:"Id"`
	PaymentID int             `xml:"PaymentId"`
	Amount    decimal.Decimal `xml:"Amount"`
}
