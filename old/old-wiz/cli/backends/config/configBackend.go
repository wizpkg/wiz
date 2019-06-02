package config

import (
	"github.com/tim15/wiz/api/backend"
	"github.com/tim15/wiz/cli/backends/config/json"
	"github.com/tim15/wiz/cli/backends/config/prototxt"
)

func GetBackends() []backend.ConfigBackend {
	return []backend.ConfigBackend{
		json.Register(),
		prototxt.Register(),
	}
}
