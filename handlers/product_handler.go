package handlers

import (
	"fmt"
	"koda-b6-backend1/models"
	"github.com/gin-gonic/gin"
)

// Create Product
func CreateProductHandler(ctx *gin.Context) {
	var input models.CreateProductRequest
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(400, models.Response{
			Status:  false,
			Message: "Invalid request body",
		})
		return
	}
	
	response, code := models.CreateComplexProductLogic(input)
	ctx.JSON(code, response)
}

// Get All Products
func GetAllProductsHandler(ctx *gin.Context) {
	ctx.JSON(200, models.Response{
		Status:  true,
		Message: "Success get all products",
		Data:    models.Products,
	})
}

// Get Product By ID
func GetProductByIdHandler(ctx *gin.Context) {
	idParam := ctx.Param("id")
	for _, prod := range models.Products {
		if idParam == fmt.Sprint(prod.IdProduct) {
			ctx.JSON(200, models.Response{
				Status:  true,
				Message: "Product Found!",
				Data:    prod,
			})
			return
		}
	}
	ctx.JSON(404, models.Response{Status: false, Message: "Product not found!"})
}

// Update Product 
func UpdateProductHandler(ctx *gin.Context) {
	idParam := ctx.Param("id")
	var input models.Product

	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(400, models.Response{
			Status:  false,
			Message: "Invalid request body",
		})
		return
	}

	for i := 0; i < len(models.Products); i++ {
		if idParam == fmt.Sprint(models.Products[i].IdProduct) {
			if input.Name != "" {
				models.Products[i].Name = input.Name
			}
			if input.Desc != "" {
				models.Products[i].Desc = input.Desc
			}
			if input.Price > 0 {
				models.Products[i].Price = input.Price
			}
			if input.Quantity >= 0 {
				models.Products[i].Quantity = input.Quantity
			}
			models.Products[i].IsActive = input.IsActive

			ctx.JSON(200, models.Response{
				Status:  true,
				Message: "Product updated successfully!",
				Data:    models.Products[i],
			})
			return
		}
	}
	ctx.JSON(404, models.Response{Status: false, Message: "Product not found"})
}

// Delete Product 
func DeleteProductHandler(ctx *gin.Context) {
	idParam := ctx.Param("id")
	for i := 0; i < len(models.Products); i++ {
		if idParam == fmt.Sprint(models.Products[i].IdProduct) {
			models.Products = append(models.Products[:i], models.Products[i+1:]...)
			ctx.JSON(200, models.Response{
				Status:  true,
				Message: "Product Deleted successfully!",
			})
			return
		}
	}
	ctx.JSON(404, models.Response{Status: false, Message: "Product not found"})
}