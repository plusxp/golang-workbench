package main

import (
	"log"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

func crawl(urlTarget url.URL, domainFilter string) {
	if urlTarget.Host != domainFilter {
		log.Printf("'%+v' didn't pass the domain filter '%s', returning.\n", urlTarget.String(), domainFilter)
		pageOutsideDomain++
		return
	}

	fetched.Lock()
	if _, ok := fetched.m[urlTarget]; ok {
		fetched.Unlock()
		log.Printf("Already fetched '%+v', returning.\n", urlTarget.String())
		return
	}

	fetched.m[urlTarget] = errFetchInProgress
	fetched.Unlock()

	doc, errRead := goquery.NewDocumentFromReader(getResponse(urlTarget))

	// TODO: get errors passed back from my getResponse() func
	fetched.Lock()
	fetched.m[urlTarget] = errRead
	fetched.Unlock()

	if errRead != nil {
		log.Fatalf("%+v\n", errRead)
	}

	log.Printf("Fetched '%+v'.\n", urlTarget.String())

	// Get all links on this page, and store them for later reference
	crawled.Lock()
	if _, ok := crawled.m[urlTarget]; ok {
		crawled.Unlock()
		log.Printf("Already crawled '%+v', returning.\n", urlTarget)
		return
	}

	// Keeping the children URLs in a seperate slice like this is a bit of a hack
	// I don't like it but it got me past some locking issues
	// TODO: learn more about locking and clean this up
	children := make([]url.URL, 0)
	crawled.m[urlTarget] = make([]url.URL, 0)

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href := s.AttrOr("href", "")
		urlHref := convertURL(href, domainFilter)

		log.Printf("'%+v' is a child of '%v'.\n", urlHref, urlTarget.String())

		children = append(children, *urlHref)
	})
	crawled.m[urlTarget] = children
	crawled.Unlock()

	// Now start crawlers on all of this page's children
	done := make(chan bool)
	for b, c := range children {
		log.Printf("Crawling child %+v/%+v of %+v: '%+v'\n", b+1, len(children), urlTarget.String(), c.String())

		go func(u url.URL) {
			crawl(u, domainFilter)
			done <- true
		}(c)
	}

	for x, y := range children {
		log.Printf("<- [%+v] %+v/%+v - waiting for child: %+v\n", urlTarget.String(), x+1, len(children), y.String())
		<-done
	}

	log.Printf("Done with '%+v'.\n", urlTarget.String())
	pageCrawled++
}