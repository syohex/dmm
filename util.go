package dmm

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

var pageRegexp = regexp.MustCompile(`(\d+)タイトル中`)

// CalculatePages returns pages of actress or genre
func CalculatePages(url string, limit int) (int, error) {
	if limit <= 0 {
		return 0, fmt.Errorf("'limit' must be greater than 0(limit=%d)", limit)
	}

	doc, err := goquery.NewDocument(url)
	if err != nil {
		return 0, err
	}

	pageInfo := doc.Find(".list-boxpagenation > p").Text()
	matches := pageRegexp.FindAllStringSubmatch(pageInfo, 1)
	if len(matches) == 0 {
		return 0, fmt.Errorf("can't get page information")
	}

	titles, err := strconv.Atoi(matches[0][1])
	if err != nil {
		return 0, err
	}

	return (titles / limit) + 1, nil
}
