package weblinks

import (
	"golang.org/x/net/html"
	"io"
)

func Get(httpBody io.Reader) []string {

	// Initialize slice of string to store URL
	links := make([]string, 0)

	// NewTokenizer returns a new HTML Tokenizer for the given Reader.
	htmlToken := html.NewTokenizer(httpBody)

	for {

		// Next scans the next token and returns its type.
		tokenType := htmlToken.Next()

		// Token returns the next Token
		token := htmlToken.Token()

		switch tokenType {

		case html.ErrorToken:
			htmlToken.Err()
			return links

		case html.StartTagToken:
			// StartTagToken looks like <a>
			// DataAtom is integer representation of html tagname
			if token.DataAtom.String() == "a" {

				for _, attr := range token.Attr {

					if attr.Key == "href" {
						// Push all the URL to links array of slice
						links = append(links, attr.Val)
					}
				}
			}
		}
	}
}
