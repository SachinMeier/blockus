package bitcoind

import (
	"bytes"
	"fmt"
	"time"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
)

func (rpc *Client) RawTransaction(txid string) (*btcjson.TxRawResult, error) {
	txID, err := chainhash.NewHashFromStr(txid)
	if err != nil {
		return nil, fmt.Errorf("invalid txid : %w", err)
	}
	tx, err := rpc.GetRawTransactionVerbose(txID)
	if err != nil {
		return nil, fmt.Errorf("failed to get transaction from bitcoind : %w", err)
	}
	return tx, nil
}

func (rpc *Client) FullTransaction(txidStr string) (*Tx, error) {
	txid, err := chainhash.NewHashFromStr(txidStr)
	if err != nil {
		return nil, fmt.Errorf("invalid txid : %w", err)
	}

	txJSON, err := rpc.GetRawTransactionVerbose(txid)
	if err != nil {
		return nil, fmt.Errorf("failed to get transaction from bitcoind : %w", err)
	}

	// query all inputs for valueIn and
	// create TxInput objects for each
	txInputs := make([]*TxInput, 0, len(txJSON.Vin))
	for _, in := range txJSON.Vin {
		prevTxid, err := chainhash.NewHashFromStr(in.Txid)
		if err != nil {
			return nil, fmt.Errorf("invalid previous txid for input : %w", err)
		}
		prevTx, err := rpc.GetRawTransactionVerbose(prevTxid)
		if err != nil {
			return nil, fmt.Errorf("failed to get previous transaction from bitcoind : %w", err)
		}
		input := NewTxInput(in.Txid, in.Vout, int64(prevTx.Vout[in.Vout].Value*Coin))
		txInputs = append(txInputs, input)
	}

	// map btcjson.Vout to TxOutput
	txOutputs := make([]*TxOutput, 0, len(txJSON.Vout))
	for _, out := range txJSON.Vout {
		output := NewTxOutput(txidStr, out.N, out.ScriptPubKey.Hex, out.ScriptPubKey.Addresses[0], int64(out.Value*Coin))
		txOutputs = append(txOutputs, output)
	}

	return &Tx{
		Txid:      txidStr,
		Hex:       txJSON.Hex,
		Version:   txJSON.Version,
		Locktime:  txJSON.LockTime,
		Size:      txJSON.Size,
		Weight:    txJSON.Weight,
		Inputs:    txInputs,
		Outputs:   txOutputs,
		BlockHash: txJSON.BlockHash,
		BlockTime: time.Unix(txJSON.Blocktime, 0),
	}, nil
}

func (rpc *Client) BroadcastTransaction(txBytes []byte) (string, error) {
	msgTx := wire.NewMsgTx(1)
	err := msgTx.Deserialize(bytes.NewBuffer(txBytes))
	if err != nil {
		return "", fmt.Errorf("failed to deserialize tx %X : %w", txBytes, err)
	}
	txid, err := rpc.SendRawTransaction(msgTx, true)
	if err != nil {
		return "", fmt.Errorf("failed to broadcast tx %X : %w", txBytes, err)
	}
	return txid.String(), nil
}
