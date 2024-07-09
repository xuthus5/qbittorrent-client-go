package qbittorrent

import (
	"os"
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
	err := c.Torrent().PauseTorrents([]string{"202382999be6a4fab395cd9c2c9d294177587904"})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("torrent paused")
}

func TestClient_ResumeTorrents(t *testing.T) {
	err := c.Torrent().ResumeTorrents([]string{"fd3b4bf1937c59a8fd1a240cddc07172e0b979a2"})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("torrent resumed")
}

func TestClient_DeleteTorrents(t *testing.T) {
	err := c.Torrent().DeleteTorrents([]string{"202382999be6a4fab395cd9c2c9d294177587904"}, true)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("torrent deleted")
}

func TestClient_RecheckTorrents(t *testing.T) {
	err := c.Torrent().RecheckTorrents([]string{"fd3b4bf1937c59a8fd1a240cddc07172e0b979a2"})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("torrent rechecked")
}

func TestClient_ReAnnounceTorrents(t *testing.T) {
	err := c.Torrent().ReAnnounceTorrents([]string{"fd3b4bf1937c59a8fd1a240cddc07172e0b979a2"})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("torrent reannounceed")
}

func TestClient_AddNewTorrent(t *testing.T) {
	fileContent, err := os.ReadFile("C:\\Users\\xuthu\\Downloads\\bbbbb.torrent")
	if err != nil {
		t.Fatal(err)
	}
	err = c.Torrent().AddNewTorrent(&TorrentAddOption{
		Torrents: []*TorrentAddFileMetadata{
			{
				//Filename: "ttttt.torrent",
				Data: fileContent,
			},
		},
		Category:           "movies",
		Tags:               []string{"d", "e", "f"},
		SkipChecking:       false,
		Paused:             false,
		RootFolder:         false,
		Rename:             "",
		UpLimit:            0,
		DlLimit:            0,
		RatioLimit:         0,
		SeedingTimeLimit:   0,
		AutoTMM:            false,
		SequentialDownload: "",
		FirstLastPiecePrio: "",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("torrent added")
}

func TestClient_AddTrackers(t *testing.T) {
	err := c.Torrent().AddTrackers("ca4523a3db9c6c3a13d7d7f3a545f97b75083032", []string{"https://hddtime.org/announce"})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("torrent trackers added")
}

func TestClient_EditTrackers(t *testing.T) {
	err := c.Torrent().EditTrackers("ca4523a3db9c6c3a13d7d7f3a545f97b75083032", "https://hddtime.org/announce", "https://hdctime.org/announce")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("torrent trackers edited")
}

func TestClient_RemoveTrackers(t *testing.T) {
	err := c.Torrent().RemoveTrackers("ca4523a3db9c6c3a13d7d7f3a545f97b75083032", []string{"https://hdctime.org/announce"})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("torrent trackers removed")
}

func TestClient_AddPeers(t *testing.T) {
	// todo no test
	//c.Torrent().AddPeers([]string{"ca4523a3db9c6c3a13d7d7f3a545f97b75083032"}, []string{"10.0.0.1:38080"})
}

func TestClient_IncreasePriority(t *testing.T) {
	err := c.Torrent().IncreasePriority([]string{"916a250d32822adca39eb2b53efadfda1a15f902"})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("torrent priority increased")
}

func TestClient_DecreasePriority(t *testing.T) {
	err := c.Torrent().DecreasePriority([]string{"916a250d32822adca39eb2b53efadfda1a15f902"})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("torrent priority decreased")
}

func TestClient_MaxPriority(t *testing.T) {
	err := c.Torrent().MaxPriority([]string{"916a250d32822adca39eb2b53efadfda1a15f902"})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("torrent priority maxed")
}

func TestClient_MinPriority(t *testing.T) {
	err := c.Torrent().MinPriority([]string{"916a250d32822adca39eb2b53efadfda1a15f902"})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("torrent priority mined")
}

func TestClient_SetFilePriority(t *testing.T) {
	// todo no test
}

func TestClient_GetDownloadLimit(t *testing.T) {
	downloadLimit, err := c.Torrent().GetDownloadLimit([]string{"916a250d32822adca39eb2b53efadfda1a15f902"})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("torrent download limit", downloadLimit)
}

func TestClient_SetDownloadLimit(t *testing.T) {
	err := c.Torrent().SetDownloadLimit([]string{"916a250d32822adca39eb2b53efadfda1a15f902"}, 0)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("torrent download limit setted")
}

func TestClient_SetShareLimit(t *testing.T) {
	err := c.Torrent().SetShareLimit([]string{"916a250d32822adca39eb2b53efadfda1a15f902"}, -2, -2, -2)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("torrent share limit setted")
}

func TestClient_GetUploadLimit(t *testing.T) {
	limit, err := c.Torrent().GetUploadLimit([]string{"916a250d32822adca39eb2b53efadfda1a15f902"})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("torrent upload limit", limit)
}

func TestClient_SetUploadLimit(t *testing.T) {
	err := c.Torrent().SetUploadLimit([]string{"916a250d32822adca39eb2b53efadfda1a15f902"}, 0)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("torrent upload limit setted")
}

func TestClient_SetLocation(t *testing.T) {
	// todo test
}
