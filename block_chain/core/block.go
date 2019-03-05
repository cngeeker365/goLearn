package core

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Index 			int64	//区块编号，标识该区块在链中的位置，方便查找
	Timestamp 		int64	//区块创建的时间戳
	PreBlockHash	string	//上一个区块的哈希值
	Hash 			string	//本区块的哈希值
	Data 			string	//区块数据（简单模拟，实际包含诸多交易数据）
}

func (b Block) CalHash() string{
	blockData := string(b.Index)+string(b.Timestamp)+string(b.Data)
	hashBytes := sha256.Sum256([]byte(blockData))
	return hex.EncodeToString(hashBytes[:])
}

func NewBlock(preBlock Block, data string) Block{
	newBlock := Block{
		Index: preBlock.Index+1,
		PreBlockHash: preBlock.Hash,
		Timestamp:time.Now().Unix(),
		Data:data,
	}
	newBlock.Hash = newBlock.CalHash()
	return newBlock
}

func NewGenesisBlock() Block{
	preBlock := Block{
		Index: -1,
		Hash: "",
	}
	return NewBlock(preBlock, "Genesis Block")
}

