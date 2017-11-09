package main

import (
	"github.com/graphql-go/graphql"
)

var ProductObject = graphql.NewObject(graphql.ObjectConfig{
	Name: "Product",
	Fields: graphql.Fields{
		"product_id": &graphql.Field{
			Name: "ProductID",
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				product := p.Source.(Product)
				return product.ProductId, nil
			},
		},
		"product_name": &graphql.Field{
			Name: "ProductName",
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				product := p.Source.(Product)
				return product.ProductName, nil
			},
		},
		"product_price": &graphql.Field{
			Name: "ProductPrice",
			Type: graphql.Float,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				product := p.Source.(Product)
				return product.ProductPrice, nil
			},
		},
		"shop": &graphql.Field{
			Type: ShopObject,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return Shop{
					ShopID:   1,
					ShopName: "Tokopedia",
				}, nil
			},
		},
	},
})

type Product struct {
	ProductId    int
	ProductName  string
	ProductPrice float64
}

func getProduct() Product {
	return Product{
		ProductId:    1,
		ProductName:  "Sepatu",
		ProductPrice: 30000,
	}
}

func getProductByID(productID int) Product {
	return Product{
		ProductId:    1,
		ProductName:  "Sepatu",
		ProductPrice: 30000,
	}
}

func getProducts() []Product {
	return []Product{
		{
			ProductId:    1,
			ProductName:  "Sepatu",
			ProductPrice: 30000,
		},
		{
			ProductId:    2,
			ProductName:  "Baju",
			ProductPrice: 30000,
		},
		{
			ProductId:    3,
			ProductName:  "Celana",
			ProductPrice: 30000,
		},
	}
}
