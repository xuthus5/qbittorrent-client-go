package qbittorrent

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/bytedance/sonic"
)

type Torrent interface {
	// GetTorrents get torrent list
	GetTorrents(opt *TorrentOption) ([]*TorrentInfo, error)
	// GetProperties get torrent generic properties
	GetProperties(hash string) (*TorrentProperties, error)
	// GetTrackers get torrent trackers
	GetTrackers(hash string) ([]*TorrentTracker, error)
	// GetWebSeeds get torrent web seeds
	GetWebSeeds(hash string) ([]*TorrentWebSeed, error)
	// GetContents get torrent contents, indexes(optional) of the files you want to retrieve
	GetContents(hash string, indexes ...string) ([]*TorrentContent, error)
	// GetPiecesStates get torrent pieces states
	GetPiecesStates(hash string) ([]int, error)
	// GetPiecesHashes get torrent pieces hashes
	GetPiecesHashes(hash string) ([]string, error)
	// PauseTorrents the hashes of the torrents you want to pause
	PauseTorrents(hashes []string) error
	// ResumeTorrents the hashes of the torrents you want to resume
	ResumeTorrents(hashes []string) error
	// DeleteTorrents the hashes of the torrents you want to delete, if set deleteFile to true,
	// the downloaded data will also be deleted, otherwise has no effect.
	DeleteTorrents(hashes []string, deleteFile bool) error
	// RecheckTorrents the hashes of the torrents you want to recheck
	RecheckTorrents(hashes []string) error
	// ReAnnounceTorrents the hashes of the torrents you want to reannounce
	ReAnnounceTorrents(hashes []string) error
	// AddNewTorrent add torrents from server local file or from URLs. http://, https://, magnet: and bc://bt/ links are supported
	AddNewTorrent(opt *TorrentAddOption) error
	// AddTrackers add trackers to torrent
	AddTrackers(hash string, urls []string) error
	// EditTrackers edit trackers
	EditTrackers(hash, origUrl, newUrl string) error
	// RemoveTrackers remove trackers
	RemoveTrackers(hash string, urls []string) error
	// AddPeers add peers for torrent, each peer is host:port
	AddPeers(hashes []string, peers []string) error
	// IncreasePriority increase torrent priority
	IncreasePriority(hashes []string) error
	// DecreasePriority decrease torrent priority
	DecreasePriority(hashes []string) error
	// MaxPriority maximal torrent priority
	MaxPriority(hashes []string) error
	// MinPriority minimal torrent priority
	MinPriority(hashes []string) error
	// SetFilePriority set file priority
	SetFilePriority(hash string, id string, priority int) error
	// GetDownloadLimit get torrent download limit
	GetDownloadLimit(hashes []string) (map[string]int, error)
	// SetDownloadLimit set torrent download limit, limit in bytes per second
	SetDownloadLimit(hashes []string, limit int) error
	// SetShareLimit set torrent share limit
	SetShareLimit(hashes []string, ratioLimit float64, seedingTimeLimit, inactiveSeedingTimeLimit int) error
	// GetUploadLimit get torrent upload limit
	GetUploadLimit(hashes []string) (map[string]int, error)
	// SetUploadLimit set torrent upload limit
	SetUploadLimit(hashes []string, limit int) error
	// SetLocation set torrent location
	SetLocation(hashes []string, location string) error
	// SetName set torrent name
	SetName(hash string, name string) error
	// SetCategory set torrent category
	SetCategory(hashes []string, category string) error
	// GetCategories get all categories
	GetCategories() (map[string]*TorrentCategory, error)
	// AddNewCategory add new category
	AddNewCategory(category, savePath string) error
	// EditCategory edit category
	EditCategory(category, savePath string) error
	// RemoveCategories remove categories
	RemoveCategories(categories []string) error
	// AddTags add torrent tags
	AddTags(hashes []string, tags []string) error
	// RemoveTags remove torrent tags
	RemoveTags(hashes []string, tags []string) error
	// GetTags get all tags
	GetTags() ([]string, error)
	// CreateTags create tags
	CreateTags(tags []string) error
	// DeleteTags delete tags
	DeleteTags(tags []string) error
	// SetAutomaticManagement set automatic torrent management
	SetAutomaticManagement(hashes []string, enable bool) error
	// ToggleSequentialDownload toggle sequential download
	ToggleSequentialDownload(hashes []string) error
	// SetFirstLastPiecePriority set first/last piece priority
	SetFirstLastPiecePriority(hashes []string) error
	// SetForceStart set force start
	SetForceStart(hashes []string, force bool) error
	// SetSuperSeeding set super seeding
	SetSuperSeeding(hashes []string, enable bool) error
	// RenameFile rename file
	RenameFile(hash, oldPath, newPath string) error
	// RenameFolder rename folder
	RenameFolder(hash, oldPath, newPath string) error
}

type TorrentOption struct {
	// Filter torrent list by state. Allowed state filters: all,downloading,seeding,completed,paused,
	// active,inactive,resumed,stalled,stalled_uploading,stalled_downloading,errored
	Filter string `schema:"filter,omitempty"`
	// Category get torrents with the given category, empty string means "without category"; no "category"
	// parameter means "any category"
	Category string `schema:"category,omitempty"`
	// Tag get torrents with the given tag, empty string means "without tag"; no "tag" parameter means "any tag"
	Tag string `schema:"tag,omitempty"`
	// Sort  torrents by given key, they can be sorted using any field of the response's JSON array (which are documented below) as the sort key.
	Sort string `schema:"sort,omitempty"`
	// Reverse enable reverse sorting. Defaults to false
	Reverse bool `schema:"reverse,omitempty"`
	// Limit the number of torrents returned
	Limit int `schema:"limit,omitempty"`
	// Offset set offset (if less than 0, offset from end)
	Offset int `schema:"offset,omitempty"`
	// Hashes filter by hashes
	Hashes []string `schema:"-"`
}

type TorrentInfo struct {
	AddedOn                  int     `json:"added_on"`
	AmountLeft               int     `json:"amount_left"`
	AutoTmm                  bool    `json:"auto_tmm"`
	Availability             int     `json:"availability"`
	Category                 string  `json:"category"`
	Completed                int     `json:"completed"`
	CompletionOn             int     `json:"completion_on"`
	ContentPath              string  `json:"content_path"`
	DlLimit                  int     `json:"dl_limit"`
	Dlspeed                  int     `json:"dlspeed"`
	DownloadPath             string  `json:"download_path"`
	Downloaded               int     `json:"downloaded"`
	DownloadedSession        int     `json:"downloaded_session"`
	Eta                      int     `json:"eta"`
	FLPiecePrio              bool    `json:"f_l_piece_prio"`
	ForceStart               bool    `json:"force_start"`
	Hash                     string  `json:"hash"`
	InactiveSeedingTimeLimit int     `json:"inactive_seeding_time_limit"`
	InfohashV1               string  `json:"infohash_v1"`
	InfohashV2               string  `json:"infohash_v2"`
	LastActivity             int     `json:"last_activity"`
	MagnetURI                string  `json:"magnet_uri"`
	MaxInactiveSeedingTime   int     `json:"max_inactive_seeding_time"`
	MaxRatio                 int     `json:"max_ratio"`
	MaxSeedingTime           int     `json:"max_seeding_time"`
	Name                     string  `json:"name"`
	NumComplete              int     `json:"num_complete"`
	NumIncomplete            int     `json:"num_incomplete"`
	NumLeechs                int     `json:"num_leechs"`
	NumSeeds                 int     `json:"num_seeds"`
	Priority                 int     `json:"priority"`
	Progress                 int     `json:"progress"`
	Ratio                    float64 `json:"ratio"`
	RatioLimit               int     `json:"ratio_limit"`
	SavePath                 string  `json:"save_path"`
	SeedingTime              int     `json:"seeding_time"`
	SeedingTimeLimit         int     `json:"seeding_time_limit"`
	SeenComplete             int     `json:"seen_complete"`
	SeqDl                    bool    `json:"seq_dl"`
	Size                     int     `json:"size"`
	State                    string  `json:"state"`
	SuperSeeding             bool    `json:"super_seeding"`
	Tags                     string  `json:"tags"`
	TimeActive               int     `json:"time_active"`
	TotalSize                int     `json:"total_size"`
	Tracker                  string  `json:"tracker"`
	TrackersCount            int     `json:"trackers_count"`
	UpLimit                  int     `json:"up_limit"`
	Uploaded                 int     `json:"uploaded"`
	UploadedSession          int     `json:"uploaded_session"`
	Upspeed                  int     `json:"upspeed"`
}

type TorrentProperties struct {
	AdditionDate           int     `json:"addition_date,omitempty"`
	Comment                string  `json:"comment,omitempty"`
	CompletionDate         int     `json:"completion_date,omitempty"`
	CreatedBy              string  `json:"created_by,omitempty"`
	CreationDate           int     `json:"creation_date,omitempty"`
	DlLimit                int     `json:"dl_limit,omitempty"`
	DlSpeed                int     `json:"dl_speed,omitempty"`
	DlSpeedAvg             int     `json:"dl_speed_avg,omitempty"`
	DownloadPath           string  `json:"download_path,omitempty"`
	Eta                    int     `json:"eta,omitempty"`
	Hash                   string  `json:"hash,omitempty"`
	InfohashV1             string  `json:"infohash_v1,omitempty"`
	InfohashV2             string  `json:"infohash_v2,omitempty"`
	IsPrivate              bool    `json:"is_private,omitempty"`
	LastSeen               int     `json:"last_seen,omitempty"`
	Name                   string  `json:"name,omitempty"`
	NbConnections          int     `json:"nb_connections,omitempty"`
	NbConnectionsLimit     int     `json:"nb_connections_limit,omitempty"`
	Peers                  int     `json:"peers,omitempty"`
	PeersTotal             int     `json:"peers_total,omitempty"`
	PieceSize              int     `json:"piece_size,omitempty"`
	PiecesHave             int     `json:"pieces_have,omitempty"`
	PiecesNum              int     `json:"pieces_num,omitempty"`
	Reannounce             int     `json:"reannounce,omitempty"`
	SavePath               string  `json:"save_path,omitempty"`
	SeedingTime            int     `json:"seeding_time,omitempty"`
	Seeds                  int     `json:"seeds,omitempty"`
	SeedsTotal             int     `json:"seeds_total,omitempty"`
	ShareRatio             float64 `json:"share_ratio,omitempty"`
	TimeElapsed            int     `json:"time_elapsed,omitempty"`
	TotalDownloaded        int64   `json:"total_downloaded,omitempty"`
	TotalDownloadedSession int64   `json:"total_downloaded_session,omitempty"`
	TotalSize              int64   `json:"total_size,omitempty"`
	TotalUploaded          int64   `json:"total_uploaded,omitempty"`
	TotalUploadedSession   int64   `json:"total_uploaded_session,omitempty"`
	TotalWasted            int     `json:"total_wasted,omitempty"`
	UpLimit                int     `json:"up_limit,omitempty"`
	UpSpeed                int     `json:"up_speed,omitempty"`
	UpSpeedAvg             int     `json:"up_speed_avg,omitempty"`
}

type TorrentTracker struct {
	Msg           string `json:"msg,omitempty"`
	NumDownloaded int    `json:"num_downloaded,omitempty"`
	NumLeeches    int    `json:"num_leeches,omitempty"`
	NumPeers      int    `json:"num_peers,omitempty"`
	NumSeeds      int    `json:"num_seeds,omitempty"`
	Status        int    `json:"status,omitempty"`
	Tier          int    `json:"tier,omitempty"`
	URL           string `json:"url,omitempty"`
}

type TorrentWebSeed struct {
	URL string `json:"url"`
}

type TorrentContent struct {
	Availability int    `json:"availability,omitempty"`
	Index        int    `json:"index,omitempty"`
	IsSeed       bool   `json:"is_seed,omitempty"`
	Name         string `json:"name,omitempty"`
	PieceRange   []int  `json:"piece_range,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Progress     int    `json:"progress,omitempty"`
	Size         int64  `json:"size,omitempty"`
}

type TorrentAddOption struct {
	URLs               []string `schema:"-"`
	Torrents           []byte   `schema:"-"`
	SavePath           string   `schema:"save_path,omitempty"`
	Cookies            string   `schema:"cookie,omitempty"`
	Category           string   `schema:"category,omitempty"`
	Tags               []string `schema:"-"`
	SkipChecking       bool     `schema:"skip_checking,omitempty"`
	Paused             bool     `schema:"paused,omitempty"`
	RootFolder         bool     `schema:"root_folder,omitempty"`
	Rename             string   `schema:"rename,omitempty"`
	UpLimit            int      `schema:"upLimit,omitempty"`
	DlLimit            int      `schema:"dlLimit,omitempty"`
	RatioLimit         int      `schema:"ratioLimit,omitempty"`
	SeedingTimeLimit   int      `schema:"seedingTimeLimit,omitempty"`
	AutoTMM            bool     `schema:"autoTMM,omitempty"`
	SequentialDownload bool     `schema:"sequentialDownload,omitempty"`
	FirstLastPiecePrio bool     `schema:"firstLastPiecePrio,omitempty"`
}

type TorrentCategory struct {
	Name     string `json:"name,omitempty"`
	SavePath string `json:"savePath,omitempty"`
}

func (c *client) GetTorrents(opt *TorrentOption) ([]*TorrentInfo, error) {
	var formData = url.Values{}
	err := encoder.Encode(opt, formData)
	if err != nil {
		return nil, err
	}
	if len(opt.Hashes) != 0 {
		formData.Add("hashes", strings.Join(opt.Hashes, "|"))
	}

	apiUrl := fmt.Sprintf("%s/api/v2/torrents/info?%s", c.config.Address, formData.Encode())
	result, err := c.doRequest(&requestData{
		url: apiUrl,
	})
	if err != nil {
		return nil, err
	}

	if result.code != 200 {
		return nil, errors.New("get torrents failed: " + string(result.body))
	}

	fmt.Println(string(result.body))

	var mainData []*TorrentInfo
	if err := sonic.Unmarshal(result.body, &mainData); err != nil {
		return nil, err
	}

	return mainData, nil
}

func (c *client) GetProperties(hash string) (*TorrentProperties, error) {
	apiUrl := fmt.Sprintf("%s/api/v2/torrents/properties?hash=%s", c.config.Address, hash)
	result, err := c.doRequest(&requestData{
		url: apiUrl,
	})
	if err != nil {
		return nil, err
	}

	if result.code != 200 {
		return nil, errors.New("get torrent properties failed: " + string(result.body))
	}

	var mainData = new(TorrentProperties)
	if err := sonic.Unmarshal(result.body, mainData); err != nil {
		return nil, err
	}

	return mainData, nil
}

func (c *client) GetTrackers(hash string) ([]*TorrentTracker, error) {
	apiUrl := fmt.Sprintf("%s/api/v2/torrents/trackers?hash=%s", c.config.Address, hash)
	result, err := c.doRequest(&requestData{
		url: apiUrl,
	})
	if err != nil {
		return nil, err
	}

	if result.code != 200 {
		return nil, errors.New("get torrent trackers failed: " + string(result.body))
	}

	var mainData []*TorrentTracker
	if err := sonic.Unmarshal(result.body, &mainData); err != nil {
		return nil, err
	}

	return mainData, nil
}

func (c *client) GetWebSeeds(hash string) ([]*TorrentWebSeed, error) {
	apiUrl := fmt.Sprintf("%s/api/v2/torrents/webseeds?hash=%s", c.config.Address, hash)
	result, err := c.doRequest(&requestData{
		url: apiUrl,
	})
	if err != nil {
		return nil, err
	}

	if result.code != 200 {
		return nil, errors.New("get torrent web seeds failed: " + string(result.body))
	}

	var mainData []*TorrentWebSeed
	if err := sonic.Unmarshal(result.body, &mainData); err != nil {
		return nil, err
	}

	return mainData, nil
}

func (c *client) GetContents(hash string, indexes ...string) ([]*TorrentContent, error) {
	var apiUrl string
	if len(indexes) != 0 {
		apiUrl = fmt.Sprintf("%s/api/v2/torrents/files?hash=%s&indexes=%s", c.config.Address, hash, strings.Join(indexes, "|"))
	} else {
		apiUrl = fmt.Sprintf("%s/api/v2/torrents/files?hash=%s", c.config.Address, hash)
	}
	result, err := c.doRequest(&requestData{
		url: apiUrl,
	})
	if err != nil {
		return nil, err
	}

	if result.code != 200 {
		return nil, errors.New("get torrent web seeds failed: " + string(result.body))
	}

	var mainData []*TorrentContent
	if err := sonic.Unmarshal(result.body, &mainData); err != nil {
		return nil, err
	}

	return mainData, nil
}

func (c *client) GetPiecesStates(hash string) ([]int, error) {
	var apiUrl = fmt.Sprintf("%s/api/v2/torrents/pieceStates?hash=%s", c.config.Address, hash)
	result, err := c.doRequest(&requestData{
		url: apiUrl,
	})
	if err != nil {
		return nil, err
	}

	if result.code != 200 {
		return nil, errors.New("get torrent pieces states failed: " + string(result.body))
	}

	var mainData []int
	if err := sonic.Unmarshal(result.body, &mainData); err != nil {
		return nil, err
	}

	return mainData, nil
}

func (c *client) GetPiecesHashes(hash string) ([]string, error) {
	var apiUrl = fmt.Sprintf("%s/api/v2/torrents/pieceHashes?hash=%s", c.config.Address, hash)
	result, err := c.doRequest(&requestData{
		url: apiUrl,
	})
	if err != nil {
		return nil, err
	}

	if result.code != 200 {
		return nil, errors.New("get torrent pieces states failed: " + string(result.body))
	}

	var mainData []string
	if err := sonic.Unmarshal(result.body, &mainData); err != nil {
		return nil, err
	}

	return mainData, nil
}

func (c *client) PauseTorrents(hashes []string) error {
	if len(hashes) == 0 {
		return errors.New("no torrent hashes provided")
	}
	var apiUrl = fmt.Sprintf("%s/api/v2/torrents/pieceHashes?hashes=%s", c.config.Address, strings.Join(hashes, "|"))
	result, err := c.doRequest(&requestData{
		url: apiUrl,
	})
	if err != nil {
		return err
	}

	if result.code != 200 {
		return errors.New("pause torrents failed: " + string(result.body))
	}
	return nil
}

func (c *client) ResumeTorrents(hashes []string) error {
	if len(hashes) == 0 {
		return errors.New("no torrent hashes provided")
	}
	var apiUrl = fmt.Sprintf("%s/api/v2/torrents/resume?hashes=%s", c.config.Address, strings.Join(hashes, "|"))
	result, err := c.doRequest(&requestData{
		url: apiUrl,
	})
	if err != nil {
		return err
	}

	if result.code != 200 {
		return errors.New("resume torrents failed: " + string(result.body))
	}
	return nil
}

func (c *client) DeleteTorrents(hashes []string, deleteFile bool) error {
	if len(hashes) == 0 {
		return errors.New("no torrent hashes provided")
	}
	var apiUrl = fmt.Sprintf("%s/api/v2/torrents/resume?hashes=%s&deleteFiles=%t", c.config.Address, strings.Join(hashes, "|"), deleteFile)
	result, err := c.doRequest(&requestData{
		url: apiUrl,
	})
	if err != nil {
		return err
	}

	if result.code != 200 {
		return errors.New("delete torrents failed: " + string(result.body))
	}
	return nil
}

func (c *client) RecheckTorrents(hashes []string) error {
	if len(hashes) == 0 {
		return errors.New("no torrent hashes provided")
	}
	var apiUrl = fmt.Sprintf("%s/api/v2/torrents/recheck?hashes=%s", c.config.Address, strings.Join(hashes, "|"))
	result, err := c.doRequest(&requestData{
		url: apiUrl,
	})
	if err != nil {
		return err
	}

	if result.code != 200 {
		return errors.New("recheck torrents failed: " + string(result.body))
	}
	return nil
}

func (c *client) ReAnnounceTorrents(hashes []string) error {
	if len(hashes) == 0 {
		return errors.New("no torrent hashes provided")
	}
	var apiUrl = fmt.Sprintf("%s/api/v2/torrents/reannounce?hashes=%s", c.config.Address, strings.Join(hashes, "|"))
	result, err := c.doRequest(&requestData{
		url: apiUrl,
	})
	if err != nil {
		return err
	}

	if result.code != 200 {
		return errors.New("reannounce torrents failed: " + string(result.body))
	}
	return nil
}

func (c *client) AddNewTorrent(opt *TorrentAddOption) error {
	panic("implement me")
}

func (c *client) AddTrackers(hash string, urls []string) error {
	if len(urls) == 0 {
		return errors.New("no trackers urls provided")
	}
	var apiUrl = fmt.Sprintf("%s/api/v2/torrents/addTrackers?hash=%s&urls=%s", c.config.Address, hash, strings.Join(urls, "%0A"))
	result, err := c.doRequest(&requestData{
		url: apiUrl,
	})
	if err != nil {
		return err
	}

	if result.code != 200 {
		return errors.New("addTrackers torrents failed: " + string(result.body))
	}
	return nil
}

func (c *client) EditTrackers(hash, origUrl, newUrl string) error {
	var apiUrl = fmt.Sprintf("%s/api/v2/torrents/editTracker?hash=%s&origUrl=%s&newUrl=%s", c.config.Address, hash, origUrl, newUrl)
	result, err := c.doRequest(&requestData{
		url: apiUrl,
	})
	if err != nil {
		return err
	}

	if result.code != 200 {
		return errors.New("editTracker torrents failed: " + string(result.body))
	}
	return nil
}

func (c *client) RemoveTrackers(hash string, urls []string) error {
	if len(urls) == 0 {
		return errors.New("no trackers urls provided")
	}
	var apiUrl = fmt.Sprintf("%s/api/v2/torrents/removeTrackers?hash=%s&urls=%s", c.config.Address, hash, strings.Join(urls, "|"))
	result, err := c.doRequest(&requestData{
		url: apiUrl,
	})
	if err != nil {
		return err
	}

	if result.code != 200 {
		return errors.New("removeTrackers torrents failed: " + string(result.body))
	}
	return nil
}

func (c *client) AddPeers(hashes []string, peers []string) error {
	if len(hashes) == 0 {
		return errors.New("no hashes provided")
	}
	if len(peers) == 0 {
		return errors.New("no peers provided")
	}
	var apiUrl = fmt.Sprintf("%s/api/v2/torrents/addPeers?hashes=%s&peers=%s", c.config.Address, strings.Join(hashes, "|"), strings.Join(peers, "|"))
	result, err := c.doRequest(&requestData{
		url: apiUrl,
	})
	if err != nil {
		return err
	}

	if result.code != 200 {
		return errors.New("addPeers torrents failed: " + string(result.body))
	}
	return nil
}

func (c *client) IncreasePriority(hashes []string) error {
	if len(hashes) == 0 {
		return errors.New("no hashes provided")
	}
	var apiUrl = fmt.Sprintf("%s/api/v2/torrents/increasePrio?hashes=%s", c.config.Address, strings.Join(hashes, "|"))
	result, err := c.doRequest(&requestData{
		url: apiUrl,
	})
	if err != nil {
		return err
	}

	if result.code != 200 {
		return errors.New("increasePrio torrents failed: " + string(result.body))
	}
	return nil
}

func (c *client) DecreasePriority(hashes []string) error {
	if len(hashes) == 0 {
		return errors.New("no hashes provided")
	}
	var apiUrl = fmt.Sprintf("%s/api/v2/torrents/decreasePrio?hashes=%s", c.config.Address, strings.Join(hashes, "|"))
	result, err := c.doRequest(&requestData{
		url: apiUrl,
	})
	if err != nil {
		return err
	}

	if result.code != 200 {
		return errors.New("decreasePrio torrents failed: " + string(result.body))
	}
	return nil
}

func (c *client) MaxPriority(hashes []string) error {
	if len(hashes) == 0 {
		return errors.New("no hashes provided")
	}
	var apiUrl = fmt.Sprintf("%s/api/v2/torrents/topPrio?hashes=%s", c.config.Address, strings.Join(hashes, "|"))
	result, err := c.doRequest(&requestData{
		url: apiUrl,
	})
	if err != nil {
		return err
	}

	if result.code != 200 {
		return errors.New("topPrio torrents failed: " + string(result.body))
	}
	return nil
}

func (c *client) MinPriority(hashes []string) error {
	if len(hashes) == 0 {
		return errors.New("no hashes provided")
	}
	var apiUrl = fmt.Sprintf("%s/api/v2/torrents/bottomPrio?hashes=%s", c.config.Address, strings.Join(hashes, "|"))
	result, err := c.doRequest(&requestData{
		url: apiUrl,
	})
	if err != nil {
		return err
	}

	if result.code != 200 {
		return errors.New("bottomPrio torrents failed: " + string(result.body))
	}
	return nil
}

func (c *client) SetFilePriority(hash string, id string, priority int) error {
	var apiUrl = fmt.Sprintf("%s/api/v2/torrents/bottomPrio?hash=%s&id=%s&priority=%d", c.config.Address, hash, id, priority)
	result, err := c.doRequest(&requestData{
		url: apiUrl,
	})
	if err != nil {
		return err
	}

	if result.code != 200 {
		return errors.New("bottomPrio torrents failed: " + string(result.body))
	}
	return nil
}

func (c *client) GetDownloadLimit(hashes []string) (map[string]int, error) {
	//TODO implement me
	panic("implement me")
}

func (c *client) SetDownloadLimit(hashes []string, limit int) error {
	//TODO implement me
	panic("implement me")
}

func (c *client) SetShareLimit(hashes []string, ratioLimit float64, seedingTimeLimit, inactiveSeedingTimeLimit int) error {
	//TODO implement me
	panic("implement me")
}

func (c *client) GetUploadLimit(hashes []string) (map[string]int, error) {
	//TODO implement me
	panic("implement me")
}

func (c *client) SetUploadLimit(hashes []string, limit int) error {
	//TODO implement me
	panic("implement me")
}

func (c *client) SetLocation(hashes []string, location string) error {
	//TODO implement me
	panic("implement me")
}

func (c *client) SetName(hash string, name string) error {
	//TODO implement me
	panic("implement me")
}

func (c *client) SetCategory(hashes []string, category string) error {
	//TODO implement me
	panic("implement me")
}

func (c *client) GetCategories() (map[string]*TorrentCategory, error) {
	//TODO implement me
	panic("implement me")
}

func (c *client) AddNewCategory(category, savePath string) error {
	//TODO implement me
	panic("implement me")
}

func (c *client) EditCategory(category, savePath string) error {
	//TODO implement me
	panic("implement me")
}

func (c *client) RemoveCategories(categories []string) error {
	//TODO implement me
	panic("implement me")
}

func (c *client) AddTags(hashes []string, tags []string) error {
	//TODO implement me
	panic("implement me")
}

func (c *client) RemoveTags(hashes []string, tags []string) error {
	//TODO implement me
	panic("implement me")
}

func (c *client) GetTags() ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (c *client) CreateTags(tags []string) error {
	//TODO implement me
	panic("implement me")
}

func (c *client) DeleteTags(tags []string) error {
	//TODO implement me
	panic("implement me")
}

func (c *client) SetAutomaticManagement(hashes []string, enable bool) error {
	//TODO implement me
	panic("implement me")
}

func (c *client) ToggleSequentialDownload(hashes []string) error {
	//TODO implement me
	panic("implement me")
}

func (c *client) SetFirstLastPiecePriority(hashes []string) error {
	//TODO implement me
	panic("implement me")
}

func (c *client) SetForceStart(hashes []string, force bool) error {
	//TODO implement me
	panic("implement me")
}

func (c *client) SetSuperSeeding(hashes []string, enable bool) error {
	//TODO implement me
	panic("implement me")
}

func (c *client) RenameFile(hash, oldPath, newPath string) error {
	//TODO implement me
	panic("implement me")
}

func (c *client) RenameFolder(hash, oldPath, newPath string) error {
	//TODO implement me
	panic("implement me")
}
