package qbittorrent

import (
	"net/http"
	"time"
)

// clientPool defines a pool of HTTP clients
type clientPool struct {
	// clients channel to store http.Client instances
	clients chan *http.Client
	// maxIdle maximum number of idle clients in the pool, default 128
	maxIdle int
	// timeout for each http.Client, default 3 seconds
	timeout time.Duration
}

// newClientPool creates and returns a new clientPool
func newClientPool(maxIdle int, timeout time.Duration) *clientPool {
	if maxIdle == 0 {
		maxIdle = 128
	}
	if timeout == 0 {
		timeout = time.Second * 3
	}
	pool := &clientPool{
		clients: make(chan *http.Client, maxIdle),
		maxIdle: maxIdle,
		timeout: timeout,
	}
	for i := 0; i < maxIdle; i++ {
		// pre-create maxIdle number of http.Client instances
		pool.clients <- &http.Client{
			Timeout: timeout,
		}
	}
	return pool
}

// GetClient retrieves a http.Client from the pool
func (p *clientPool) GetClient() *http.Client {
	var client *http.Client
	select {
	case client = <-p.clients:
	default:
		// if the channel is empty, create a new client
		client = &http.Client{
			Timeout: p.timeout,
		}
	}
	return client
}

// ReleaseClient returns a http.Client back to the pool
func (p *clientPool) ReleaseClient(client *http.Client) {
	select {
	case p.clients <- client:
	default:
		// if the channel is full, discard and close the client
		client.CloseIdleConnections()
	}
}
