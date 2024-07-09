package qbittorrent

import "errors"

var (
	ErrNotLogin   = errors.New("not login")
	ErrAuthFailed = errors.New("auth failed")
)
