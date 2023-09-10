package reload

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestShouldSuccessfullyReloadConfigurations(t *testing.T) {
	// given
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("invalid http method")
		}
		if r.Header.Get("Content-Type") != "application/json" {
			t.Fatalf("invalid Content-Type header")
		}
		switch r.URL.RequestURI() {
		case "/valid":
			w.WriteHeader(http.StatusNoContent)
		default:
			t.Fatalf("invalid uri")
		}

	})
	server := httptest.NewServer(handler)

	// when
	err := reloadConfigurations(server.URL + "/valid")

	// then
	if err != nil {
		t.Fatalf("error should be nil: %v", err)
	}
}

func TestShouldFailWhileReloadingConfigurations(t *testing.T) {
	// given
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("invalid http method")
		}
		if r.Header.Get("Content-Type") != "application/json" {
			t.Fatalf("invalid Content-Type header")
		}
		switch r.URL.RequestURI() {
		case "/invalid":
			w.WriteHeader(http.StatusNotAcceptable)
		default:
			t.Fatalf("invalid uri")
		}

	})
	server := httptest.NewServer(handler)

	// when
	err := reloadConfigurations(server.URL + "/invalid")

	// then
	if err == nil {
		t.Fatalf("error should not be nil: %v", err)
	}
	if err.Error() != "invalid response: 406 Not Acceptable" {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestReloadConfigurations(t *testing.T) {
	// given
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("invalid http method")
		}
		if r.Header.Get("Content-Type") != "application/json" {
			t.Fatalf("invalid Content-Type header")
		}
		switch r.URL.RequestURI() {
		case "/valid":
			w.WriteHeader(http.StatusNoContent)
		case "/invalid":
			w.WriteHeader(http.StatusNotAcceptable)
		default:
			t.Fatalf("invalid uri")
		}

	})
	server := httptest.NewServer(handler)

	tests := []struct {
		name              string
		uri               string
		expectedErrString string
	}{
		{
			name:              "reloadConfigurations should succeed",
			uri:               "/valid",
			expectedErrString: "",
		},
		{
			name:              "reloadConfigurations should fail",
			uri:               "/invalid",
			expectedErrString: "invalid response: 406 Not Acceptable",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// when
			err := reloadConfigurations(server.URL + tc.uri)

			// then

			if err != nil {
				if err.Error() != tc.expectedErrString {
					t.Fatalf("unexpected err, want: %v, got: %v", tc.expectedErrString, err)
				}
			} else {
				if tc.expectedErrString != "" {
					t.Fatalf("unexpected err: %s", tc.expectedErrString)
				}
			}

		})
	}
}
