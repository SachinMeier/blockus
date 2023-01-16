package config

type DefaultProvider struct{}

func NewDefaultProvider() Provider {
	return &DefaultProvider{}
}

func (d *DefaultProvider) GetInt64(keyName string, defaultVal int64) int64 {
	return defaultVal
}
func (d *DefaultProvider) GetString(keyName string, defaultVal string) string {
	return defaultVal
}
func (d *DefaultProvider) GetStrings(keyName string, defaultVals []string) []string {
	return defaultVals
}

func (d *DefaultProvider) GetBool(keyname string, defaultVal bool) bool {
	return defaultVal
}
