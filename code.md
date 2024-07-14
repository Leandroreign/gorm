Driver := storage.Postgres
storage.New(Driver)

var products []model.Product
storage.DB().Find(&products)

for _, product := range products {
    fmt.Printf("%+v\n", product)
}

product := model.Product{}
storage.DB().First(&product, 2)

fmt.Printf("\nImprimiendo el producto con el id %v\n%+v", product.ID, product)

product.Name = "producto modificado"

storage.DB().Save(&product)

storage.DB().Model(&product).Updates(model.Product{Name: "modificado con el update"})

storage.DB().Delete(&product)