package dmm

import (
	"regexp"
)

// Product is product in dmm.co.jp
type Product struct {
	ID    string
	Title string
	Image string
}

var smallImgRe = regexp.MustCompile(`ps\.jpg$`)

// LargeImage returns large image URL of product
func (p *Product) LargeImage() string {
	return smallImgRe.ReplaceAllLiteralString(p.Image, "pl.jpg")
}
