package gop5js

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouter(t *testing.T) {
	ts := httptest.NewServer(newRouter())
	defer ts.Close()
	testUrls := []string{
		"lib/p5.js",
		"sketch.js",
		"",
	}
	for _, tURL := range testUrls {
		res, err := http.Get(ts.URL + PathPrefix + "/" + tURL)

		if err != nil {
			t.Error(err)
		}
		if res.StatusCode != 200 {
			t.Errorf("Did not find %s", tURL)
		}
	}
}
