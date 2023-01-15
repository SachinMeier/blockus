package bitcoind

type Config struct {
	Host       string
	Port       string
	Network    string
	User       string
	Pass       string
	ZMQHost    string
	DisableTLS bool
}

func NewConfig(host string, port string, network string, user string, pass string, zmqHost string, disableTLS bool) Config {
	return Config{
		Host:       host,
		Port:       port,
		User:       user,
		Pass:       pass,
		Network:    network,
		ZMQHost:    zmqHost,
		DisableTLS: disableTLS,
	}
}
