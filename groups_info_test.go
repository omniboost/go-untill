package untill_test

import (
	"log"
	"net/url"
	"os"
	"testing"

	_ "github.com/joho/godotenv/autoload"
	untill "github.com/omniboost/go-untill"
)

func TestGroupsInfo(t *testing.T) {
	client := untill.NewClient(nil, "", "")
	client.SetDebug(true)
	client.SetUsername(os.Getenv("UNTILL_USERNAME"))
	client.SetPassword(os.Getenv("UNTILL_PASSWORD"))
	client.SetBaseURL(url.URL{
		Scheme: "http",
		Host:   os.Getenv("UNTILL_HOST"),
		Path:   "/soap/ITPAPIPOS",
	})

	req := client.NewGetGroupsInfoRequest()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	log.Printf("%+v", resp)
}
