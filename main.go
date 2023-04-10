package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.github.com/",
		"https://www.reddit.com",
		"https://golang.org",
	}

	titles := fetchTitlesConcurrently(urls, 4)
	for _, title := range titles {
		fmt.Println(title)
	}
}

func fetchTitle(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "Error fetching url", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	r := regexp.MustCompile(`<title>(.+)</title>`)
	matches := r.FindStringSubmatch(string(body))
	if len(matches) < 2 {
		return "", fmt.Errorf("no title found")
	}

	title := strings.TrimSpace(matches[1])
	return title, nil
}

func fetchTitlesConcurrently(urls []string, numWorkers int) []string {
	urlsCh := make(chan string, len(urls))
	titlesCh := make(chan string, len(urls))

	for i := 0; i < numWorkers; i++ {
		go func() {
			for url := range urlsCh {
				title, err := fetchTitle(url)
				if err != nil {
					titlesCh <- fmt.Sprintf("Error for %s: %v", url, err)
				} else {
					titlesCh <- fmt.Sprintf("Title for %s: %s", url, title)
				}
			}
		}()
	}

	for _, url := range urls {
		urlsCh <- url
	}
	close(urlsCh)

	titles := make([]string, 0, len(urls))
	for i := 0; i < len(urls); i++ {
		title := <-titlesCh
		titles = append(titles, title)
	}

	return titles
}
