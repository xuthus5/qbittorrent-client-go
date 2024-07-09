package qbittorrent

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

type Authentication interface {
	// Login cookie-based authentication, after calling NewClient, do not need to call Login again,
	// it is the default behavior
	Login() error
	// Logout deactivate cookies
	Logout() error
}

func (c *client) Login() error {
	if c.config.Username == "" || c.config.Password == "" {
		return errors.New("username or password is empty")
	}

	formData := url.Values{}
	formData.Set("username", c.config.Username)
	formData.Set("password", c.config.Password)
	encodedFormData := formData.Encode()

	apiUrl := fmt.Sprintf("%s/api/v2/auth/login", c.config.Address)

	result, err := c.doRequest(&requestData{
		method: http.MethodPost,
		url:    apiUrl,
		body:   strings.NewReader(encodedFormData),
	})
	if err != nil {
		return err
	}

	if result.code != 200 {
		return errors.New("login failed: " + string(result.body))
	}

	if string(result.body) == "Fails." {
		return ErrAuthFailed
	}

	if string(result.body) != "Ok." {
		return errors.New("login failed: " + string(result.body))
	}

	if c.cookieJar == nil {
		c.cookieJar, err = cookiejar.New(nil)
		if err != nil {
			return err
		}
	}

	u, err := url.Parse(c.config.Address)
	if err != nil {
		return err
	}
	c.cookieJar.SetCookies(u, result.cookies)

	return nil
}

func (c *client) Logout() error {
	apiUrl := fmt.Sprintf("%s/api/v2/auth/logout", c.config.Address)
	result, err := c.doRequest(&requestData{
		method: http.MethodPost,
		url:    apiUrl,
	})
	if err != nil {
		return err
	}

	if result.code != 200 {
		return errors.New("logout failed: " + string(result.body))
	}

	return nil
}
