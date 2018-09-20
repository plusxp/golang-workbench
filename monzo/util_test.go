package main

import (
	"fmt"
	"net/url"
	"testing"
)

func TestConvertURL(t *testing.T) {
	// Domain filter argument used in each convertURL() call
	// This is normally initilised/parsed from CLI in: main.go > init()
	flagURL = "monzo.com"

	testCases := []struct {
		base      string
		href      string
		urlResult url.URL
	}{
		{
			base:      flagURL,
			href:      "/",
			urlResult: url.URL{Scheme: "https", Host: flagURL, Path: "/"},
		},
		{
			base:      fmt.Sprintf("http://%s", flagURL),
			href:      "/",
			urlResult: url.URL{Scheme: "http", Host: flagURL, Path: "/"},
		},
		{
			base:      fmt.Sprintf("https://%s", flagURL),
			href:      "/",
			urlResult: url.URL{Scheme: "https", Host: flagURL, Path: "/"},
		},
		{
			base:      flagURL,
			href:      "/about",
			urlResult: url.URL{Scheme: "https", Host: flagURL, Path: "/about"},
		},
		{
			base:      fmt.Sprintf("http://%s", flagURL),
			href:      "/about",
			urlResult: url.URL{Scheme: "http", Host: flagURL, Path: "/about"},
		},
		{
			base:      fmt.Sprintf("https://%s", flagURL),
			href:      "/about",
			urlResult: url.URL{Scheme: "https", Host: flagURL, Path: "/about"},
		},
		{
			base:      flagURL,
			href:      fmt.Sprintf("//%s/about", flagURL),
			urlResult: url.URL{Scheme: "https", Host: flagURL, Path: "/about"},
		},
		{
			base:      fmt.Sprintf("http://%s", flagURL),
			href:      fmt.Sprintf("//%s/about", flagURL),
			urlResult: url.URL{Scheme: "http", Host: flagURL, Path: "/about"},
		},
		{
			base:      fmt.Sprintf("https://%s", flagURL),
			href:      fmt.Sprintf("//%s/about", flagURL),
			urlResult: url.URL{Scheme: "https", Host: flagURL, Path: "/about"},
		},
		{
			base:      flagURL,
			href:      fmt.Sprintf("http://%s/about", flagURL),
			urlResult: url.URL{Scheme: "http", Host: flagURL, Path: "/about"},
		},
		{
			base:      fmt.Sprintf("https://%s", flagURL),
			href:      fmt.Sprintf("http://%s/about", flagURL),
			urlResult: url.URL{Scheme: "http", Host: flagURL, Path: "/about"},
		},
		{
			base:      fmt.Sprintf("https://%s", flagURL),
			href:      "//facebook.com/about",
			urlResult: url.URL{Scheme: "https", Host: "facebook.com", Path: "/about"},
		},
		{
			base:      fmt.Sprintf("https://%s", flagURL),
			href:      "http://twitter.com/monzo",
			urlResult: url.URL{Scheme: "http", Host: "twitter.com", Path: "/monzo"},
		},
		{
			base:      fmt.Sprintf("https://%s", flagURL),
			href:      "/fragment#foo",
			urlResult: url.URL{Scheme: "https", Host: flagURL, Path: "/fragment", Fragment: "foo"},
		},
		{
			base:      fmt.Sprintf("https://%s", flagURL),
			href:      "/query?foo=bar",
			urlResult: url.URL{Scheme: "https", Host: flagURL, Path: "/query", RawQuery: "foo=bar"},
		},
		{
			base:      fmt.Sprintf("https://%s", flagURL),
			href:      fmt.Sprintf("%s/bare-domain-same", flagURL),
			urlResult: url.URL{Scheme: "https", Host: flagURL, Path: "/bare-domain-same"},
		},
		{
			base:      fmt.Sprintf("https://%s", flagURL),
			href:      fmt.Sprintf("%s/bare/domain/same", flagURL),
			urlResult: url.URL{Scheme: "https", Host: flagURL, Path: "/bare/domain/same"},
		},
		{
			base:      fmt.Sprintf("https://%s", flagURL),
			href:      "twitter.com/bare-domain-different",
			urlResult: url.URL{Scheme: "https", Host: "twitter.com", Path: "/bare-domain-different"},
		},
		{
			base:      fmt.Sprintf("https://%s", flagURL),
			href:      "twitter.com/bare/domain/different",
			urlResult: url.URL{Scheme: "https", Host: "twitter.com", Path: "/bare/domain/different"},
		},
		{
			base:      fmt.Sprintf("https://%s", flagURL),
			href:      "samelevel",
			urlResult: url.URL{Scheme: "https", Host: flagURL, Path: "/samelevel"},
		},
		{
			base:      fmt.Sprintf("https://%s/about", flagURL),
			href:      "..",
			urlResult: url.URL{Scheme: "https", Host: flagURL, Path: "/"},
		},
		{
			base:      fmt.Sprintf("https://%s/about/stuff/and/things", flagURL),
			href:      "../also/widgets",
			urlResult: url.URL{Scheme: "https", Host: flagURL, Path: "/about/stuff/also/widgets"},
		},
		{
			base:      fmt.Sprintf("https://%s/about/stuff/and/things/", flagURL),
			href:      "../../also/widgets",
			urlResult: url.URL{Scheme: "https", Host: flagURL, Path: "/about/stuff/also/widgets"},
		},
		{
			base:      fmt.Sprintf("https://%s/one-directory/", flagURL),
			href:      "../../../go/up/and/down/a/lot/of/levels",
			urlResult: url.URL{Scheme: "https", Host: flagURL, Path: "/go/up/and/down/a/lot/of/levels"},
		},
		{
			base:      fmt.Sprintf("https://%s/one-directory/", flagURL),
			href:      "../../../go//up/../and/../down/a/../lot/../of/../levels",
			urlResult: url.URL{Scheme: "https", Host: flagURL, Path: "/go/down/levels"},
		},
		{
			base:      flagURL,
			href:      "/double//slashes//in//the//path",
			urlResult: url.URL{Scheme: "https", Host: flagURL, Path: "/double/slashes/in/the/path"},
		},
		{
			base:      fmt.Sprintf("https://%s", flagURL),
			href:      "/double//slashes//in//the//path",
			urlResult: url.URL{Scheme: "https", Host: flagURL, Path: "/double/slashes/in/the/path"},
		},
	}

	for _, tC := range testCases {
		desc := fmt.Sprintf("[%s + %s -> %s]", tC.base, tC.href, tC.urlResult.String())
		t.Run(desc, func(t *testing.T) {
			actual := convertURL(tC.base, tC.href)
			if *actual != tC.urlResult {
				t.Errorf("Got '%q', want '%q'.\n[%#v]\n[%#v]\n", actual.String(), tC.urlResult.String(), actual, tC.urlResult)
			}
		})
	}
}
