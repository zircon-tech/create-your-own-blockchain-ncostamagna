package crypto

import (
	"time"
	"crypto/sha256"
)

interface Block { }

type block struct {
	timestamp *time.Time
	data []byte
	hash []byte
	//prevHash []byte
	prevBlock Block
}

func (s *block) generateHash(){
	prevHash := ""
	if s.prevBlock != nil {
		prevHash = s.prevBlock.hash
	}
	s.hash = sha256.Sum256([]byte(fmt.Sprintf("%s%s%s",prevHash,s.timestamp,string(s.data))))
}

func New(data []byte, timestamp *time.Time, prev Block) Block {
	return &block{
		timestamp: timestamp,
		data: data,
		prevBlock: prev,
	}
}