package registry

import "github.com/wizpkg/wizd/src/core/config"

func FindPackage(pkg string) {
	r := config.GetConfig("registryLocation")
	l := r + pkg
}
