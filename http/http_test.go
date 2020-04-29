package http

import (
	"net/url"
	"strings"
	"testing"
)

func TestUrl(t *testing.T) {
	_url := "http://www.kuberun.com/coffee   xfdfdf?adfd=cc"
	u, err := url.Parse(_url)
	if err != nil {
		t.Fatalf("parse url error:%s", err.Error())
	}

	t.Logf("host:%s", u.Hostname())
	t.Logf("path:%s", url.PathEscape(u.Path))
	t.Logf("url string:%s", u.String())
	_u, _ := url.Parse(u.String())
	t.Logf("url path:%s", _u.Path)
	t.Logf("url path:%s", strings.ReplaceAll(strings.ToLower(url.PathEscape(u.Path)), "%2f", "/"))

}
