package qbittorrent

import "testing"

func TestClient_Login(t *testing.T) {
	if err := c.Authentication().Login(); err != nil {
		t.Fatal(err)
	}
}

func TestClient_Logout(t *testing.T) {
	if err := c.Authentication().Login(); err != nil {
		t.Fatal(err)
	}

	if err := c.Authentication().Logout(); err != nil {
		t.Fatal(err)
	}
}
