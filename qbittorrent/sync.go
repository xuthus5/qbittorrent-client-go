package qbittorrent

import (
	"errors"
	"fmt"

	"github.com/bytedance/sonic"
)

type Sync interface {
	// MainData get sync main data, rid is Response ID. if not provided, will be assumed.
	// if the given is different from the one of last server reply, will be (see the server reply details for more info)
	MainData(rid int) (*SyncMainData, error)
	// TorrentPeersData get sync torrent peer data, hash is torrent hash, rid is response id
	TorrentPeersData(hash string, rid int) (*SyncTorrentPeers, error)
}

type SyncMainData struct {
	Rid         int                        `json:"rid,omitempty"`
	FullUpdate  bool                       `json:"full_update,omitempty"`
	ServerState ServerState                `json:"server_state,omitempty"`
	Torrents    map[string]SyncTorrentInfo `json:"torrents,omitempty"`
}

type ServerState struct {
	AllTimeDl          int64  `json:"alltime_dl,omitempty"`
	AllTimeUl          int64  `json:"alltime_ul,omitempty"`
	AverageTimeQueue   int    `json:"average_time_queue,omitempty"`
	DlInfoData         int64  `json:"dl_info_data,omitempty"`
	DlInfoSpeed        int    `json:"dl_info_speed,omitempty"`
	QueuedIoJobs       int    `json:"queued_io_jobs,omitempty"`
	TotalBuffersSize   int    `json:"total_buffers_size,omitempty"`
	UpInfoData         int64  `json:"up_info_data,omitempty"`
	UpInfoSpeed        int    `json:"up_info_speed,omitempty"`
	WriteCacheOverload string `json:"write_cache_overload,omitempty"`
}

type SyncTorrentInfo struct {
	AmountLeft        int64   `json:"amount_left,omitempty"`
	Completed         int     `json:"completed,omitempty"`
	DlSpeed           int     `json:"dlspeed,omitempty"`
	Downloaded        int     `json:"downloaded,omitempty"`
	DownloadedSession int     `json:"downloaded_session,omitempty"`
	Eta               int     `json:"eta,omitempty"`
	Progress          float64 `json:"progress,omitempty"`
	SeenComplete      int     `json:"seen_complete,omitempty"`
	TimeActive        int     `json:"time_active,omitempty"`
}

type SyncTorrentPeers struct {
	Rid        int                        `json:"rid,omitempty"`
	FullUpdate bool                       `json:"full_update,omitempty"`
	ShowFlags  bool                       `json:"show_flags,omitempty"`
	Peers      map[string]SyncTorrentPeer `json:"peers,omitempty"`
}

type SyncTorrentPeer struct {
	Client       string  `json:"client,omitempty"`
	Connection   string  `json:"connection,omitempty"`
	Country      string  `json:"country,omitempty"`
	CountryCode  string  `json:"country_code,omitempty"`
	DlSpeed      int     `json:"dl_speed,omitempty"`
	Downloaded   int     `json:"downloaded,omitempty"`
	Files        string  `json:"files,omitempty"`
	Flags        string  `json:"flags,omitempty"`
	FlagsDesc    string  `json:"flags_desc,omitempty"`
	IP           string  `json:"ip,omitempty"`
	PeerIDClient string  `json:"peer_id_client,omitempty"`
	Port         int     `json:"port,omitempty"`
	Progress     float64 `json:"progress,omitempty"`
	Relevance    float64 `json:"relevance,omitempty"`
	UpSpeed      int     `json:"up_speed,omitempty"`
	Uploaded     int     `json:"uploaded,omitempty"`
}

func (c *client) MainData(rid int) (*SyncMainData, error) {
	apiUrl := fmt.Sprintf("%s/api/v2/sync/maindata?rid=%d", c.config.Address, rid)
	result, err := c.doRequest(&requestData{
		url: apiUrl,
	})
	if err != nil {
		return nil, err
	}

	if result.code != 200 {
		return nil, errors.New("get main data failed: " + string(result.body))
	}

	var mainData = new(SyncMainData)
	if err := sonic.Unmarshal(result.body, mainData); err != nil {
		return nil, err
	}

	return mainData, nil
}

func (c *client) TorrentPeersData(hash string, rid int) (*SyncTorrentPeers, error) {
	apiUrl := fmt.Sprintf("%s/api/v2/sync/torrentPeers?hash=%s&rid=%d", c.config.Address, hash, rid)
	result, err := c.doRequest(&requestData{
		url: apiUrl,
	})
	if err != nil {
		return nil, err
	}

	if result.code != 200 {
		return nil, errors.New("get torrent peers data failed: " + string(result.body))
	}

	var mainData = new(SyncTorrentPeers)
	if err := sonic.Unmarshal(result.body, mainData); err != nil {
		return nil, err
	}

	return mainData, nil
}
