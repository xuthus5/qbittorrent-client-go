package qbittorrent

import (
	"net/url"
	"testing"
	"time"
)

var (
	c Client
)

func init() {
	var err error
	c, err = NewClient(&Config{
		Address:           "http://192.168.3.33:38080",
		Username:          "admin",
		Password:          "J0710cz5",
		RefreshIntervals:  time.Hour,
		ConnectionTimeout: time.Second * 3,
		CustomHeaders: map[string]string{
			//"Origin":     "http://192.168.3.33:8080",
			//"Referer":    "http://192.168.3.33:8080",
			"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36 Edg/125.0.0.0",
		},
	})
	if err != nil {
		panic(err)
	}
}

func TestFormEncoder(t *testing.T) {
	var option = LogOption{
		Normal:      true,
		Info:        true,
		Warning:     false,
		Critical:    false,
		LastKnownId: 0,
	}
	var form = url.Values{}
	err := encoder.Encode(option, form)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(form)
}
