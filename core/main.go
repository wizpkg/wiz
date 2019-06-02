package main

import (
	// "fmt"
	"github.com/wizpkg/wizd/src/utils/log"
	"github.com/wizpkg/wizd/src/config"
	"github.com/wizpkg/wizd/src/api/install"
)

var BuildDate string
var BuildTime string

func main() {
	log.Log("Wiz Test v1")
	
	log.Log(config.GetConfig("storeLocation"))
	install.Install("test", "0.1.0")
}
