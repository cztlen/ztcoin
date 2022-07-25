package block

import (
	"crypto/sha256"
	"encoding/hex"
)

//创建区块

type Block struct {
	Data     string
	Hash     string
	PrevHash string
	Nonce    int
}

//生成区块
func (b Block) GenerateBlock(data string) (Block, error) {
	var newBlock Block
	newBlock.Data = data
	newBlock.Hash = newBlock.CalculateHash()
	return newBlock, nil
}

//生成散列值
//符合区块难度的hash
func (b Block) CalculateHash() string {

	record := b.Data + b.PrevHash
	hash := sha256.New()
	hash.Write([]byte(record))
	h := hash.Sum(nil)
	return hex.EncodeToString(h)
}

//创世块
func BigBang() Block {
	b := Block{}
	b.Data = "我是祖先"
	b.PrevHash = ""
	b.Hash = b.CalculateHash()
	return b

}
