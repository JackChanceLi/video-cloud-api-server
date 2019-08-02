package url

import (
	"testing"
)

func TestNewRtmpUrl(t *testing.T) {
	url, err := NewRtmpUrl()
	if err != nil {
		t.Errorf("Error of generate rtmpUrl:%v", err)
	}
	t.Log(url)
}

func TestNewFlvUrl(t *testing.T) {
	url, err := NewFlvUrl()
	if err != nil {
		t.Errorf("Error of generate flvUrl:%v", err)
	}
	t.Log(url)
}

func TestNewHlsUrl(t *testing.T) {
	url, err := NewHlsUrl()
	if err != nil {
		t.Errorf("Error of generate hlsUrl:%v", err)
	}
	t.Log(url)
}