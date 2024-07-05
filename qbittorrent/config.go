package qbittorrent

import "time"

type Config struct {
	// Address qBittorrent endpoint
	Address string
	// Username used to access the WebUI
	Username string
	// Password used to access the WebUI
	Password string

	// HTTP configuration

	// CustomHeaders custom headers
	CustomHeaders map[string]string
	// ConnectionTimeout request timeout, default 3 seconds
	ConnectionTimeout time.Duration
	// ConnectionMaxIdles http client pool, default 128
	ConnectionMaxIdles int
	// RefreshCookie whether to automatically refresh cookies
	RefreshCookie bool
	// SessionTimeout interval for refreshing cookies, default 1 hour
	RefreshIntervals time.Duration
}
