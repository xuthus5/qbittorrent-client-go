package qbittorrent

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/bytedance/sonic"
)

type TransferStatusBar struct {
	ConnectionStatus  string `json:"connection_status,omitempty"`
	DhtNodes          int    `json:"dht_nodes,omitempty"`
	DlInfoData        int64  `json:"dl_info_data,omitempty"`
	DlInfoSpeed       int    `json:"dl_info_speed,omitempty"`
	DlRateLimit       int    `json:"dl_rate_limit,omitempty"`
	UpInfoData        int    `json:"up_info_data,omitempty"`
	UpInfoSpeed       int    `json:"up_info_speed,omitempty"`
	UpRateLimit       int    `json:"up_rate_limit,omitempty"`
	Queueing          bool   `json:"queueing,omitempty"`
	UseAltSpeedLimits bool   `json:"use_alt_speed_limits,omitempty"`
	RefreshInterval   int    `json:"refresh_interval,omitempty"`
}

type Transfer interface {
	// GlobalStatusBar usually see in qBittorrent status bar
	GlobalStatusBar() (*TransferStatusBar, error)
	// BanPeers the peer to ban, or multiple peers separated by a pipe.
	// each peer is host:port
	BanPeers(peers []string) error
	// GetSpeedLimitsMode get alternative speed limits state
	GetSpeedLimitsMode() (string, error)
	// ToggleSpeedLimitsMode toggle alternative speed limits
	ToggleSpeedLimitsMode() error
	// GetGlobalUploadLimit get global upload limit, the response is the value of current global download speed
	// limit in bytes/second; this value will be zero if no limit is applied.
	GetGlobalUploadLimit() (string, error)
	// SetGlobalUploadLimit set global upload limit, set in bytes/second
	SetGlobalUploadLimit(int) error
	// GetGlobalDownloadLimit get global download limit, the response is the value of current global download speed
	// limit in bytes/second; this value will be zero if no limit is applied.
	GetGlobalDownloadLimit() (string, error)
	// SetGlobalDownloadLimit set global download limit, set in bytes/second
	SetGlobalDownloadLimit(int) error
}

func (c *client) GlobalStatusBar() (*TransferStatusBar, error) {
	apiUrl := fmt.Sprintf("%s/api/v2/transfer/info", c.config.Address)
	result, err := c.doRequest(&requestData{
		url: apiUrl,
	})
	if err != nil {
		return nil, err
	}

	if result.code != 200 {
		return nil, errors.New("get global transfer status bar failed: " + string(result.body))
	}

	var data = new(TransferStatusBar)
	if err := sonic.Unmarshal(result.body, data); err != nil {
		return nil, err
	}

	return data, nil
}

func (c *client) BanPeers(peers []string) error {
	apiUrl := fmt.Sprintf("%s/api/v2/transfer/banPeers", c.config.Address)
	var form = url.Values{}
	form.Add("peers", strings.Join(peers, "|"))
	result, err := c.doRequest(&requestData{
		url:    apiUrl,
		method: http.MethodPost,
		body:   strings.NewReader(form.Encode()),
	})
	if err != nil {
		return err
	}

	if result.code != 200 {
		return errors.New("ban peers failed: " + string(result.body))
	}

	return nil
}

func (c *client) GetSpeedLimitsMode() (string, error) {
	apiUrl := fmt.Sprintf("%s/api/v2/transfer/speedLimitsMode", c.config.Address)
	result, err := c.doRequest(&requestData{
		url: apiUrl,
	})
	if err != nil {
		return "", err
	}

	if result.code != 200 {
		return "", errors.New("ban peers failed: " + string(result.body))
	}

	return string(result.body), nil
}

func (c *client) ToggleSpeedLimitsMode() error {
	apiUrl := fmt.Sprintf("%s/api/v2/transfer/toggleSpeedLimitsMode", c.config.Address)
	result, err := c.doRequest(&requestData{
		url:    apiUrl,
		method: http.MethodPost,
	})
	if err != nil {
		return err
	}

	if result.code != 200 {
		return errors.New("ban peers failed: " + string(result.body))
	}

	return nil
}

func (c *client) GetGlobalUploadLimit() (string, error) {
	apiUrl := fmt.Sprintf("%s/api/v2/transfer/uploadLimit", c.config.Address)
	result, err := c.doRequest(&requestData{
		url: apiUrl,
	})
	if err != nil {
		return "", err
	}

	if result.code != 200 {
		return "", errors.New("get global upload limit failed: " + string(result.body))
	}

	return string(result.body), nil
}

func (c *client) SetGlobalUploadLimit(limit int) error {
	apiUrl := fmt.Sprintf("%s/api/v2/transfer/setUploadLimit?limit=%d", c.config.Address, limit)
	result, err := c.doRequest(&requestData{
		url: apiUrl,
	})
	if err != nil {
		return err
	}

	if result.code != 200 {
		return errors.New("set global upload limit failed: " + string(result.body))
	}

	return nil
}

func (c *client) GetGlobalDownloadLimit() (string, error) {
	apiUrl := fmt.Sprintf("%s/api/v2/transfer/downloadLimit", c.config.Address)
	result, err := c.doRequest(&requestData{
		url: apiUrl,
	})
	if err != nil {
		return "", err
	}

	if result.code != 200 {
		return "", errors.New("get global download limit failed: " + string(result.body))
	}

	return string(result.body), nil
}

func (c *client) SetGlobalDownloadLimit(limit int) error {
	apiUrl := fmt.Sprintf("%s/api/v2/transfer/setDownloadLimit?limit=%d", c.config.Address, limit)
	result, err := c.doRequest(&requestData{
		url: apiUrl,
	})
	if err != nil {
		return err
	}

	if result.code != 200 {
		return errors.New("set global download limit failed: " + string(result.body))
	}

	return nil
}
