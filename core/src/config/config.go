package config

// GetConfig returns the configuration for the value provided
func GetConfig(key string) (value string) {
	defaultConfig := map[string]string{
		"storeLocation": "~/wiz",
	}
	return defaultConfig[key]
}
