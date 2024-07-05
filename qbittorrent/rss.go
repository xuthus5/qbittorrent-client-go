package qbittorrent

type RSS interface {
	// AddFolder create new folder for rss, full path of added folder such as "The Pirate Bay\Top100"
	AddFolder(path string) error
	// AddFeed add feed to path, url of rss feed such as http://thepiratebay.org/rss//top100/200
	AddFeed(url string, path ...string) error
	// RemoveItem remove folder or feed
	RemoveItem(path string) error
	// MoveItem move or rename folder or feed
	MoveItem(srcPath, destPath string) error
	// GetItems list all items
	GetItems(withData bool) (map[string]interface{}, error)
	// MarkAsRead if articleId is provided only the article is marked as read otherwise the whole feed
	// is going to be marked as read.
	MarkAsRead(path, articleId string) error
	// RefreshItem refresh folder or feed
	RefreshItem(itemPath string) error
	// SetAutoDownloadingRule set auto-downloading rule
	SetAutoDownloadingRule(ruleName, ruleDef string) error
	// RenameAutoDownloadingRule rename auto-downloading rule
	RenameAutoDownloadingRule(ruleName, newRuleName string) error
	// RemoveAutoDownloadingRule remove auto-downloading rule
	RemoveAutoDownloadingRule(ruleName string) error
	// GetAllAutoDownloadingRules get all auto-downloading rules
	GetAllAutoDownloadingRules()
	// GetAllArticlesMatchingRule get all articles matching a rule
	GetAllArticlesMatchingRule(ruleName string) (map[string]interface{}, error)
}

func (c *client) AddFolder(path string) error {
	//TODO implement me
	panic("implement me")
}

func (c *client) AddFeed(url string, path ...string) error {
	//TODO implement me
	panic("implement me")
}

func (c *client) RemoveItem(path string) error {
	//TODO implement me
	panic("implement me")
}

func (c *client) MoveItem(srcPath, destPath string) error {
	//TODO implement me
	panic("implement me")
}

func (c *client) GetItems(withData bool) (map[string]interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (c *client) MarkAsRead(path, articleId string) error {
	//TODO implement me
	panic("implement me")
}

func (c *client) RefreshItem(itemPath string) error {
	//TODO implement me
	panic("implement me")
}

func (c *client) SetAutoDownloadingRule(ruleName, ruleDef string) error {
	//TODO implement me
	panic("implement me")
}

func (c *client) RenameAutoDownloadingRule(ruleName, newRuleName string) error {
	//TODO implement me
	panic("implement me")
}

func (c *client) RemoveAutoDownloadingRule(ruleName string) error {
	//TODO implement me
	panic("implement me")
}

func (c *client) GetAllAutoDownloadingRules() {
	//TODO implement me
	panic("implement me")
}

func (c *client) GetAllArticlesMatchingRule(ruleName string) (map[string]interface{}, error) {
	//TODO implement me
	panic("implement me")
}
