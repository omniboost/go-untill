package untill_test

import (
	"log"
	"net/url"
	"os"
	"testing"
	"time"

	_ "github.com/joho/godotenv/autoload"
	untill "github.com/omniboost/go-untill"
)

func TestGetTurnoverReport(t *testing.T) {
	client := untill.NewClient(nil, "", "")
	client.SetDebug(true)
	client.SetUsername(os.Getenv("UNTILL_USERNAME"))
	client.SetPassword(os.Getenv("UNTILL_PASSWORD"))
	client.SetBaseURL(url.URL{
		Scheme: "http",
		Host:   os.Getenv("UNTILL_HOST"),
		Path:   "/soap/ITPAPIPOS",
	})

	// b, _ := ioutil.ReadFile("log.formatted.xml")
	// body := untill.GetTurnoverReportResponseBody{}
	// envelope := untill.Envelope{Body: &body}
	// errz := xml.Unmarshal(b, &envelope)
	// log.Println(errz)
	// log.Println(len(body.Transactions))
	// os.Exit(12)

	req := client.NewGetTurnoverReportRequest()
	from := time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)
	req.RequestBody().From = untill.DateTime{from}
	till := time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC)
	req.RequestBody().Till = untill.DateTime{till}
	req.RequestBody().SalesAreaID = 5000000259
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	// log.Println(len(resp.Transactions))
	// for _, tr := range resp.Transactions {
	// 	log.Println(tr.ID)
	// }
	log.Printf("%+v", resp)
}
