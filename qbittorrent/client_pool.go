package qbittorrent

import (
	"crypto/tls"
	"net"
	"net/http"
	"sync"
	"time"
)

// clientPool defines a pool of HTTP clients
type clientPool struct {
	// pool store http.Client instances
	*sync.Pool
}

// newClientPool creates and returns a new clientPool
func newClientPool(maxIdle int, timeout time.Duration) *clientPool {
	if maxIdle == 0 {
		maxIdle = 128
	}
	if timeout == 0 {
		timeout = time.Second * 3
	}
	return &clientPool{
		Pool: &sync.Pool{
			New: func() any {
				return &http.Client{
					Transport: &http.Transport{
						Proxy: http.ProxyFromEnvironment,
						DialContext: (&net.Dialer{
							Timeout:   30 * time.Second,
							KeepAlive: 30 * time.Second,
						}).DialContext,
						TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
						MaxIdleConns:    maxIdle,
					},
					Timeout: timeout,
				}
			},
		},
	}
}

// GetClient retrieves a http.Client from the pool
func (p *clientPool) GetClient() *http.Client {
	return p.Get().(*http.Client)
}

// ReleaseClient returns a http.Client back to the pool
func (p *clientPool) ReleaseClient(client *http.Client) {
	p.Put(client)
}
