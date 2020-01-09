package untill

import "time"

type DateTime struct {
	time.Time
}

func (d DateTime) MarshalText() ([]byte, error) {
	return []byte(d.Format("2006-01-02T15:04:05")), nil
}
