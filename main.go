package main

import (
	"fmt"
	"ztcoin/block"
	"ztcoin/chain"
)

func main() {

	// b := 1

	// p := pow.ProofOfWork("ztcoin")
	fmt.Println(block.BigBang(3))
}

func Test() {
	b := block.BigBang(3)
	chain := chain.New(b, 1) //初始化chain

	nb, _ := b.GenerateBlock("转账 10元") //生成新区块
	chain = chain.AddBlockToChain(nb)  //向链上添加区块

	nb1, _ := b.GenerateBlock("转账 100元")
	chain = chain.AddBlockToChain(nb1)

	nb2, _ := b.GenerateBlock("转账 1000元")
	chain = chain.AddBlockToChain(nb2)

	//验证区块
	if err := chain.VerifyBlock(); err != nil {
		fmt.Println(err.Error())
	}
	for _, b := range chain.Chain {
		fmt.Printf("%+v\n", b)
	}
	chain.Chain[1].Data = "转账 1元"
	chain.Chain[1].Hash = chain.Chain[1].CalculateHash()
	for _, b := range chain.Chain {
		fmt.Printf("%+v\n", b)
	}
}
