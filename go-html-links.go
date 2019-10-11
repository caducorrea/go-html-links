package links

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/caducorrea/gontains"
	"golang.org/x/net/html"
)

// All to get all link from a string content
func All(content io.Reader) []string {
	links := []string{}
	col := []string{}

	page := html.NewTokenizer(content)
	for {
		tokenType := page.Next()
		if tokenType == html.ErrorToken {
			return links
		}
		token := page.Token()

		if tokenType == html.StartTagToken && token.DataAtom.String() == "a" {
			for _, attr := range token.Attr {
				if attr.Key == "href" {
					fmt.Println(attr.Val)
					tl := trimHash(attr.Val)
					col = append(col, tl)
					resolv(&links, col)
				}
			}
		}
	}
}

func trimHash(l string) string {
	if strings.Contains(l, "#") {
		var index int
		for n, str := range l {
			if strconv.QuoteRune(str) == "'#'" {
				index = n
				break
			}
		}
		return l[:index]
	}
	return l
}

func resolv(sl *[]string, ml []string) {
	for _, str := range ml {
		if gontains.Gontains(*sl, str) == false {
			*sl = append(*sl, str)
		}
	}
}
