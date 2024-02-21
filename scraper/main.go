package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Struct to hold YouTube video details
type YouTubeVideo struct {
	URL   string
	Title string
	Depth int
}

var visited = make(map[string]bool)
var videos []YouTubeVideo

func scrapeWebsite(baseURL string, currentURL string, depth int) {
	// Avoid revisiting URLs
	if visited[currentURL] {
		return
	}
	visited[currentURL] = true

	// Make HTTP GET request
	response, err := http.Get(currentURL)
	if err != nil {
		log.Printf("Error fetching URL %s: %s", currentURL, err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		log.Printf("Status code error: %d %s", response.StatusCode, response.Status)
		return
	}

	// Parse the HTML document
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Printf("Error parsing HTML: %s", err)
		return
	}

	// Find all links and recursively scrape them
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Attr("href")
		if link == "" {
			return
		}

		// Resolve relative URLs
		absoluteURL, err := url.Parse(link)
		if err != nil {
			return
		}
		base, _ := url.Parse(baseURL)
		absoluteURL = base.ResolveReference(absoluteURL)

		if strings.Contains(absoluteURL.String(), "youtube.com") || strings.Contains(absoluteURL.String(), "youtu.be") {
			title := s.Text()
			videos = append(videos, YouTubeVideo{URL: absoluteURL.String(), Title: title, Depth: depth})
		} else if strings.HasPrefix(absoluteURL.String(), baseURL) {
			// Follow the link with increased depth
			scrapeWebsite(baseURL, absoluteURL.String(), depth+1)
		}
	})
}

func writeToFile(videos []YouTubeVideo) {
	file, err := os.Create("youtube_links.txt")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	for _, video := range videos {
		line := fmt.Sprintf("Depth: %d, URL: %s, Title: %s\n", video.Depth, video.URL, video.Title)
		_, err := file.WriteString(line)
		if err != nil {
			log.Fatal("Cannot write to file", err)
		}
	}
}

func main() {
	startURL := "https://feranowibisono.com" // Change to your target website
	scrapeWebsite(startURL, startURL, 0)
	writeToFile(videos)
}
