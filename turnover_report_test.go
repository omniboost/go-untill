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
	from := time.Date(2019, 1, 22, 0, 0, 0, 0, time.UTC)
	req.RequestBody().From = untill.DateTime{from}
	till := time.Date(2019, 1, 23, 0, 0, 0, 0, time.UTC)
	req.RequestBody().Till = untill.DateTime{till}
	req.RequestBody().SalesAreaID = 5000000366
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp.Bills, "", "  ")
	log.Printf("%+v", (string(b)))
}
