package main

import (
	"io/ioutil"
	"os"

	"github.com/gostack/views"
)

const (
	// assetsPath represents the location of templates, default to assets/templates
	// folder in the current directory.
	templatesPath = "_assets/templates"

	// cacheTemplates toggle template caching, should only be enable on prod.
	cacheTemplates = false
)

// viewsManager handles view templates and rendering with layouts.
var viewsManager = views.NewManager(templatesPath, loadAsset, cacheTemplates)

// loadAsset loads the current asset from disk
func loadAsset(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(f)
}
