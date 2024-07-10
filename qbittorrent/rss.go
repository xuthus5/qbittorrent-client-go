package qbittorrent

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/bytedance/sonic"
)

type RSS interface {
	// AddFolder create new folder for rss, full path of added folder such as "The Pirate Bay\Top100"
	AddFolder(path string) error
	// AddFeed add feed
	AddFeed(*RssAddFeedOption) error
	// RemoveItem remove folder or feed
	RemoveItem(path string) error
	// MoveItem move or rename folder or feed
	MoveItem(srcPath, destPath string) error
	// GetItems list all items, if withData is true, will return all data
	GetItems(withData bool) (map[string]interface{}, error)
	// MarkAsRead if articleId is provided only the article is marked as read otherwise the whole feed
	// is going to be marked as read.
	MarkAsRead(*RssMarkAsReadOption) error
	// RefreshItem refresh folder or feed
	RefreshItem(itemPath string) error
	// SetAutoDownloadingRule set auto-downloading rule
	SetAutoDownloadingRule(ruleName string, ruleDef *RssAutoDownloadingRuleDef) error
	// RenameAutoDownloadingRule rename auto-downloading rule
	RenameAutoDownloadingRule(ruleName, newRuleName string) error
	// RemoveAutoDownloadingRule remove auto-downloading rule
	RemoveAutoDownloadingRule(ruleName string) error
	// GetAllAutoDownloadingRules get all auto-downloading rules
	GetAllAutoDownloadingRules() (map[string]*RssAutoDownloadingRuleDef, error)
	// GetAllArticlesMatchingRule get all articles matching a rule
	GetAllArticlesMatchingRule(ruleName string) (map[string][]string, error)
}

type RssAddFeedOption struct {
	// URL feed of rss such as http://thepiratebay.org/rss//top100/200
	URL string `schema:"url"`
	// Folder full path of added folder, optional
	Folder string `schema:"path,omitempty"`
}

type RssMarkAsReadOption struct {
	// ItemPath current full path of item
	ItemPath string `schema:"itemPath"`
	// ArticleId id of article, optional
	ArticleId string `schema:"articleId,omitempty"`
}

type RssAutoDownloadingRuleDefTorrentParams struct {
	Category                 string   `json:"category,omitempty"`
	DownloadLimit            int      `json:"download_limit,omitempty"`
	DownloadPath             int      `json:"download_path,omitempty"`
	InactiveSeedingTimeLimit int      `json:"inactive_seeding_time_limit,omitempty"`
	OperatingMode            string   `json:"operating_mode,omitempty"`
	RatioLimit               int      `json:"ratio_limit,omitempty"`
	SavePath                 string   `json:"save_path,omitempty"`
	SeedingTimeLimit         int      `json:"seeding_time_limit,omitempty"`
	SkipChecking             bool     `json:"skip_checking,omitempty"`
	Tags                     []string `json:"tags,omitempty"`
	UploadLimit              int      `json:"upload_limit,omitempty"`
	Stopped                  bool     `json:"stopped,omitempty"`
	UseAutoTMM               bool     `json:"use_auto_tmm,omitempty"`
}

type RssAutoDownloadingRuleDef struct {
	AddPaused                 bool                                    `json:"addPaused,omitempty"`
	AffectedFeeds             []string                                `json:"affectedFeeds,omitempty"`
	AssignedCategory          string                                  `json:"assignedCategory,omitempty"`
	Enabled                   bool                                    `json:"enabled,omitempty"`
	EpisodeFilter             string                                  `json:"episodeFilter,omitempty"`
	IgnoreDays                int                                     `json:"ignoreDays,omitempty"`
	LastMatch                 string                                  `json:"lastMatch,omitempty"`
	MustContain               string                                  `json:"mustContain,omitempty"`
	MustNotContain            string                                  `json:"mustNotContain,omitempty"`
	PreviouslyMatchedEpisodes []string                                `json:"previouslyMatchedEpisodes,omitempty"`
	Priority                  int                                     `json:"priority,omitempty"`
	SavePath                  string                                  `json:"savePath,omitempty"`
	SmartFilter               bool                                    `json:"smartFilter,omitempty"`
	TorrentParams             *RssAutoDownloadingRuleDefTorrentParams `json:"torrentParams,omitempty"`
	UseRegex                  bool                                    `json:"useRegex,omitempty"`
}

func (c *client) AddFolder(path string) error {
	var formData = url.Values{}
	formData.Add("path", path)
	var apiUrl = fmt.Sprintf("%s/api/v2/rss/addFolder", c.config.Address)
	result, err := c.doRequest(&requestData{
		url:    apiUrl,
		method: http.MethodPost,
		body:   strings.NewReader(formData.Encode()),
	})
	if err != nil {
		return err
	}

	if result.code != 200 {
		return errors.New("add rss folder failed: " + string(result.body))
	}
	return nil
}

func (c *client) AddFeed(opt *RssAddFeedOption) error {
	var formData = url.Values{}
	err := encoder.Encode(opt, formData)
	if err != nil {
		return err
	}
	var apiUrl = fmt.Sprintf("%s/api/v2/rss/addFolder", c.config.Address)
	result, err := c.doRequest(&requestData{
		url:    apiUrl,
		method: http.MethodPost,
		body:   strings.NewReader(formData.Encode()),
	})
	if err != nil {
		return err
	}

	if result.code != 200 {
		return errors.New("add rss feed failed: " + string(result.body))
	}
	return nil
}

func (c *client) RemoveItem(path string) error {
	var formData = url.Values{}
	formData.Add("path", path)
	var apiUrl = fmt.Sprintf("%s/api/v2/rss/removeItem", c.config.Address)
	result, err := c.doRequest(&requestData{
		url:    apiUrl,
		method: http.MethodPost,
		body:   strings.NewReader(formData.Encode()),
	})
	if err != nil {
		return err
	}

	if result.code != 200 {
		return errors.New("remove rss item failed: " + string(result.body))
	}
	return nil
}

func (c *client) MoveItem(srcPath, destPath string) error {
	var formData = url.Values{}
	formData.Add("itemPath", srcPath)
	formData.Add("destPath", destPath)
	var apiUrl = fmt.Sprintf("%s/api/v2/rss/moveItem", c.config.Address)
	result, err := c.doRequest(&requestData{
		url:    apiUrl,
		method: http.MethodPost,
		body:   strings.NewReader(formData.Encode()),
	})
	if err != nil {
		return err
	}

	if result.code != 200 {
		return errors.New("move rss item failed: " + string(result.body))
	}
	return nil
}

func (c *client) GetItems(withData bool) (map[string]interface{}, error) {
	var apiUrl = fmt.Sprintf("%s/api/v2/rss/items?withData=%t", c.config.Address, withData)
	result, err := c.doRequest(&requestData{
		url:    apiUrl,
		method: http.MethodGet,
	})
	if err != nil {
		return nil, err
	}

	if result.code != 200 {
		return nil, errors.New("get rss items failed: " + string(result.body))
	}
	var data = make(map[string]interface{})
	err = sonic.Unmarshal(result.body, &data)
	return data, err
}

func (c *client) MarkAsRead(opt *RssMarkAsReadOption) error {
	var formData = url.Values{}
	err := encoder.Encode(opt, formData)
	if err != nil {
		return err
	}
	var apiUrl = fmt.Sprintf("%s/api/v2/rss/markAsRead", c.config.Address)
	result, err := c.doRequest(&requestData{
		url:    apiUrl,
		method: http.MethodPost,
		body:   strings.NewReader(formData.Encode()),
	})
	if err != nil {
		return err
	}

	if result.code != 200 {
		return errors.New("mark as read rss item failed: " + string(result.body))
	}
	return nil
}

func (c *client) RefreshItem(itemPath string) error {
	var formData = url.Values{}
	formData.Add("itemPath", itemPath)
	var apiUrl = fmt.Sprintf("%s/api/v2/rss/refreshItem", c.config.Address)
	result, err := c.doRequest(&requestData{
		url:    apiUrl,
		method: http.MethodPost,
		body:   strings.NewReader(formData.Encode()),
	})
	if err != nil {
		return err
	}

	if result.code != 200 {
		return errors.New("refresh rss item failed: " + string(result.body))
	}
	return nil
}

func (c *client) SetAutoDownloadingRule(ruleName string, ruleDef *RssAutoDownloadingRuleDef) error {
	var formData = url.Values{}
	formData.Add("ruleName", ruleName)
	ruleDefBytes, err := sonic.Marshal(ruleDef)
	if err != nil {
		return err
	}
	formData.Add("ruleDef", string(ruleDefBytes))
	var apiUrl = fmt.Sprintf("%s/api/v2/rss/setRule", c.config.Address)
	result, err := c.doRequest(&requestData{
		url:    apiUrl,
		method: http.MethodPost,
		body:   strings.NewReader(formData.Encode()),
	})
	if err != nil {
		return err
	}

	if result.code != 200 {
		return errors.New("set auto downloading rule failed: " + string(result.body))
	}
	return nil
}

func (c *client) RenameAutoDownloadingRule(ruleName, newRuleName string) error {
	var formData = url.Values{}
	formData.Add("ruleName", ruleName)
	formData.Add("newRuleName", newRuleName)
	var apiUrl = fmt.Sprintf("%s/api/v2/rss/renameRule", c.config.Address)
	result, err := c.doRequest(&requestData{
		url:    apiUrl,
		method: http.MethodPost,
		body:   strings.NewReader(formData.Encode()),
	})
	if err != nil {
		return err
	}

	if result.code != 200 {
		return errors.New("rename auto downloading rule failed: " + string(result.body))
	}
	return nil
}

func (c *client) RemoveAutoDownloadingRule(ruleName string) error {
	var formData = url.Values{}
	formData.Add("ruleName", ruleName)
	var apiUrl = fmt.Sprintf("%s/api/v2/rss/removeRule", c.config.Address)
	result, err := c.doRequest(&requestData{
		url:    apiUrl,
		method: http.MethodPost,
		body:   strings.NewReader(formData.Encode()),
	})
	if err != nil {
		return err
	}

	if result.code != 200 {
		return errors.New("remove auto downloading rule failed: " + string(result.body))
	}
	return nil
}

func (c *client) GetAllAutoDownloadingRules() (map[string]*RssAutoDownloadingRuleDef, error) {
	var apiUrl = fmt.Sprintf("%s/api/v2/rss/matchingArticles", c.config.Address)
	result, err := c.doRequest(&requestData{
		url: apiUrl,
	})
	if err != nil {
		return nil, err
	}
	if result.code != 200 {
		return nil, errors.New("get rss rules failed: " + string(result.body))
	}
	var data = make(map[string]*RssAutoDownloadingRuleDef)
	err = sonic.Unmarshal(result.body, &data)
	return data, err
}

func (c *client) GetAllArticlesMatchingRule(ruleName string) (map[string][]string, error) {
	var formData = url.Values{}
	formData.Add("ruleName", ruleName)
	var apiUrl = fmt.Sprintf("%s/api/v2/rss/matchingArticles?%s", c.config.Address, formData.Encode())
	result, err := c.doRequest(&requestData{
		url: apiUrl,
	})
	if err != nil {
		return nil, err
	}
	if result.code != 200 {
		return nil, errors.New("get rss rule match articles failed: " + string(result.body))
	}
	var data = make(map[string][]string)
	err = sonic.Unmarshal(result.body, &data)
	return data, err
}
