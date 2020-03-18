package untill

import (
	"encoding/json"
	"time"
)

type DateTime struct {
	time.Time
}

func (d *DateTime) MarshalJSON() ([]byte, error) {
	if d.Time.IsZero() {
		return json.Marshal(nil)
	}

	return json.Marshal(d.Time.Format("2006-01-02T15:04:05-07:00"))
}

func (d *DateTime) UnmarshalJSON(text []byte) (err error) {
	var value string
	err = json.Unmarshal(text, &value)
	if err != nil {
		return err
	}

	if value == "" {
		return nil
	}

	// first try standard date
	d.Time, err = time.Parse(time.RFC3339, value)
	if err == nil {
		return nil
	}

	// try untill date format
	d.Time, err = time.Parse("2006-01-02T15:04:05-07:00", value)
	return
}

func (d DateTime) MarshalText() ([]byte, error) {
	return []byte(d.Format("2006-01-02T15:04:05-07:00")), nil
}
