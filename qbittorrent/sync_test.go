package qbittorrent

import (
	"testing"
	"time"

	"github.com/bytedance/sonic"
)

func TestClient_MainData(t *testing.T) {
	syncMainData, err := c.Sync().MainData(0)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("sync main data: %+v", syncMainData)

	time.Sleep(time.Second)
	syncMainData, err = c.Sync().MainData(0)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("sync main data: %+v", syncMainData)
}

func TestClient_TorrentPeersData(t *testing.T) {
	peersData, err := c.Sync().TorrentPeersData("f23daefbe8d24d3dd882b44cb0b4f762bc23b4fc", 0)
	if err != nil {
		t.Fatal(err)
	}
	bytes, err := sonic.Marshal(peersData)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))
}
