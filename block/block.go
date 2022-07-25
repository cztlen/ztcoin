package block

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"time"
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
	//hash 满足区块链设置的条件
	return newBlock, nil
}

//生成散列值
func (b Block) CalculateHash() string {

	nonce := getRand()
	record := b.Data + b.PrevHash + strconv.Itoa(nonce)
	// fmt.Println(record)
	hash := sha256.New()
	hash.Write([]byte(record))
	h := hash.Sum(nil)
	return hex.EncodeToString(h)
}

//创世块
func BigBang(difficulty int) Block {
	b := Block{}
	b.Data = "创世块"
	b.PrevHash = ""
	b.Nonce = getRand()
	b.Hash = b.Mine(difficulty)
	return b

}

//获取的hash值规则
func GetAnswer(difficulty int) string {
	answer := ""
	for i := 0; i < difficulty; i++ {
		answer += "0"
	}
	return answer

}

////计算符合区块难度的hash mine
// proof of work

func (b Block) Mine(difficulty int) string {
	hash := b.CalculateHash()
	for {
		if hash[0:difficulty] == GetAnswer(difficulty) {
			break
		} else {
			hash = b.CalculateHash()
			// fmt.Println(hash[0:difficulty])
		}
	}
	fmt.Println("挖矿成功")
	return hash
}

//获取随机数

func getRand() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(100000000)
}
