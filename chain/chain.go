package chain

import (
	"errors"
	"ztcoin/block"
)

type Blockchain struct {
	Chain      []block.Block
	Difficulty int //区块难度
}

func New(b block.Block, level int) Blockchain {
	bc := Blockchain{}
	bc.Chain = append(bc.Chain, b)
	// cache.GoCache.Set("cztCoin", chain, 2*time.Minute)
	bc.Difficulty = level
	return bc
}

//从链上获取前一个区块的hash
func (bc Blockchain) GetLastBlockHash() string {
	lastBlock := bc.Chain[len(bc.Chain)-1]
	return lastBlock.Hash
}

//给链上添加一个区块
func (bc Blockchain) AddBlockToChain(block block.Block) Blockchain {
	prevHash := bc.GetLastBlockHash()
	block.PrevHash = prevHash
	block.Hash = block.Mine(bc.Difficulty)
	bc.Chain = append(bc.Chain, block)
	// cache.GoCache.Set("cztCoin", bc, 2*time.Minute)
	return bc
}

//验证区块是否合法

func (bc Blockchain) VerifyBlock() error {

	if len(bc.Chain) == 1 {
		if bc.Chain[0].Hash == bc.Chain[0].CalculateHash() {
			return nil
		}
		return errors.New("创世块被篡改")
	}
	for i := 1; i < len(bc.Chain)-1; i++ {
		if bc.Chain[i].Hash != bc.Chain[i].CalculateHash() {
			return errors.New("数据被篡改")
		}
		if bc.Chain[i].PrevHash != bc.Chain[i-1].CalculateHash() {
			return errors.New("前后区块链接断裂")
		}
	}
	return nil
}
