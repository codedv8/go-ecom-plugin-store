package main

import (
	ecomapp "github.com/codedv8/go-ecom-app"
)

// Store - The main struct for this module
type Store struct {
	App     *ecomapp.Application
	Message string
}

var storeList []Store

// Exports

// SysInit - Pre initialization of this module
func SysInit(app *ecomapp.Application) error {
	store := &Store{
		App:     app,
		Message: "Welcome to the CMS plugin",
	}
	store.SysInit(app)

	storeList = append(storeList, *store)

	return nil
}

// Init - Initialization of this module
func Init(app *ecomapp.Application) error {
	for _, store := range storeList {
		store.Init(app)
	}

	return nil
}

// Done - Shut down of this module
func Done(app *ecomapp.Application) error {
	for _, store := range storeList {
		store.Done(app)
	}

	return nil
}

var store Store
