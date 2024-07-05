package qbittorrent

// Client represents a qBittorrent client
type Client interface {
	// Authentication auth qBittorrent client
	Authentication() Authentication
	// Application get qBittorrent application info
	Application() Application
	// Log get qBittorrent log
	Log() Log
	// Sync get qBittorrent events
	Sync() Sync
	// Transfer transfer info
	Transfer() Transfer
	// Torrent manage
	Torrent() Torrent
	// Search api
	Search() Search
	// RSS api
	RSS() RSS
}

func NewClient(cfg *Config) (Client, error) {
	var c = &client{config: cfg, clientPool: newClientPool(cfg.ConnectionMaxIdles, cfg.ConnectionTimeout)}
	if err := c.Authentication().Login(); err != nil {
		return nil, err
	}
	if cfg.RefreshCookie {
		go c.refreshCookie()
	}
	return c, nil
}
