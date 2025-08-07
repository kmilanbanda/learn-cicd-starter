package auth

import(
	"testing"
	"net/http"
)

func TestGetAPIKeyNoHeader(t *testing.T) {	
	headers := http.Header{}
	str, err := GetAPIKey(headers)
	if str != "" || err != ErrNoAuthHeaderIncluded {
		 t.Errorf("GetAPIKey() = (%v, %v), want ('', ErrNoAuthHeaderIncluded)", str, err)
	}
}

func TestGetAPIKeyWithHeader(t *testing.T) {
	headers := http.Header{}
	headers.Add("Authorization", "ApiKey 12345")
	str, err := GetAPIKey(headers)
	if str != "12345" || err != nil {
		t.Errorf("GetAPIKey() = (%v, %v), want (\"12345\", nil)", str, err)
	}
}

func TestGetAPIKeyMalformedHeader(t *testing.T) {
	headers := http.Header{}
	headers.Add("Authorization", "12345")
	str, err := GetAPIKey(headers)
	if str != "" || err.Error() != "malformed authorization header" {
		t.Errorf("GetAPIKey() = (%v, %v), want ('', errors.New(\"malformed authorization header\")", str, err)
	}
}
