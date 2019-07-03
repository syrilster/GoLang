package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type UrlLookupMap struct {
	value map[string]string
	mutex sync.Mutex
}

func (urlMap *UrlLookupMap) put(key string, body string) {
	urlMap.mutex.Lock()
	defer urlMap.mutex.Unlock()
	urlMap.value[key] = body
}

func (urlMap *UrlLookupMap) get(key string) (string, bool) {
	urlMap.mutex.Lock()
	defer urlMap.mutex.Unlock()
	val, ok := urlMap.value[key]
	return val, ok
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, urlMap UrlLookupMap) {
	defer waitGroup.Done()
	body, urls, err := fetcher.Fetch(url)
	urlMap.put(url, body)
	if depth <= 0 {
		return
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		if _, ok := urlMap.get(u); !ok {
			waitGroup.Add(1)
			go Crawl(u, depth-1, fetcher, urlMap)
		}
	}
	return
}

var waitGroup sync.WaitGroup

func main() {
	urlMap := UrlLookupMap{value: make(map[string]string)}
	waitGroup.Add(1)
	go Crawl("https://golang.org/", 4, fetcher, urlMap)
	defer waitGroup.Done()

	for url := range urlMap.value {
		body, _ := urlMap.get(url)
		fmt.Printf("found: %s %q\n", url, body)
	}
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
