package dmm

import (
	"testing"
)

func TestCalculatePages(t *testing.T) {
	actressURL := `http://www.dmm.co.jp/digital/videoa/-/list/=/article=actress/id=1011199/limit=120/sort=release_date/`
	pages, err := CalculatePages(actressURL, 120)
	if err != nil {
		t.Fatal(err)
	}

	if pages < 1 {
		t.Fatalf("Can't get pages: pages=%d", pages)
	}

	t.Logf("id=1011199 actress has %d pages", pages)

	genreURL := `http://www.dmm.co.jp/digital/videoa/-/list/=/article=keyword/id=2001/limit=120/sort=release_date/`
	pages, err = CalculatePages(genreURL, 30)
	if err != nil {
		t.Fatal(err)
	}

	if pages < 10 {
		t.Fatalf("Can't get pages: pages=%d", pages)
	}

	t.Logf("id=2001 genre has %d pages", pages)
}
