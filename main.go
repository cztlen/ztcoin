package main

import (
	"fmt"
	"math/rand"
	"time"
	"ztcoin/block"
	"ztcoin/chain"
)

func main() {
	rand.Seed(time.Now().Unix())
	r := rand.Intn(100)

	// p := pow.ProofOfWork("ztcoin")
	fmt.Println(r)
}

func Test() {
	b := block.BigBang()
	chain := chain.New(b) //初始化chain

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
	for _, b := range chain {
		fmt.Printf("%+v\n", b)
	}
	chain[1].Data = "转账 1元"
	chain[1].Hash = chain[1].CalculateHash()
	for _, b := range chain {
		fmt.Printf("%+v\n", b)
	}
}
