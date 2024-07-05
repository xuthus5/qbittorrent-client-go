# qbittorrent-client-go

qbittorrent-client-go is a Go client for qBittorrent, this library applies to qBittorrent v4.1+.

qBittorrent is an open-source BitTorrent client for all platforms, find more about qBittorrent at https://github.com/qbittorrent/qBittorrent

## Requirements

- Go 1.19

## Usage

Import the Client Library:

```go
import "github.com/xuthus5/qbittorrent-client-go/qbittorrent"
```

Create a Client:

```go
config := &qbittorrent.Config{
    Address:           "http://localhost:8080",
    Username:          "admin",
    Password:          "admin",
    RefreshIntervals:  time.Hour,
    ConnectionTimeout: time.Second * 3,
}
client, err := qbittorrent.NewClient(config)
if err != nil {
	// do something
}

if err := client.Authentication().Login(); err != nil {
    // do something
}
```