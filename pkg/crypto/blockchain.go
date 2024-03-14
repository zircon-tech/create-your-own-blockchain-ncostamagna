package crypto

type BlockChain struct {
	chain []Block
}

func NewBC() *BlockChain {
	return &BlockChain{}
}

func (bc *BlockChain) Add(b Block) *BlockChain {

	if lb := bc.GetLastBlock(); lb != nil {
		b.prevBlock = lb
	}

	b.hash = b.GetHash()
	bc.chain = append(bc.chain, b)

	return bc
}

func (bc *BlockChain) GetLastBlock() *Block {

	if len(bc.chain) < 1 {
		return nil
	}

	return &bc.chain[len(bc.chain)-1]
}

func (bc *BlockChain) IsValid() bool {

	for i := range bc.chain {

		if bc.chain[i].hash != bc.chain[i].GetHash() {
			return false
		}

		if i == 0 {
			continue
		}

		if bc.chain[i-1].hash != bc.chain[i].prevBlock.hash {
			return false
		}
	}

	return true

}
