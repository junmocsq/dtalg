package proxy

type productData struct {
	price     int
	name, sku string
}

var dbMock = make(map[string]productData)

type db struct {
}

func NewDb() *db {
	return &db{}
}

func (d *db) store(p productData) {
	dbMock[p.sku] = p
}

func (d *db) getProductData(sku string) productData {
	return dbMock[sku]
}

func (d *db) deleteProductData(sku string) {
	delete(dbMock, sku)
}
