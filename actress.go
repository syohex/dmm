package dmm

import (
	"fmt"
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

// Actress is actress information in dmm.co.jp
type Actress struct {
	ID    int
	Name  string
	Image string
}

func (a *Actress) pageURL(page int) string {
	url := fmt.Sprintf(`http://www.dmm.co.jp/digital/videoa/-/list/=/article=actress/id=%d/limit=120/sort=release_date/`,
		a.ID)
	if page == 1 {
		return url
	}

	return url + fmt.Sprintf("page=%d/", page)
}

// ImageURL returns actress image URL
func (a *Actress) ImageURL() string {
	return fmt.Sprintf("http://pics.dmm.co.jp/mono/actjpgs/%s", a.Image)
}

var productIDRegexp = regexp.MustCompile("cid=([^/]+)")

// Products returns actress products in dmm.co.jp
func (a *Actress) Products(limit int) ([]Product, error) {
	firstPage := a.pageURL(1)
	pages, err := CalculatePages(firstPage, 120)
	if err != nil {
		return nil, err
	}

	var products []Product
	for i := 1; i < pages; i++ {
		doc, err := goquery.NewDocument(a.pageURL(i))
		if err != nil {
			return nil, err
		}

		doc.Find("p.tmb").Each(func(n int, s *goquery.Selection) {
			a := s.Find("a")
			href, ok := a.Attr("href")
			if !ok {
				return
			}

			matches := productIDRegexp.FindAllStringSubmatch(href, 1)
			if len(matches) == 0 {
				return
			}

			id := matches[0][1]

			img := s.Find("img")
			title, ok := img.Attr("alt")
			if !ok {
				return
			}

			imgSrc, ok := img.Attr("src")
			if !ok {
				return
			}

			product := Product{
				ID:    id,
				Title: title,
				Image: imgSrc,
			}

			products = append(products, product)
		})

		if len(products) > limit {
			break
		}
	}

	return products[0:limit], nil
}
