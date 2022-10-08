package concurrency

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	if url == "waat://imjustlearning.here" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://imjustlearning.here",
	}

	want := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://imjustlearning.here": false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	assert.Equal(t, want, got)
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
