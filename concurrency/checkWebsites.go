package main

import "time"

type WebsiteChecker func(string) bool

// So A note for young n!ggas if you're reading this commit,
// actually maps and slices are not thread-safe in Golang
// i.e if you're trying to access a map or slice concurrently
// via two different goroutines the code will run into race conditions!

// To solve this data race we use channels!

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		go func() {
			results[url] = wc(url)
		}()
	}

	time.Sleep(2 * time.Second)
	return results
}
