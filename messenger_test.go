package messenger

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func setClient(code int, body []byte) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(code)
		w.Write(body)
	}))

	http.DefaultClient.Transport = &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}
	return server
}

func TestCheckIntegrity(t *testing.T) {
	if !checkIntegrity("e6af24be1d683c8c911949f897eea1f6", []byte(`{"object":"page","entry":[{"id":1751036168465324,"time":1460923697656,"messaging":[{"sender":{"id":982337261802700},"recipient":{"id":1751036168465324},"timestamp":1460923697635,"message":{"mid":"mid.1460923697625:5c96e8279b55505308","seq":614,"text":"Test \u00e4\u00eb\u00ef"}}]}]}`), "f1a4569dcf02a9829a15696d949b386b7d6d0272") {
		t.Error("Message integrity verification does not work")
	}

	if checkIntegrity("e6af24be1d683c8c911949f897eea1f6", []byte(`{"object":"page","entry":[]}`), "f1a4569dcf02a9829a15696d949b386b7d6d0272") {
		t.Error("Message integrity verification does not work")
	}
}
