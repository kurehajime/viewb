package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(handler))
	defer ts.Close()

	//HTTP OK?
	res, err := http.Get(ts.URL)
	if err != nil {
		t.Error("unexpected")
		return
	}

	//STATUS OK?
	if res.StatusCode != 200 {
		t.Error("Status code error")
		return
	}
}
func TestCmd(t *testing.T) {
	//Command OK?
	ans := "HelloWorld"
	res := cmd("echo HelloWorld")
	if strings.Trim(ans, " \r\n") != strings.Trim(res, " \r\n") {
		t.Error("cmd Error:", ans, res)
		return
	}
}
