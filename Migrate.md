
## Migrate tables
```go
driver := storage.Postgres
storage.New(driver)
storage.DB().AutoMigrate(
    &model.Product{},
    &model.InvoiceHeader{},
    &model.InvoiceItem{},
)
```

## Adding Foreign keys
```go
driver := storage.Postgres
storage.New(driver)

storage.DB().Model(&model.InvoiceItem{}).AddForeignKey("invoice_header_id",
    "invoice_headers(id)", "RESTRICT", "RESTRICT")

storage.DB().Model(&model.InvoiceItem{}).AddForeignKey("product_id",
    "products(id)", "RESTRICT", "RESTRICT")
```