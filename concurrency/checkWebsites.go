package main

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

// So A note for young n!ggas if you're reading this commit,
// actually maps and slices are not thread-safe in Golang
// i.e if you're trying to access a map or slice concurrently
// via two different goroutines the code will run into race conditions!

// To solve this data race we use channels!

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func() {
			resultChannel <- result{url, wc(url)}
		}()
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}

// Goroutines, the basic unit of concurrency in GO,
// which let us manage more than one website check request

// anonymous functions, which we used to start each of the concurrent processes that check websites.

// channels, to help organize and control the communication between the different
// processes, allowing us to avoid a race condition bug.

// race detector(bhai isne time le liya boht gcc CGO enable karna pada iske liye)
// helps us to debug problems with concurrent code by using
// go test -race

// Make It Work, Make It Right, Make it FAST - Kent Beck
