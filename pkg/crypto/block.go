package crypto

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type Block struct {
	timestamp time.Time
	data      []byte
	hash      [32]byte
	prevBlock *Block
}

func (s *Block) GetHash() [32]byte {
	var prevHash [32]byte
	if s.prevBlock != nil {
		b := s.prevBlock
		prevHash = b.hash
	}

	return sha256.Sum256([]byte(fmt.Sprintf("%s%s%s", prevHash, s.timestamp, string(s.data))))
}

func NewBlock(data []byte, timestamp time.Time) Block {
	return Block{
		timestamp: timestamp,
		data:      data,
	}
}
