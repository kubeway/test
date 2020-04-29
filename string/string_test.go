package http

import (
	"testing"
)

func TestStr(t *testing.T) {
	_url := "*.ask.kuberun.com"
	if _url[0] == '*' {
		t.Log(_url[1:])
	}

}
