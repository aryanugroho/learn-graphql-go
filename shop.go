package main

import "github.com/graphql-go/graphql"

type Shop struct {
	ShopID   int    `json:"shop_id"`
	ShopName string `json:"shop_name"`
}

var ShopObject = createShopObj()

func createShopObj() *graphql.Object {
	fields := graphql.BindFields(Shop{})

	return graphql.NewObject(graphql.ObjectConfig{
		Name:   "Shop",
		Fields: fields,
	})
}
