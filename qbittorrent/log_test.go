package qbittorrent

import (
	"testing"

	"github.com/bytedance/sonic"
)

func TestClient_GetLog(t *testing.T) {
	entries, err := c.Log().GetLog(&LogOption{
		Normal:      true,
		Info:        true,
		Warning:     true,
		Critical:    true,
		LastKnownId: 0,
	})
	if err != nil {
		t.Fatal(err)
	}
	bytes, err := sonic.Marshal(entries)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))
}

func TestClient_GetPeerLog(t *testing.T) {
	entries, err := c.Log().GetPeerLog(-1)
	if err != nil {
		t.Fatal(err)
	}
	bytes, err := sonic.Marshal(entries)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))
}
