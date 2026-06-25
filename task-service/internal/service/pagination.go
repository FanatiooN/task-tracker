package service

import (
	"encoding/base64"
	"time"
)

func encodeToken(t time.Time) string {
	return base64.StdEncoding.EncodeToString([]byte(t.Format(time.RFC3339Nano)))
}
func decodeToken(token string) (*time.Time, error) {
	b, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return nil, err
	}

	t, err := time.Parse(time.RFC3339Nano, string(b))
	if err != nil {
		return nil, err
	}

	return &t, nil
}
