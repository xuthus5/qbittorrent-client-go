package qbittorrent

import "testing"

func TestClient_Version(t *testing.T) {
	version, err := c.Application().Version()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(version)
}

func TestClient_WebApiVersion(t *testing.T) {
	version, err := c.Application().WebApiVersion()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(version)
}

func TestClient_BuildInfo(t *testing.T) {
	info, err := c.Application().BuildInfo()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("build: %+v", info)
}

func TestClient_Shutdown(t *testing.T) {
	//if err := c.Application().Shutdown(); err != nil {
	//	t.Fatal(err)
	//}
	t.Log("shutting down")
}

func TestClient_GetPreferences(t *testing.T) {
	prefs, err := c.Application().GetPreferences()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("prefs: %+v", prefs)
}

func TestClient_SetPreferences(t *testing.T) {
	prefs, err := c.Application().GetPreferences()
	if err != nil {
		t.Fatal(err)
	}

	prefs.FileLogAge = 301
	if err := c.Application().SetPreferences(prefs); err != nil {
		t.Fatal(err)
	}
	t.Logf("success")
}

func TestClient_DefaultSavePath(t *testing.T) {
	path, err := c.Application().DefaultSavePath()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("path: %s", path)
}
