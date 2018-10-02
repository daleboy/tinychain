package types

import (
	"encoding/json"
	"github.com/tinychain/tinychain/common"
)

// Log represents a contract log event. These events are generated by the LOG opcode and
// stored/indexed by the node.
type Log struct {
	// Consensus fields:
	// address of the contract that generated the event
	Address common.Address `json:"address" gencodec:"required"`
	// list of topics provided by the contract.
	Topics []common.Hash `json:"topics" gencodec:"required"`
	// supplied by the contract, usually ABI-encoded
	Data []byte `json:"data" gencodec:"required"`

	// Derived fields. These fields are filled in by the node
	// but not secured by consensus.
	// block in which the transaction was included
	BlockHeight uint64 `json:"blockHeight"`
	// hash of the transaction
	TxHash common.Hash `json:"transactionHash" gencodec:"required"`
	// index of the transaction in the block
	TxIndex uint `json:"transactionIndex" gencodec:"required"`
	// hash of the block in which the transaction was included
	BlockHash common.Hash `json:"blockHash"`
	// index of the log in the receipt
	Index uint `json:"logIndex" gencodec:"required"`

	// The Removed field is true if this log was reverted due to a chain reorganisation.
	// You must pay attention to this field if you receive logs through a filter query.
	Removed bool `json:"removed"`
}

func (log *Log) Serialize() ([]byte, error) {
	return json.Marshal(log)
}

func (log *Log) Deserialize(d []byte) error {
	return json.Unmarshal(d, log)
}

type Logs []*Log

func (logs Logs) Serialize() ([]byte, error) {
	return json.Marshal(logs)
}

func (logs Logs) Deserialize(d []byte) error {
	return json.Unmarshal(d, logs)
}
