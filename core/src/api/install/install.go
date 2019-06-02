package install

import (
	"fmt"
	"os"

	"github.com/wizpkg/wizd/src/config"
	"github.com/wizpkg/wizd/src/utils/log"
)

// Install a package to the specified store
func Install(packageName string, packageVersion string) string {
	store := config.GetConfig("storeLocation")

	log.Log("Installing to store: ", store)
	log.Log(packageName, packageVersion)

	path := store + "/packages/" + packageName + "/" + packageVersion
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println(path, "does not exist")
	} else {
		fmt.Println("path exists")
	}
	return packageName + packageVersion
}
