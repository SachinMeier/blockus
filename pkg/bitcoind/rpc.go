package bitcoind

import (
	"fmt"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
)

type RPCClient interface {
	Host() string

	// rpcclient.Client funcs
	Shutdown()
	GetBlockCount() (int64, error)
	BroadcastTransaction(txBytes []byte) (string, error)
	GetRawTransactionVerbose(txid *chainhash.Hash) (*btcjson.TxRawResult, error)

	// addt'l funcs
	RawTransaction(txid string) (*btcjson.TxRawResult, error)
	FullTransaction(txid string) (*Tx, error)
	GetBlockHeader(blockHashStr string) (*BlockHeader, error)
	GetBlockHeaderByHeight(height int64) (*BlockHeader, error)
}

type Client struct {
	*rpcclient.Client

	cfg Config
}

func (rpc *Client) Host() string {
	return fmt.Sprintf("%s:%s", rpc.cfg.Host, rpc.cfg.Port)
}

func NewClient(cfg Config) (*Client, error) {
	client := &Client{
		cfg: cfg,
	}

	connConfig := &rpcclient.ConnConfig{
		Host:         client.Host(),
		User:         cfg.User,
		Pass:         cfg.Pass,
		HTTPPostMode: true,
		Params:       cfg.Network,
		DisableTLS:   cfg.DisableTLS,
	}

	connClient, err := rpcclient.New(connConfig, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create rpc client for bitcoind node : %w", err)
	}
	client.Client = connClient

	return client, nil
}
