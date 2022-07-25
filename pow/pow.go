package pow

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
)

//工作量证明
func ProofOfWork(data string) string {
	hash := sha256.New()
	x := 1
	hash.Write([]byte(data + strconv.Itoa(x)))
	h := hash.Sum(nil)
	hashStr := hex.EncodeToString(h)
	for {
		if hashStr[0:1] == "0" {
			break
		} else {
			x++
		}
		hash.Write([]byte(data + strconv.Itoa(x)))
		h := hash.Sum(nil)
		hashStr = hex.EncodeToString(h)
	}
	fmt.Println(x)
	return hashStr

}
