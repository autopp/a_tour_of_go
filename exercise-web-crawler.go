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

type Cache struct {
  urls map[string]bool
  mutex sync.Mutex
}

func (c *Cache) Add(url string) {
  c.mutex.Lock()
  c.urls[url] = true
  c.mutex.Unlock()
}

func (c *Cache) Contains(url string) bool {
  c.mutex.Lock()
  defer c.mutex.Unlock()
  _, b := c.urls[url]
  return b
}

func crawlWith(url string, depth int, fetcher Fetcher, cache *Cache, ch chan bool) {
  if depth <= 0 {
		ch <- false
    return
	}
  if cache.Contains(url) {
    ch <- false
    return
  }
	body, urls, err := fetcher.Fetch(url)
  cache.Add(url)
	if err != nil {
		fmt.Println(err)
		ch <- false
    return
	}
	fmt.Printf("found: %s %q\n", url, body)
  n := len(urls)
  chs := make([]chan bool, n)
	for i, u := range urls {
    chs[i] = make(chan bool)
		go crawlWith(u, depth-1, fetcher, cache, chs[i])
	}
  for _, c := range chs {
    <- c
  }
	ch <- true
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
  ch := make(chan bool)
	go crawlWith(url, depth, fetcher, &Cache{urls: make(map[string]bool)}, ch)
  <- ch
}

func main() {
	Crawl("http://golang.org/", 4, fetcher)
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
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
