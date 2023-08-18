package main

import (
	"fmt"
	"net/http"
	"net/url"

	"golang.org/x/net/html"
)

var (
	urls    []string
	visited map[string]bool = map[string]bool{}
)

func main() {
	visitUrl("https://github.com/Gabrielgqa")

	fmt.Println(len(urls))
}

func visitUrl(url string) {
	if ok := visited[url]; ok {
		return
	}

	visited[url] = true
	fmt.Println(url)
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("status diferente de 200: %d", resp.StatusCode))
	}

	doc, err := html.Parse(resp.Body)

	if err != nil {
		panic(err)
	}

	extractLinksFromUrl(doc)
}

func extractLinksFromUrl(element *html.Node) {
	if element.Type == html.ElementNode && element.Data == "a" {
		for _, attr := range element.Attr {
			if attr.Key != "href" {
				continue
			}

			link, err := url.Parse(attr.Val)
			if err != nil || link.Scheme == "" {
				continue
			}

			urls = append(urls, link.String())

			visitUrl(link.String())
		}
	}

	for c := element.FirstChild; c != nil; c = c.NextSibling {
		extractLinksFromUrl(c)
	}
}
