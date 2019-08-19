package blc

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
	. "zeroChain/utils"
)

var (
	//nonce最大值
	maxNonce = math.MaxInt64
)

//挖矿难度值
const targetBits = 16

//工作量证明结构体
type ProofOfWork struct {
	//当前需要验证的区块
	block *Block

	//大数存储
	target *big.Int
}

//创建工作量证明
func NewProofOfWork(block *Block) *ProofOfWork {
	target := big.NewInt(1)
	//左移 223位 例如:0000000000000000000000010000000.......  target = 2 的223次方
	target.Lsh(target, uint(256-targetBits))
	pow := &ProofOfWork{block, target}
	return pow
}

//工作量证明
//哈希值只要小于 某个固定难度值那么就是正确的
func (pow *ProofOfWork) Run() (int, []byte) {

	var hashInt big.Int
	var hash [32]byte
	//设置难度初始值
	nonce := 0
	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
	for nonce < maxNonce {
		data := pow.prepareData(nonce)

		//计算当前字节数组 哈希值
		hash = sha256.Sum256(data)

		// \r 表示控制台只输出一条信息
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])
		//如果当前大数值 小于 pow中存储的大数值 则表示哈希是正确的
		//否则 nonce ++
		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}

	}
	return nonce, hash[:]
}

// 数据拼接
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join([][]byte{pow.block.PrevBlockHash,
		pow.block.Data,
		IntToHex(pow.block.TimeStamp),
		IntToHex(int64(targetBits)),
		IntToHex(int64(nonce))}, []byte{})

	return data
}

func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)

	hashInt.SetBytes(hash[:])
	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}
