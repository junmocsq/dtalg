package proxy

// 数据库访问和业务规则分离
type product interface {
	getPrice() int
	getName() string
	getSku() string
}

type productImp struct {
	itsPrice        int
	itsName, itsSku string
}

func NewProductImp(price int, name, sku string) product {
	return &productImp{itsPrice: price, itsName: name, itsSku: sku}
}

func (p *productImp) getPrice() int {
	// 业务规则。。。。

	return p.itsPrice
}

func (p *productImp) getName() string {
	return p.itsName
}

func (p *productImp) getSku() string {
	return p.itsSku
}

type productProxy struct {
	itsSku string
}

func NewProductProxy(sku string) product {
	return &productProxy{itsSku: sku}
}

func (p *productProxy) getPrice() int {
	d := NewDb().getProductData(p.itsSku)
	// 代理使用另一个对象
	i := NewProductImp(d.price, d.name, d.sku)
	return i.getPrice()
}

func (p *productProxy) getName() string {
	d := NewDb().getProductData(p.itsSku)
	i := NewProductImp(d.price, d.name, d.sku)
	return i.getName()
}

func (p *productProxy) getSku() string {
	return p.itsSku
}
