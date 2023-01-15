package config

type MapProvider struct {
	vals map[string]interface{}
}

func (cfg *MapProvider) GetInt64(keyName string, defaultVal int64) int64 {
	val, exists := cfg.vals[keyName]
	if !exists {
		return defaultVal
	}
	switch t := val.(type) {
	case int64:
		return t
	default:
		return defaultVal
	}
}
func (cfg *MapProvider) GetString(keyName string, defaultVal string) string {
	val, exists := cfg.vals[keyName]
	if !exists {
		return defaultVal
	}
	switch t := val.(type) {
	case string:
		return t
	default:
		return defaultVal
	}
}
func (cfg *MapProvider) GetBool(keyName string, defaultVal bool) bool {
	val, exists := cfg.vals[keyName]
	if !exists {
		return defaultVal
	}
	switch t := val.(type) {
	case bool:
		return t
	default:
		return defaultVal
	}
}
