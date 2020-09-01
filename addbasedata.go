package main

import (
	"context"
	EComStructs "github.com/codedv8/go-ecom-structs"
	"log"
)

type Page struct {
	Title  string
	SEOURL string
}

type Product struct {
	Title  string
	SEOURL string
}

type Category struct {
	Title  string
	SEOURL string
}

type Brand struct {
	Title  string
	SEOURL string
}

func (store *Store) AddBaseData(payload interface{}) (bool, error) {
	switch c := payload.(type) {
	case *EComStructs.RouterWildcard:
		path := c.Context.Request.URL.Path
		if path == "/addbasedata" {
			log.Println("Adding base data. (removing old data first)")
			store.App.DB.Client.Database("shop").Collection("cms").Drop(context.TODO())
			store.App.DB.Client.Database("shop").Collection("product").Drop(context.TODO())
			store.App.DB.Client.Database("shop").Collection("category").Drop(context.TODO())
			store.App.DB.Client.Database("shop").Collection("brand").Drop(context.TODO())

			store.App.DB.Client.Database("shop").Collection("cms").InsertOne(context.TODO(), Page{Title: "Welcome", SEOURL: "/"})
			store.App.DB.Client.Database("shop").Collection("product").InsertOne(context.TODO(), Product{Title: "Product 1", SEOURL: "/product1"})
			store.App.DB.Client.Database("shop").Collection("product").InsertOne(context.TODO(), Product{Title: "Product 2", SEOURL: "/product2"})
			store.App.DB.Client.Database("shop").Collection("product").InsertOne(context.TODO(), Product{Title: "Product 3", SEOURL: "/product3"})
			store.App.DB.Client.Database("shop").Collection("category").InsertOne(context.TODO(), Category{Title: "Category 1", SEOURL: "/category1"})
			store.App.DB.Client.Database("shop").Collection("category").InsertOne(context.TODO(), Category{Title: "Category 2", SEOURL: "/category2"})
			store.App.DB.Client.Database("shop").Collection("category").InsertOne(context.TODO(), Category{Title: "Category 3", SEOURL: "/category3"})
			store.App.DB.Client.Database("shop").Collection("brand").InsertOne(context.TODO(), Brand{Title: "Brand 1", SEOURL: "/brand1"})
			store.App.DB.Client.Database("shop").Collection("brand").InsertOne(context.TODO(), Brand{Title: "Brand 2", SEOURL: "/brand2"})
			store.App.DB.Client.Database("shop").Collection("brand").InsertOne(context.TODO(), Brand{Title: "Brand 3", SEOURL: "/brand3"})
			c.Context.String(200, "Added some basic data")
			store.Init(store.App)
			return false, nil
		}
		return true, nil
	}
	return true, nil
}
