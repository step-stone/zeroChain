package blc

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

//基础区块结构
//16进制 hash  64个数字  32字节  256位
type Block struct {
	//时间戳，创建区块时间
	TimeStamp int64

	//上一区块hash
	PrevBlockHash []byte

	//交易数据
	Data []byte

	//当前区块hash
	Hash []byte
}

//创建新的区块
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(),
		prevBlockHash,
		[]byte(data),
		[]byte{}}

	block.SetHash()
	return block
}

// 设置当前区块哈希
func (b *Block) SetHash() {
	fmt.Println(strconv.FormatInt(b.TimeStamp, 2))
	timestamp := []byte(strconv.FormatInt(b.TimeStamp, 2))

	hashBytes := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})

	//计算哈希值
	hash := sha256.Sum256(hashBytes)
	b.Hash = hash[:]
}