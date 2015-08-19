package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	user = "user"
	pass = "pass"

	ts := httptest.NewServer(http.HandlerFunc(handler))
	defer ts.Close()

	//Basic auth failure pattern
	res, err := http.Get(ts.URL)
	if err != nil {
		t.Error("unexpected")
		return
	}

	//Basic auth working ?
	if res.StatusCode == 200 {
		t.Error("Basic auth not working")
		return
	} else if res.StatusCode != 401 {
		t.Error("Status code error:" + string(res.StatusCode))
	}

	//Basic auth success pattern
	res, err = http.Get(strings.Replace(ts.URL, "http://", "http://"+user+":"+pass+"@", 1))
	if err != nil {
		t.Error("unexpected")
		return
	}

	//STATUS OK?
	if res.StatusCode != 200 {
		t.Error("Status code error:" + string(res.StatusCode))
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
