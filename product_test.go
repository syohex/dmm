package dmm

import (
	"testing"
)

func TestLargeImage(t *testing.T) {
	p := Product{
		Title: "Foo Bar Baz",
		ID:    "c029995030",
		Image: "http://pics.dmm.co.jp/digital/video/tek00067/tek00067ps.jpg",
	}

	large := p.LargeImage()
	if large != "http://pics.dmm.co.jp/digital/video/tek00067/tek00067pl.jpg" {
		t.Fatalf("Can't convert large image URL(got: %s)", large)
	}
}
