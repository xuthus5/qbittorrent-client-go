package qbittorrent

import (
	"testing"

	"github.com/bytedance/sonic"
)

func TestClient_GetTorrents(t *testing.T) {
	torrents, err := c.Torrent().GetTorrents(&TorrentOption{
		Filter:   "",
		Category: "movies",
		Tag:      "hdtime",
		Sort:     "",
		Reverse:  false,
		Limit:    0,
		Offset:   0,
		Hashes:   nil,
	})
	if err != nil {
		t.Fatal(err)
	}
	bytes, err := sonic.Marshal(torrents)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))
}

func TestClient_GetProperties(t *testing.T) {
	properties, err := c.Torrent().GetProperties("f23daefbe8d24d3dd882b44cb0b4f762bc23b4fc")
	if err != nil {
		t.Fatal(err)
	}
	bytes, err := sonic.Marshal(properties)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))
}

func TestClient_GetTrackers(t *testing.T) {
	trackers, err := c.Torrent().GetTrackers("f23daefbe8d24d3dd882b44cb0b4f762bc23b4fc")
	if err != nil {
		t.Fatal(err)
	}
	bytes, err := sonic.Marshal(trackers)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))
}

func TestClient_GetWebSeeds(t *testing.T) {
	webSeeds, err := c.Torrent().GetWebSeeds("f23daefbe8d24d3dd882b44cb0b4f762bc23b4fc")
	if err != nil {
		t.Fatal(err)
	}
	bytes, err := sonic.Marshal(webSeeds)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))
}

func TestClient_GetContents(t *testing.T) {
	contents, err := c.Torrent().GetContents("f23daefbe8d24d3dd882b44cb0b4f762bc23b4fc")
	if err != nil {
		t.Fatal(err)
	}
	bytes, err := sonic.Marshal(contents)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))
}

func TestClient_GetPiecesStates(t *testing.T) {
	states, err := c.Torrent().GetPiecesStates("f23daefbe8d24d3dd882b44cb0b4f762bc23b4fc")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(states)
}

func TestClient_GetPiecesHashes(t *testing.T) {
	hashes, err := c.Torrent().GetPiecesHashes("f23daefbe8d24d3dd882b44cb0b4f762bc23b4fc")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(hashes)
}

func TestClient_PauseTorrents(t *testing.T) {
	err := c.Torrent().PauseTorrents([]string{})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("torrent paused")
}
