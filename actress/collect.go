package actress

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/syohex/dmm"
	"path"
)

var keywords = []string{
	"a", "i", "u", "e", "o",
	"ka", "ki", "ku", "ke", "ko",
	"sa", "si", "su", "se", "so",
	"ta", "ti", "tu", "te", "to",
	"ta", "ti", "ne", "no",
	"ha", "hi", "hu", "he", "ho",
	"ma", "mi", "mu", "me", "mo",
	"ma", "mi", "mu", "me", "mo",
	"ya", "yu", "yo",
	"ra", "ri", "ru", "re", "ro",
	"wa",
}

func actressPage(key string, page int) string {
	return fmt.Sprintf("http://www.dmm.co.jp/digital/videoa/-/actress/=/keyword=%s/page=%d/", key, page)
}

// Keywords returns all possible keywords of actresses
func Keywords() []string {
	return keywords
}

var pageRegexp = regexp.MustCompile(`全(\d+)ページ中`)

func extractPageCount(url string) (int, error) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return 0, err
	}

	pageInfo := doc.Find(".list-boxpagenation").Find("p").Text()
	matches := pageRegexp.FindAllStringSubmatch(pageInfo, 1)
	if len(matches) == 0 {
		return 0, fmt.Errorf("not found page information in %s", url)
	}

	pages, err := strconv.Atoi(matches[0][1])
	if err != nil {
		return 0, err
	}

	return pages, nil
}

var actressIDRegexp = regexp.MustCompile(`id=(\d+)`)

func extractActressID(url string) (int, error) {
	matches := actressIDRegexp.FindAllStringSubmatch(url, 1)
	if len(matches) == 0 {
		return 0, fmt.Errorf("ID is not found from '%s'", url)
	}

	id, err := strconv.Atoi(matches[0][1])
	if err != nil {
		return 0, err
	}

	return id, nil
}

// ExtractActresses returns actresses whose names start with 'key'
func ExtractActresses(key string) ([]dmm.Actress, error) {
	firstPage := actressPage(key, 1)
	pages, err := extractPageCount(firstPage)
	if err != nil {
		return nil, err
	}

	var actresses []dmm.Actress
	for i := 1; i <= pages; i++ {
		page := actressPage(key, i)

		doc, err := goquery.NewDocument(page)
		if err != nil {
			return nil, err
		}

		doc.Find("div.act-box > ul.d-item > li > a").Each(func(i int, s *goquery.Selection) {
			url, ok := s.Attr("href")
			if !ok {
				return
			}

			id, err := extractActressID(url)
			if err != nil {
				return
			}

			img := s.Find("img")
			imgURL, ok := img.Attr("src")
			if !ok {
				return
			}

			name, ok := img.Attr("alt")
			if !ok {
				return
			}

			actress := dmm.Actress{
				ID:    id,
				Name:  name,
				Image: path.Base(imgURL),
			}

			actresses = append(actresses, actress)
		})
	}

	return actresses, nil
}
