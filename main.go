package main

import (
	"fmt"

	"github.com/Leandroreign/gorm/model"
	"github.com/Leandroreign/gorm/storage"
)

func main() {
	driver := storage.Postgres
	storage.New(driver)
	storage.DB().AutoMigrate(
		&model.Product{},
		&model.InvoiceHeader{},
		&model.InvoiceItem{},
	)

	storage.DB().Model(&model.InvoiceItem{}).AddForeignKey("invoice_id",
		"invoice_headers(id)", "RESTRICT", "RESTRICT")

	storage.DB().Model(&model.InvoiceItem{}).AddForeignKey("product_id",
		"products(id)", "RESTRICT", "RESTRICT")

	var Products []model.Product
	storage.DB().Find(&Products)

	var invoiceItems []model.InvoiceItem
	invoiceItems = append(invoiceItems,
		model.InvoiceItem{
			ProductID: Products[0].ID,
			Quantity:  5,
		},
		model.InvoiceItem{
			ProductID: Products[1].ID,
			Quantity:  8,
		},
	)

	InvoiceHeader := model.InvoiceHeader{
		Client:       "Leandro",
		InvoiceItems: invoiceItems,
	}

	storage.DB().Create(&InvoiceHeader)
	fmt.Printf("%+v", InvoiceHeader)
}
