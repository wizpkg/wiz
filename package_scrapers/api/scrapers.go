package api

import (
	utils "github.com/alexkreidler/goutils"
)

type ListPackagesFunction func() <-chan utils.ConcurrentSliceItem

type Scraper interface {
	ListPackages() <-chan utils.ConcurrentSliceItem
	GetPackage(string) interface{}
}
