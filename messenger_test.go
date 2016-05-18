package messenger

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/drewolson/testflight"
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
	if !checkIntegrity("e6af24be1d683c8c911949f897eea1f6", []byte(`{"object":"page","entry":[{"id":"1751036168465324","time":1460923697656,"messaging":[{"sender":{"id":"982337261802700"},"recipient":{"id":"1751036168465324"},"timestamp":1460923697635,"message":{"mid":"mid.1460923697625:5c96e8279b55505308","seq":614,"text":"Test \u00e4\u00eb\u00ef"}}]}]}`), "da611bd448dc12acdf0cd3ab33fdb3adaee26145") {
		t.Error("Message integrity verification does not work")
	}

	if checkIntegrity("e6af24be1d683c8c911949f897eea1f6", []byte(`{"object":"page","entry":[]}`), "f1a4569dcf02a9829a15696d949b386b7d6d0272") {
		t.Error("Message integrity verification does not work")
	}
}

func TestHandler(t *testing.T) {
	mess := &Messenger{
		AccessToken: "foo",
		VerifyToken: "bar",
	}
	testflight.WithServer(http.HandlerFunc(mess.Handler), func(r *testflight.Requester) {
		// Legit verify request
		response := r.Get("/?hub.verify_token=bar&hub.challenge=zoo")
		if response.StatusCode != http.StatusOK {
			t.Errorf("Invalid status code, expected %d, got: %d", http.StatusOK, response.StatusCode)
		}
		if response.Body != "zoo" {
			t.Error("Invalid body.")
		}
		// Invalid verify_token
		response = r.Get("/?hub.verify_token=abba&hub.challenge=zoo")
		if response.StatusCode != http.StatusUnauthorized {
			t.Errorf("Invalid status code, expected %d, got: %d", http.StatusUnauthorized, response.StatusCode)
		}
		if response.Body != "" {
			t.Error("Invalid body, expected to be empty.")
		}
		// Invalid method
		response = r.Put("/", "application/json", "foo-bar")
		if response.StatusCode != http.StatusMethodNotAllowed {
			t.Errorf("Invalid status code, expected %d, got: %d", http.StatusMethodNotAllowed, response.StatusCode)
		}

		mess.AppSecret = "e6af24be1d683c8c911949f897eea1f6"
		// Legit Post request
		postRequest, err := http.NewRequest("POST", "/", strings.NewReader(`{"object":"page","entry":[{"id":"1751036168465324","time":1460923697656,"messaging":[{"sender":{"id":"982337261802700"},"recipient":{"id":"1751036168465324"},"timestamp":1460923697635,"message":{"mid":"mid.1460923697625:5c96e8279b55505308","seq":614,"text":"Test \u00e4\u00eb\u00ef"}}]}]}`))
		if err != nil {
			t.Error(err)
		}
		postRequest.Header.Add("x-hub-signature", "sha1=da611bd448dc12acdf0cd3ab33fdb3adaee26145")
		response = r.Do(postRequest)
		if response.StatusCode != http.StatusOK {
			t.Errorf("Invalid status code, expected %d, got: %d", http.StatusOK, response.StatusCode)
		}

		// Invalid signature
		mess.AppSecret = "abc"

		// Invalid signature
		response = r.Post("/", "application/json", `{"object":"page","entry":[{"id":"1751036168465324","time":1460923697656,"messaging":[{"sender":{"id":"982337261802700"},"recipient":{"id":"1751036168465324"},"timestamp":1460923697635,"message":{"mid":"mid.1460923697625:5c96e8279b55505308","seq":614,"text":"Test \u00e4\u00eb\u00ef"}}]}]}`)
		if response.StatusCode != http.StatusBadRequest {
			t.Errorf("Invalid status code, expected %d, got: %d", http.StatusBadRequest, response.StatusCode)
		}

		mess.AppSecret = ""
		// Invalid request
		response = r.Post("/", "application/json", `{"object":"page","entry":[{"id":"1751036168465324","time":1460923697656,"messaging":[{"sender":{"id":"982337261802701751036168465324"},"timestamp":1460923697635,"message":{"mid":"mid.1460923697625:5c96e8279b55505308","seq":614,"text":"Test \u00e4\u00eb\u00e`)
		if response.StatusCode != http.StatusBadRequest {
			t.Errorf("Invalid status code, expected %d, got: %d", http.StatusBadRequest, response.StatusCode)
		}

	})
}
