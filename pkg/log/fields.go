package log

import (
	"go.uber.org/zap"
)

const (
	// application
	portKey       = "port"
	httpMethodKey = "http_method"
	httpPathKey   = "http_path"
	httpFromKey   = "http_from"
	httpToKey     = "http_to"

	// general
	wantedKey = "wanted"
	gotKey    = "got"

	// bitcoin
	satoshisKey    = "sats"
	txidKey        = "txid"
	blockHeightKey = "block_height"
)

// application

func Port(port string) zap.Field {
	return zap.String(portKey, port)
}

func HTTPMethod(method string) zap.Field {
	return zap.String(httpMethodKey, method)
}

func HTTPPath(path string) zap.Field {
	return zap.String(httpPathKey, path)
}

func HTTPFrom(from string) zap.Field {
	return zap.String(httpFromKey, from)
}

func HTTPTo(to string) zap.Field {
	return zap.String(httpToKey, to)
}

// General

func Wanted(val interface{}) zap.Field {
	return zap.Any(wantedKey, val)
}

func Got(val interface{}) zap.Field {
	return zap.Any(gotKey, val)
}

// Bitcoin

func Satoshis(amount int64) zap.Field {
	return zap.Int64(satoshisKey, amount)
}

func Txid(txid string) zap.Field {
	return zap.String(txidKey, txid)
}

func BlockHeight(height int64) zap.Field {
	return zap.Int64(blockHeightKey, height)
}
