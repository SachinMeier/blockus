package bitcoind

import (
	"fmt"
	"time"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
)

func BlockJSONToBlockHeader(blockHeaderJSON *btcjson.GetBlockHeaderVerboseResult) *BlockHeader {
	return &BlockHeader{
		Hash:       blockHeaderJSON.Hash,
		Height:     blockHeaderJSON.Height,
		Version:    blockHeaderJSON.Version,
		MerkleRoot: blockHeaderJSON.MerkleRoot,
		Time:       time.Unix(blockHeaderJSON.Time, 0),
		Nonce:      blockHeaderJSON.Nonce,
		Bits:       blockHeaderJSON.Bits,
		Difficulty: blockHeaderJSON.Difficulty,
		PrevHash:   blockHeaderJSON.PreviousHash,
	}
}

func (rpc *Client) GetBlockHeader(blockHashStr string) (*BlockHeader, error) {
	blockHash, err := chainhash.NewHashFromStr(blockHashStr)
	if err != nil {
		return nil, fmt.Errorf("invalid block hash : %w", err)
	}

	blockJSON, err := rpc.GetBlockHeaderVerbose(blockHash)
	if err != nil {
		return nil, fmt.Errorf("failed to lookup block with bitcoind : %w", err)
	}
	return BlockJSONToBlockHeader(blockJSON), nil
}

func (rpc *Client) GetBlockHeaderByHeight(height int64) (*BlockHeader, error) {
	blockHash, err := rpc.GetBlockHash(height)
	if err != nil {
		return nil, fmt.Errorf("failed to get block hash for block %d : %w", height, err)
	}

	blockJSON, err := rpc.GetBlockHeaderVerbose(blockHash)
	if err != nil {
		return nil, fmt.Errorf("failed to lookup block with bitcoind : %w", err)
	}
	return BlockJSONToBlockHeader(blockJSON), nil
}
