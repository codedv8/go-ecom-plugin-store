package main

import (
	EComApp "github.com/codedv8/go-ecom-app"
)

type Store struct {
	App     *EComApp.Application
	Message string
}

var storeList []Store

// Exports
func SysInit(app *EComApp.Application) error {
	store := &Store{
		App:     app,
		Message: "Welcome to the CMS plugin",
	}
	store.SysInit(app)

	storeList = append(storeList, *store)

	return nil
}

func Init(app *EComApp.Application) error {
	for _, store := range storeList {
		store.Init(app)
	}

	return nil
}

func Done(app *EComApp.Application) error {
	for _, store := range storeList {
		store.Done(app)
	}

	return nil
}

var store Store
