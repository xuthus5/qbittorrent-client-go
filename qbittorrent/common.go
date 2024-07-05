package qbittorrent

import "github.com/gorilla/schema"

const (
	ContentTypeJSON           = "application/json"
	ContentTypeFormUrlEncoded = "application/x-www-form-urlencoded"
)

var encoder = schema.NewEncoder()
