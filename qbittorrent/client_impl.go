package qbittorrent

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"
)

type responseResult struct {
	code    int
	body    []byte
	cookies []*http.Cookie
}

type requestData struct {
	method      string
	url         string
	contentType string
	body        io.Reader
	debug       bool
}

var (
	ErrNotLogin   = errors.New("not login")
	ErrAuthFailed = errors.New("auth failed")
)

var _ Client = (*client)(nil)

type client struct {
	config     *Config
	clientPool *clientPool
	cookieJar  *cookiejar.Jar
}

func (c *client) Authentication() Authentication {
	return c
}

func (c *client) Application() Application {
	return c
}

func (c *client) Log() Log {
	return c
}

func (c *client) Sync() Sync {
	return c
}

func (c *client) Transfer() Transfer {
	return c
}

func (c *client) Torrent() Torrent {
	return c
}

func (c *client) Search() Search {
	return c
}

func (c *client) RSS() RSS {
	return c
}

// doRequest send request
func (c *client) doRequest(data *requestData) (*responseResult, error) {
	if data.method == "" {
		data.method = "GET"
	}
	if data.contentType == "" {
		data.contentType = ContentTypeFormUrlEncoded
	}
	request, err := http.NewRequest(data.method, data.url, data.body)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", data.contentType)
	for key, value := range c.config.CustomHeaders {
		request.Header.Set(key, value)
	}
	hc := c.clientPool.GetClient()
	defer c.clientPool.ReleaseClient(hc)
	if c.cookieJar != nil {
		hc.Jar = c.cookieJar
	}

	if data.debug {
		dumpRequest, err := httputil.DumpRequest(request, true)
		if err != nil {
			return nil, err
		}
		fmt.Println(string(dumpRequest))
	}

	resp, err := hc.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	readAll, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &responseResult{code: resp.StatusCode, body: readAll, cookies: resp.Cookies()}, nil
}

func (c *client) cookies() (string, error) {
	if c.cookieJar == nil {
		return "", ErrNotLogin
	}
	u, err := url.Parse(c.config.Address)
	if err != nil {
		return "", err
	}
	cookies := c.cookieJar.Cookies(u)
	if len(cookies) == 0 {
		return "", ErrNotLogin
	}
	var builder strings.Builder
	for _, cookie := range cookies {
		builder.WriteString(fmt.Sprintf("%s=%s; ", cookie.Name, cookie.Value))
	}

	return builder.String(), nil
}

func (c *client) refreshCookie() {
	if c.config.RefreshIntervals == 0 {
		c.config.RefreshIntervals = time.Hour
	}
	var ticker = time.NewTicker(c.config.RefreshIntervals)
	for range ticker.C {
		if err := c.Authentication().Logout(); err != nil {
			// todo
		}
	}
}
