package untill_test

import (
	"encoding/json"
	"log"
	"net/url"
	"os"
	"testing"

	_ "github.com/joho/godotenv/autoload"
	untill "github.com/omniboost/go-untill"
)

func TestGetDiscountReasons(t *testing.T) {
	client := untill.NewClient(nil, "", "")
	client.SetDebug(true)
	client.SetUsername(os.Getenv("UNTILL_USERNAME"))
	client.SetPassword(os.Getenv("UNTILL_PASSWORD"))
	client.SetBaseURL(url.URL{
		Scheme: "http",
		Host:   os.Getenv("UNTILL_HOST"),
		Path:   "/soap/ITPAPIPOS",
	})

	req := client.NewGetDiscountReasonsRequest()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp.DiscountReasons, "", "  ")
	log.Printf("%+v", (string(b)))
}
