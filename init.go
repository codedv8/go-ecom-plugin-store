package main

import (
	EComApp "github.com/codedv8/go-ecom-app"
	EComStructs "github.com/codedv8/go-ecom-structs"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
)

func (store *Store) SysInit(app *EComApp.Application) {
}

func (store *Store) Init(app *EComApp.Application) {
	app.URIHandler.AddURI("/cart", func(ctx *gin.Context) (bool, error) {
		ctx.String(200, "You have requested the cart.")
		return true, nil
	})
	app.ListenToHook("ROUTER_WILDCARD", func(payload interface{}) (bool, error) {
		switch c := payload.(type) {
		case *EComStructs.RouterWildcard:
			path := c.Context.Request.URL.Path
			if path == "/cart" {
				c.Context.String(200, "Plugin Store responded to this call: "+path)
				return false, nil
			}
			return true, nil
		}
		return true, nil
	})
}
