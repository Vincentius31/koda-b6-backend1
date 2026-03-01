package models

// --- Structs based on Database Schema ---

type Product struct {
	IdProduct int    `json:"id_product"`
	Name      string `json:"name"`
	Desc      string `json:"desc"`
	Price     int    `json:"price"`
	Quantity  int    `json:"quantity"`
	IsActive  bool   `json:"is_active"`
}

type ProductCategory struct {
	ProductId  int `json:"product_id"`
	CategoryId int `json:"category_id"`
}

type ProductImage struct {
	IdImage   int    `json:"id_image"`
	ProductId int    `json:"product_id"`
	Path      string `json:"path"`
}

type ProductVariant struct {
	IdVariant       int    `json:"id_variant"`
	ProductId       int    `json:"product_id"`
	VariantName     string `json:"variant_name"`
	AdditionalPrice int    `json:"additional_price"`
}

type ProductSize struct {
	IdSize          int    `json:"id_size"`
	ProductId       int    `json:"product_id"`
	SizeName        string `json:"size_name"`
	AdditionalPrice int    `json:"additional_price"`
}

type Discount struct {
	IdDiscount   int     `json:"id_discount"`
	ProductId    int     `json:"product_id"`
	DiscountRate float64 `json:"discount_rate"`
	Description  string  `json:"description"`
	IsFlashSale  bool    `json:"is_flash_sale"`
}

// --- Struct untuk Request Kompleks (POST) ---

type CreateProductRequest struct {
	Name        string           `json:"name"`
	Desc        string           `json:"desc"`
	Price       int              `json:"price"`
	Quantity    int              `json:"quantity"`
	IsActive    bool             `json:"is_active"`
	CategoryIds []int            `json:"category_ids"`
	Images      []string         `json:"images"` // Path/URLs
	Variants    []ProductVariant `json:"variants"`
	Sizes       []ProductSize    `json:"sizes"`
	Discount    *Discount        `json:"discount"`
}

// --- In-Memory Data Storage (Simulasi DB) ---

var Products []Product
var NextProductId = 1

var ProductCategoryTable []ProductCategory
var ProductImagesTable []ProductImage
var ProductVariantTable []ProductVariant
var ProductSizeTable []ProductSize
var DiscountTable []Discount

// --- Business Logic ---

func CreateComplexProductLogic(input CreateProductRequest) (Response, int) {
	// 1. Validasi Dasar
	if input.Name == "" || input.Price <= 0 {
		return Response{Status: false, Message: "Name and Price are required!"}, 400
	}

	// 2. Simpan Produk Dasar
	newProduct := Product{
		IdProduct: NextProductId,
		Name:      input.Name,
		Desc:      input.Desc,
		Price:     input.Price,
		Quantity:  input.Quantity,
		IsActive:  input.IsActive,
	}
	Products = append(Products, newProduct)
	
	// Simulasi ID dari DB
	currentId := NextProductId
	NextProductId++

	// 3. Simpan Relasi (Simulasi Transaksi)

	// Relasi Category (Many-to-Many)
	for _, catId := range input.CategoryIds {
		ProductCategoryTable = append(ProductCategoryTable, ProductCategory{
			ProductId:  currentId,
			CategoryId: catId,
		})
	}

	// Relasi Images (One-to-Many)
	for _, imgPath := range input.Images {
		ProductImagesTable = append(ProductImagesTable, ProductImage{
			ProductId: currentId,
			Path:      imgPath,
		})
	}

	// Relasi Variants (One-to-Many)
	for _, v := range input.Variants {
		v.ProductId = currentId
		ProductVariantTable = append(ProductVariantTable, v)
	}

	// Relasi Sizes (One-to-Many)
	for _, s := range input.Sizes {
		s.ProductId = currentId
		ProductSizeTable = append(ProductSizeTable, s)
	}

	// Relasi Discount (One-to-One)
	if input.Discount != nil {
		input.Discount.ProductId = currentId
		DiscountTable = append(DiscountTable, *input.Discount)
	}

	return Response{
		Status:  true,
		Message: "Product and all relations created successfully",
		Data:    newProduct,
	}, 201
}