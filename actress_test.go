package dmm

import (
	"testing"
)

func TestActressProduct(t *testing.T) {
	actress := Actress{
		ID: 26225,
	}

	products, err := actress.Products(42)
	if err != nil {
		t.Fatal(err)
	}

	if len(products) != 42 {
		t.Fatalf("invalid return products number(got: %d)", len(products))
	}

	t.Log(products)
}
