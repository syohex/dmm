package actress

import (
	"testing"
)

func TestCollectActressPageInfo(t *testing.T) {
	url := "http://www.dmm.co.jp/digital/videoa/-/actress/=/keyword=a/"
	pages, err := extractPageCount(url)
	if err != nil {
		t.Fatalf("Can't get page information from %s", url)
	}

	if pages <= 0 {
		t.Fatalf("Page should be larger than 0")
	}
}

func TestExtractActresses(t *testing.T) {
	actresses, err := ExtractActresses("wa")
	if err != nil {
		t.Fatalf("Can't extract 'wa' actresses")
	}

	t.Logf("@@ %v", actresses)
}