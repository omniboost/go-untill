package untill_test

import (
	"encoding/json"
	"log"
	"net/url"
	"os"
	"testing"
	"time"

	_ "github.com/joho/godotenv/autoload"
	untill "github.com/omniboost/go-untill"
)

func TestGetDetailedTurnoverReport(t *testing.T) {
	client := untill.NewClient(nil, "", "")
	client.SetDebug(true)
	client.SetUsername(os.Getenv("UNTILL_USERNAME"))
	client.SetPassword(os.Getenv("UNTILL_PASSWORD"))
	client.SetBaseURL(url.URL{
		Scheme: "http",
		Host:   os.Getenv("UNTILL_HOST"),
		Path:   "/soap/ITPAPIPOS",
	})

	salesAreasReq := client.NewGetSalesAreasInfoRequest()
	resp, err := salesAreasReq.Do()
	if err != nil {
		t.Error(err)
	}

	for _, sa := range resp.SalesAreas {
		req := client.NewGetDetailedTurnoverReportRequest()
		from := time.Date(2019, 4, 6, 0, 0, 0, 0, time.UTC)
		req.RequestBody().From = untill.DateTime{from}
		till := time.Date(2019, 4, 7, 0, 0, 0, 0, time.UTC)
		req.RequestBody().Till = untill.DateTime{till}
		req.RequestBody().SalesAreaID = sa.ID
		resp, err := req.Do()
		if err != nil {
			t.Error(err)
		}

		b, _ := json.MarshalIndent(resp.Transactions, "", "  ")
		log.Printf("%+v", (string(b)))
	}
}
