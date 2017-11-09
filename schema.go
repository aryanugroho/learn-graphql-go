package main

import (
	"github.com/graphql-go/graphql"
	"log"
)

func CreateSchema() graphql.Schema {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Root",
			Fields: graphql.Fields{
				"get_product": &graphql.Field{
					Name: "GetProduct",
					Type: ProductObject,
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						product := getProduct()
						return product, nil
					},
				},
				"get_product_by_id": &graphql.Field{
					Name: "GetProductById",
					Type: ProductObject,
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						productID := p.Args["product_id"].(int)
						product := getProductByID(productID)
						return product, nil
					},
				},
				"get_products": &graphql.Field{
					Name: "GetProducts",
					Type: graphql.NewList(ProductObject),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						product := getProducts()
						return product, nil
					},
				},
			},
		}),
	})

	if err != nil {
		log.Fatal(err)
	}

	return schema

}
