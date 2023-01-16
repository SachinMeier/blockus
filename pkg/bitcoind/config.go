package bitcoind

import "github.com/SachinMeier/blockus/pkg/config"

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

// Environment variable names for bitcoind Config values
const (
	bitcoindHostEnvVar      = "BITCOIND_HOST"
	bitcoindPortEnvVar      = "BITCOIND_PORT"
	bitcoindNetworkEnvVar   = "BITCOIN_NETWORK"
	bitcoindRPCUserEnvVar   = "BITCOIND_RPC_USER"
	bitcoindRPCPassEnvVar   = "BITCOIND_RPC_PASS"
	bitcoindZMQHostEnvVar   = "BITCOIND_ZMQ_HOST"
	bitcoindEnableTLSEnvVar = "BITCOIND_ENABLE_TLS"
)

func LoadConfig(cfg config.Provider) Config {
	host := cfg.GetString(bitcoindHostEnvVar, "127.0.0.1")
	port := cfg.GetString(bitcoindPortEnvVar, "18443")
	network := cfg.GetString(bitcoindNetworkEnvVar, "regtest")
	rpcUser := cfg.GetString(bitcoindRPCUserEnvVar, "bitcoin")
	rpcPass := cfg.GetString(bitcoindRPCPassEnvVar, "password")
	zmqHost := cfg.GetString(bitcoindZMQHostEnvVar, "tcp://127.0.0.1:28333")
	enableTLS := cfg.GetBool(bitcoindEnableTLSEnvVar, false)
	return NewConfig(host, port, network, rpcUser, rpcPass, zmqHost, !enableTLS)
}
