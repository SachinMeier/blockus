package bitcoind

import (
	"context"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/SachinMeier/blockus/pkg/log"
	zmq "github.com/pebbe/zmq4"
)

const (
	TopicRawTx     string = "rawtx"
	TopicRawBlock  string = "rawblock"
	TopicHashTx    string = "hashtx"
	TopicHashBlock string = "hashblock"
	TopicSequence  string = "sequence"
)

type ZMQMsg interface {
	Topic() string
	Data() string
}

type ZMQMsgParser func([][]byte) (ZMQMsg, error)

// NewZMQSubscription returns a new ZMQ subscription with the passed msgFilter filter
func (c *Client) NewZMQSubscription(msgFilter string) (*zmq.Socket, error) {
	notifs, err := zmq.NewSocket(zmq.SUB)
	if err != nil {
		return nil, fmt.Errorf("failed to create zmq subscription : %w", err)
	}
	err = notifs.SetSubscribe(msgFilter)
	if err != nil {
		return nil, fmt.Errorf("failed to create zmq subcription : %w", err)
	}
	err = notifs.Connect(c.cfg.ZMQHost)
	if err != nil {
		return nil, fmt.Errorf("failed to create zmq subcription : %w", err)
	}
	return notifs, nil
}

// ListenForZMQNotifications is a long-running routine to listen to ZMQ notifications and handle them appropriately
func ListenForZMQNotifications(ctx context.Context, socket *zmq.Socket, handle func(topic string, msg ZMQMsg) error, handleErr func(err error)) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			recvZMQMessage(socket, handle, handleErr)
		}
	}
}

func recvZMQMessage(socket *zmq.Socket, handle func(topic string, msg ZMQMsg) error, handleErr func(err error)) {
	msg, err := socket.RecvMessageBytes(0)
	if err != nil {
		log.Errorf("error listening to zmq : %v", err)
	}
	topic := string(msg[0])

	var parser ZMQMsgParser

	switch topic {
	case TopicRawTx:
		parser = ParseRawTxMsg
	case TopicRawBlock:
		parser = ParseRawBlockMsg
	case TopicHashTx:
		parser = ParseHashTxMsg
	case TopicHashBlock:
		parser = ParseHashBlockMsg
	case TopicSequence:
		parser = ParseSequenceMsg
	default:
		handleErr(fmt.Errorf("unknown topic for zmq message: %s with %d parts", topic, len(msg)))
		return
	}

	zmqMsg, err := parser(msg[1:])
	if err != nil {
		handleErr(fmt.Errorf("failed to parse zmq msg : %w", err))
		return
	}
	if err := handle(topic, zmqMsg); err != nil {
		handleErr(err)
	}
}

type RawTxMsg struct {
	TxHex    string
	Sequence int64
}

func (msg *RawTxMsg) Topic() string {
	return TopicRawTx
}

func (msg *RawTxMsg) Data() string {
	return msg.TxHex
}

// NewZMQRawTxSubscription returns a new ZMQ subscription to raw txs
func (c *Client) NewZMQRawTxSubscription() (*zmq.Socket, error) {
	return c.NewZMQSubscription(TopicRawTx)
}

func ParseRawTxMsg(parts [][]byte) (ZMQMsg, error) {
	// 2 parts: tx hex & sequence
	partCt := len(parts)
	if partCt != 2 {
		return nil, fmt.Errorf("message has unexpected number of parts: %d", partCt)
	}
	txHex := hex.EncodeToString(parts[0])
	seq := int64(binary.LittleEndian.Uint32(parts[1]))
	return &RawTxMsg{
		TxHex:    txHex,
		Sequence: seq,
	}, nil
}

type RawBlockMsg struct {
	BlockHex string
	Sequence int64
}

func (msg *RawBlockMsg) Topic() string {
	return TopicRawBlock
}

func (msg *RawBlockMsg) Data() string {
	return msg.BlockHex
}

// NewZMQRawBlockSubscription returns a ZMQ subscription to raw blocks
func (c *Client) NewZMQRawBlockSubscription() (*zmq.Socket, error) {
	return c.NewZMQSubscription(TopicRawBlock)
}

func ParseRawBlockMsg(parts [][]byte) (ZMQMsg, error) {
	// 2 parts: block hex & sequence
	partCt := len(parts)
	if partCt != 2 {
		return nil, fmt.Errorf("message has unexpected number of parts: %d", partCt)
	}
	block := hex.EncodeToString(parts[0])
	seq := int64(binary.LittleEndian.Uint32(parts[1]))
	return &RawBlockMsg{
		BlockHex: block,
		Sequence: seq,
	}, nil
}

type HashTxMsg struct {
	TxHash   string
	Sequence int64
}

func (msg *HashTxMsg) Topic() string {
	return TopicHashTx
}

func (msg *HashTxMsg) Data() string {
	return msg.TxHash
}

// NewZMQHashTxSubscription returns a new ZMQ subscription to tx hashes
func (c *Client) NewZMQHashTxSubscription() (*zmq.Socket, error) {
	return c.NewZMQSubscription(TopicHashTx)
}

func ParseHashTxMsg(parts [][]byte) (ZMQMsg, error) {
	// 2 parts: block hex & sequence
	partCt := len(parts)
	if partCt != 2 {
		return nil, fmt.Errorf("message has unexpected number of parts: %d", partCt)
	}
	txid := hex.EncodeToString(reverseBytes(parts[0]))
	seq := int64(binary.LittleEndian.Uint32(parts[1]))
	return &HashTxMsg{
		TxHash:   txid,
		Sequence: seq,
	}, nil
}

type HashBlockMsg struct {
	BlockHash string
	Sequence  int64
}

func (msg *HashBlockMsg) Topic() string {
	return TopicHashBlock
}

func (msg *HashBlockMsg) Data() string {
	return msg.BlockHash
}

// NewZMQHashBlockSubscription returns a ZMQ subscription to block hashes
func (c *Client) NewZMQHashBlockSubscription() (*zmq.Socket, error) {
	return c.NewZMQSubscription(TopicHashBlock)
}

// ParseHashBlockMsg parses a HashBlockMsg
func ParseHashBlockMsg(parts [][]byte) (ZMQMsg, error) {
	// 2 parts: block hex & sequence
	partCt := len(parts)
	if partCt != 2 {
		return nil, fmt.Errorf("message has unexpected number of parts: %d", partCt)
	}
	blockHash := hex.EncodeToString(reverseBytes(parts[0]))
	seq := int64(binary.LittleEndian.Uint32(parts[1]))
	return &HashBlockMsg{
		BlockHash: blockHash,
		Sequence:  seq,
	}, nil
}

type SequenceMsg struct {
	Hash            string
	Type            SequenceType
	MempoolSequence uint64
}

type SequenceType string

const (
	SequenceTypeBlockConnect    SequenceType = "C"
	SequenceTypeBlockDisconnect SequenceType = "D"
	SequenceTypeTxRemove        SequenceType = "R"
	SequenceTypeTxAdd           SequenceType = "A"
)

var validSequenceTypes = []SequenceType{
	SequenceTypeBlockConnect,
	SequenceTypeBlockDisconnect,
	SequenceTypeTxRemove,
	SequenceTypeTxAdd,
}

func (t SequenceType) Name() string {
	switch t {
	case SequenceTypeBlockConnect:
		return "block connected"
	case SequenceTypeBlockDisconnect:
		return "block disconnected"
	case SequenceTypeTxRemove:
		return "tx removed from mempool"
	case SequenceTypeTxAdd:
		return "tx added to mempool"
	default:
		return "invalid sequence event"
	}
}

func (t SequenceType) Valid() bool {
	for _, valid := range validSequenceTypes {
		if t == valid {
			return true
		}
	}
	return false
}

func (msg *SequenceMsg) Topic() string {
	return TopicSequence
}

func (msg *SequenceMsg) Data() string {
	if msg.MempoolSequence != 0 {
		return fmt.Sprintf("%s: %s (%d)", msg.Type.Name(), msg.Hash, msg.MempoolSequence)
	}
	return fmt.Sprintf("%s: %s", msg.Type.Name(), msg.Hash)
}

// NewZMQSequenceSubscription returns a new ZMQ subscription to sequence events
func (c *Client) NewZMQSequenceSubscription() (*zmq.Socket, error) {
	return c.NewZMQSubscription(TopicSequence)
}

func ParseSequenceMsg(parts [][]byte) (ZMQMsg, error) {
	// 3 parts: hash, type, sequence?
	var seq uint64
	partCt := len(parts)
	if partCt < 2 || partCt > 4 {
		return nil, fmt.Errorf("message has unexpected number of parts: %d", partCt)
	}
	hash := hex.EncodeToString(parts[0][:32])
	t := SequenceType(parts[0][32])
	if !t.Valid() {
		return nil, fmt.Errorf("sequence mesage has unrecognized event type: %s | %s | %d", hash, t, seq)
	}

	if partCt == 3 {
		seq = binary.LittleEndian.Uint64(parts[2])
	}
	return &SequenceMsg{
		Hash:            hash,
		Type:            t,
		MempoolSequence: seq,
	}, nil
}
