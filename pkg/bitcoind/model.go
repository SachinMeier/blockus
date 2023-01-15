package bitcoind

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

const Coin float64 = 100_000_000

type TxInput struct {
	PrevTxid string `json:"prev_txid"`
	PrevVout uint32 `json:"prev_vout"`
	Value    int64  `json:"value"`
}

func NewTxInput(prevTxid string, prevVout uint32, value int64) *TxInput {
	return &TxInput{
		PrevTxid: prevTxid,
		PrevVout: prevVout,
		Value:    value,
	}
}

func (in *TxInput) String() string {
	return fmt.Sprintf("%s:%d", in.PrevTxid, in.PrevVout)
}

func (in *TxInput) GetValue() int64 {
	return in.Value
}

type TxOutput struct {
	Txid         string `json:"txid"`
	Vout         uint32 `json:"vout"`
	ScriptPubkey string `json:"script_pubkey"`
	Address      string `json"address"`
	Value        int64  `json:"value"`
}

func NewTxOutput(txid string, vout uint32, scriptPubkey, address string, value int64) *TxOutput {
	return &TxOutput{
		Txid:         txid,
		Vout:         vout,
		ScriptPubkey: scriptPubkey,
		Address:      address,
		Value:        value,
	}
}

func (out *TxOutput) GetValue() int64 {
	return out.Value
}

type Tx struct {
	Txid      string      `json:"txid"`
	Hex       string      `json:"hex"`
	Version   uint32      `json:"version"`
	Locktime  uint32      `json:"locktime"`
	Size      int32       `json:"size"`
	Weight    int32       `json:"weight"`
	Inputs    []*TxInput  `json:"inputs"`
	Outputs   []*TxOutput `json:"outputs"`
	BlockHash string      `json:"block_hash"`
	BlockTime time.Time   `json:"block_time"`
	// BlockHeight int32       `json:"block_height"`
}

func (tx *Tx) ValueIn() int64 {
	var sum int64 = 0
	for _, in := range tx.Inputs {
		sum += in.Value
	}
	return sum
}

func (tx *Tx) ValueOut() int64 {
	var sum int64 = 0
	for _, out := range tx.Outputs {
		sum += out.Value
	}
	return sum
}

func (tx *Tx) Fee() int64 {
	return tx.ValueIn() - tx.ValueOut()
}

type BlockHeader struct {
	Hash       string    `json:"block_hash"`
	Height     int32     `json:"block_height"`
	Version    int32     `json:"version"`
	MerkleRoot string    `json:"merkle_root"`
	Time       time.Time `json:"block_time"`
	Nonce      uint64    `json:"nonce"`
	Bits       string    `json:"bits"`
	Difficulty float64   `json:"difficulty"`
	PrevHash   string    `json:"prev_hash"`
}

const genesisTargetPadding = 26
const genesisDifficulty int64 = 0x00ffff

// Difficulty uses a block header's bits to calculate the difficulty
// In order to avoid using big.Ints, which do not divide easily into
// floats, we reduce both the genesis target and the current
// target by their GCF with 256.
func (b *BlockHeader) GetDifficultyFromBits(bits string) (float64, error) {
	// TODO: ensure no 0x prefix to bits
	nBits, err := hex.DecodeString(bits)
	if err != nil {
		return -1, fmt.Errorf("invalid bits : %w", err)
	}
	targetPadding := int(nBits[0]) - len(nBits) + 1
	targetPrefix, _ := strconv.ParseInt(bits[2:], 16, 64)
	reducedGenesisPadding := genesisTargetPadding - targetPadding
	reducedGenesisDifficulty := genesisDifficulty << (reducedGenesisPadding * 8)
	return float64(reducedGenesisDifficulty) / float64(targetPrefix), nil
}
