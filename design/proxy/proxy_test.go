package proxy

import "testing"

func TestProxy(t *testing.T) {
	db := NewDb()
	db.store(productData{sku: "test1", name: "name1", price: 456})
	p := NewProductProxy("test1")
	if p.getName() != "name1" || p.getPrice() != 456 {
		t.Error("product proxy failed!", p.getName(), p.getPrice())
	}
}
