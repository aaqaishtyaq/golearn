package concurrency

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	if url == "wat://somerandomewebsite.io" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://aaqa.dev",
		"wat://somerandomewebsite.io",
	}

	want := map[string]bool{
		"http://google.com":           true,
		"http://aaqa.dev":             true,
		"wat://somerandomewebsite.io": false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Wanted %v, got %v", want, got)
	}
}
