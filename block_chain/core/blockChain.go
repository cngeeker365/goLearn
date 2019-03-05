package core

import (
	"encoding/json"
	"log"
)

type BlockChain struct {
	Blocks 	[]*Block
}

func (b *BlockChain) Append(block *Block){
	if len(b.Blocks)==0 {
		b.Blocks = append(b.Blocks, block)
		return
	}
	if isValid(*b.Blocks[len(b.Blocks)-1], *block){
		b.Blocks = append(b.Blocks, block)
	}else {
		log.Printf("invalid block")
	}
}

func isValid(old Block, new Block) bool{
	if old.Index+1 != new.Index {
		return false
	}
	if old.Hash != new.PreBlockHash {
		return false
	}
	if new.Hash != new.CalHash() {
		return false
	}
	return true
}

func (b *BlockChain) SendData(data string) {
	preBlock := b.Blocks[len(b.Blocks)-1]
	newBlock := NewBlock(*preBlock, data)
	b.Append(&newBlock)
}

func NewBlockChain() *BlockChain{
	genesisBlock := NewGenesisBlock()
	blockChain := BlockChain{}
	blockChain.Append(&genesisBlock)
	return &blockChain
}

func (b *BlockChain) String() string{
	data, err := json.MarshalIndent(b,"", "\t")
	if err != nil {
		log.Printf("Error occured on Json Marshalling: %v\n", err)
	}
	return string(data)
}