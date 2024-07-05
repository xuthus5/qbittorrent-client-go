package qbittorrent

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/bytedance/sonic"
)

type LogOption struct {
	Normal      bool  `schema:"normal,omitempty"`        // include normal messages
	Info        bool  `schema:"info,omitempty"`          // include info messages
	Warning     bool  `schema:"warning,omitempty"`       // include warning messages
	Critical    bool  `schema:"critical,omitempty"`      // include critical messages
	LastKnownId int64 `schema:"last_known_id,omitempty"` // exclude messages with "message id" <= (default: last_known_id-1)
}

type LogEntry struct {
	Id        int    `json:"id,omitempty"`        // id of the message or peer
	Timestamp int    `json:"timestamp,omitempty"` // seconds since epoch
	Type      int    `json:"type,omitempty"`      // type of the message, Log::NORMAL: 1, Log::INFO: 2, Log::WARNING: 4, Log::CRITICAL: 8
	Message   string `json:"message,omitempty"`   // text of the message
	IP        string `json:"ip"`                  // ip of the peer
	Blocked   bool   `json:"blocked,omitempty"`   // whether the peer was blocked
	Reason    string `json:"reason,omitempty"`    // Reason of the block
}

type Log interface {
	// GetLog get log
	GetLog(option *LogOption) ([]*LogEntry, error)
	// GetPeerLog get peer log
	GetPeerLog(lastKnownId int) ([]*LogEntry, error)
}

func (c *client) GetLog(option *LogOption) ([]*LogEntry, error) {
	var form = url.Values{}
	err := encoder.Encode(option, form)
	if err != nil {
		return nil, err
	}
	apiUrl := fmt.Sprintf("%s/api/v2/log/main?%s", c.config.Address, form.Encode())

	result, err := c.doRequest(&requestData{
		url:  apiUrl,
		body: strings.NewReader(form.Encode()),
	})
	if err != nil {
		return nil, err
	}

	if result.code != 200 {
		return nil, errors.New("get log failed: " + string(result.body))
	}

	var logs []*LogEntry
	if err := sonic.Unmarshal(result.body, &logs); err != nil {
		return nil, err
	}

	return logs, nil
}

func (c *client) GetPeerLog(lastKnownId int) ([]*LogEntry, error) {
	apiUrl := fmt.Sprintf("%s/api/v2/log/peers", c.config.Address)
	var form = url.Values{}
	form.Add("last_known_id", strconv.Itoa(lastKnownId))

	result, err := c.doRequest(&requestData{
		url:  apiUrl,
		body: strings.NewReader(form.Encode()),
	})
	if err != nil {
		return nil, err
	}

	if result.code != 200 {
		return nil, errors.New("get peer log failed: " + string(result.body))
	}

	var logs []*LogEntry
	if err := sonic.Unmarshal(result.body, &logs); err != nil {
		return nil, err
	}

	return logs, nil
}
