package config

type Provider interface {
	GetInt64(keyName string, defaultVal int64) int64
	GetString(keyName string, defaultVal string) string
	GetBool(keyname string, defaultVal bool) bool
}
