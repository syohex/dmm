package actress

import (
	"testing"
)

func TestKeywords(t *testing.T) {
	keywords := Keywords()
	if len(keywords) != 43 {
		t.Fatal("Invalid keyword length")
	}

	m := make(map[string]int)
	for _, k := range keywords {
		if m[k] != 0 {
			t.Fatalf("Found duplicate keyword: '%s'", k)
		}

		m[k]++
	}
}

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

func TestCollect(t *testing.T) {
	actresses, err := CollectFromKey("wa")
	if err != nil {
		t.Fatalf("Can't extract 'wa' actresses")
	}

	if len(actresses) == 0 {
		t.Fatalf("no actresses")
	}

	found := false
	for _, actress := range actresses {
		if actress.Name == "若槻シェルビー" {
			if actress.Works <= 0 {
				t.Fatalf("Can't extract 'Works' information")
			}
			found = true
		}
	}

	if !found {
		t.Fatalf("parsing actress page is failed")
	}
}
