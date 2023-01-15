package config

import (
	"blockus/pkg/bitcoind"
)

// BITCOIND

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

func LoadBitcoindConfig(cfg Provider) bitcoind.Config {
	host := cfg.GetString(bitcoindHostEnvVar, "127.0.0.1")
	port := cfg.GetString(bitcoindPortEnvVar, "18443")
	network := cfg.GetString(bitcoindNetworkEnvVar, "regtest")
	rpcUser := cfg.GetString(bitcoindRPCUserEnvVar, "bitcoin")
	rpcPass := cfg.GetString(bitcoindRPCPassEnvVar, "password")
	zmqHost := cfg.GetString(bitcoindZMQHostEnvVar, "tcp://127.0.0.1:28333")
	enableTLS := cfg.GetBool(bitcoindEnableTLSEnvVar, false)
	return bitcoind.NewConfig(host, port, network, rpcUser, rpcPass, zmqHost, !enableTLS)
}
