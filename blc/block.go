package blc

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
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

	//Nonce随机数 挖矿难度值
	Nonce int
}

//创建新的区块
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), prevBlockHash, []byte(data), []byte{}, 0}
	pow := NewProofOfWork(block)

	//通过工作量证明 设置区块的哈希 以及nonce值
	nonce, hash := pow.Run()

	block.Hash = hash
	block.Nonce = nonce

	//校验区块合法性
	//valid:=pow.Validate()

	//block.SetHash()
	return block
}

// 设置当前区块哈希
func (b *Block) SetHash() {
	//fmt.Println(strconv.FormatInt(b.TimeStamp, 2))
	timestamp := []byte(strconv.FormatInt(b.TimeStamp, 2))

	hashBytes := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})

	//计算哈希值
	hash := sha256.Sum256(hashBytes)
	b.Hash = hash[:]
}

// 创建Zero区块
func NewZeroBlock() *Block {
	return NewBlock("Zero Block", []byte{'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0'})
}

// 序列化
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}
	return result.Bytes()
}

//反序列化
func Deserialize(b []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(b))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}
	return &block
}
