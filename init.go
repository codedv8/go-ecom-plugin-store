package main

import (
	"context"

	ecomapp "github.com/codedv8/go-ecom-app"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// SysInit - Pre initialization of this object
func (store *Store) SysInit(app *ecomapp.Application) {
}

// SEOURL - A generic struct with the key SEOURL to be able to read seourls from the database
type SEOURL struct {
	SEOURL string `bson:"seourl"`
}

// Init - Initialization of this object
func (store *Store) Init(app *ecomapp.Application) {
	// app.URIHandler.AddURI("/cart", store.Cart)
	app.URIHandler.AddURI("/__addtocart", store.AddToCart)
	app.URIHandler.AddURI("/__removefromcart", store.RemoveFromCart)
	app.URIHandler.AddURI("/__emptycart", store.EmptyCart)

	// Add base data
	app.ListenToHook("ROUTER_WILDCARD", store.AddBaseData)

	// Load seourl's from product
	productCollection := app.DB.Client.Database("shop").Collection("product")
	productCursor, productErr := productCollection.Find(context.TODO(), bson.M{})
	if productErr == nil {
		for productCursor.Next(context.TODO()) {
			var seourl SEOURL
			productCursor.Decode(&seourl)
			if len(seourl.SEOURL) > 0 {
				app.URIHandler.AddURI(seourl.SEOURL, store.HandleProduct)
			}
		}
	}

	// Load seourl's from category
	categoryCollection := app.DB.Client.Database("shop").Collection("category")
	categoryCursor, categoryErr := categoryCollection.Find(context.TODO(), bson.M{})
	if categoryErr == nil {
		for categoryCursor.Next(context.TODO()) {
			var seourl SEOURL
			categoryCursor.Decode(&seourl)
			if len(seourl.SEOURL) > 0 {
				app.URIHandler.AddURI(seourl.SEOURL, store.HandleCategory)
			}
		}
	}

	// Load seourl's from brand
	brandCollection := app.DB.Client.Database("shop").Collection("brand")
	brandCursor, brandErr := brandCollection.Find(context.TODO(), bson.M{})
	if brandErr == nil {
		for brandCursor.Next(context.TODO()) {
			var seourl SEOURL
			brandCursor.Decode(&seourl)
			if len(seourl.SEOURL) > 0 {
				app.URIHandler.AddURI(seourl.SEOURL, store.HandleBrand)
			}
		}
	}
}

// Cart - Handling requests to cart
func (store *Store) Cart(ctx *gin.Context) (bool, error) {
	ctx.String(200, "You have requested the cart.")
	return true, nil
}

// AddToCart - Handling adding items to the cart
func (store *Store) AddToCart(ctx *gin.Context) (bool, error) {
	ctx.String(200, "You are trying to add items to the cart.")
	return true, nil
}

// RemoveFromCart - Handling removing items from the cart
func (store *Store) RemoveFromCart(ctx *gin.Context) (bool, error) {
	ctx.String(200, "You are trying to remove items from the cart.")
	return true, nil
}

// EmptyCart - Handling emptying the the cart
func (store *Store) EmptyCart(ctx *gin.Context) (bool, error) {
	ctx.String(200, "You are trying empty the cart.")
	return true, nil
}

// HandleProduct - Handling a product
func (store *Store) HandleProduct(ctx *gin.Context) (bool, error) {
	ctx.String(200, "You are trying access a product.")
	return true, nil
}

// HandleCategory - Handling a category
func (store *Store) HandleCategory(ctx *gin.Context) (bool, error) {
	ctx.String(200, "You are trying access a category.")
	return true, nil
}

// HandleBrand - Handling a brand
func (store *Store) HandleBrand(ctx *gin.Context) (bool, error) {
	ctx.String(200, "You are trying access a brand.")
	return true, nil
}
