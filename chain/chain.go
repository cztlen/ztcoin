package chain

import (
	"errors"
	"ztcoin/block"
)

type Blockchain []block.Block

func New(b block.Block) Blockchain {
	chain := Blockchain{}
	chain = append(chain, b)
	// cache.GoCache.Set("cztCoin", chain, 2*time.Minute)
	return chain
}

//从链上获取前一个区块的hash
func (bc Blockchain) GetLastBlockHash() string {
	lastBlock := bc[len(bc)-1]
	return lastBlock.Hash
}

//给链上添加一个区块
func (bc Blockchain) AddBlockToChain(block block.Block) Blockchain {
	prevHash := bc.GetLastBlockHash()
	block.PrevHash = prevHash
	bc = append(bc, block)
	// cache.GoCache.Set("cztCoin", bc, 2*time.Minute)
	return bc
}

//验证区块是否合法

func (bc Blockchain) VerifyBlock() error {

	if len(bc) == 1 {
		if bc[0].Hash == bc[0].CalculateHash() {
			return nil
		}
		return errors.New("创世块被篡改")
	}
	for i := 1; i < len(bc)-1; i++ {
		if bc[i].Hash != bc[i].CalculateHash() {
			return errors.New("数据被篡改")
		}
		if bc[i].PrevHash != bc[i-1].CalculateHash() {
			return errors.New("前后区块链接断裂")
		}
	}
	return nil
}
