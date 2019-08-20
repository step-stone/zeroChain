package blc

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

const dbFile = "block.db"
const blockBuckets = "block"

//区块链结构
type BlockChain struct {
	//存储有序区块信息
	//Blocks []*Block

	//区块链最后一个区块的哈希值
	Tip []byte

	//数据库
	DB *bolt.DB
}

//添加新区块
func (chain *BlockChain) AddBlock(data string) {

	newBlock := NewBlock(data, chain.Tip)

	err := chain.DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBuckets))
		err := bucket.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			log.Panic(err)
		}

		err = bucket.Put([]byte("l"), newBlock.Hash)

		if err != nil {
			log.Panic(err)
		}

		chain.Tip = newBlock.Hash
		return nil
	})

	if err != nil {
		log.Panic(err)
	}
}

// 创建区块链
func NewBlockChain() *BlockChain {

	//最后一个区块的哈希值
	var tip []byte

	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {

		bucket := tx.Bucket([]byte(blockBuckets))
		if bucket == nil {
			fmt.Println("No exist table [block]. creat a new block table")

			//创建第一个区块
			zeroBlock := NewZeroBlock()
			bucket, err = tx.CreateBucket([]byte(blockBuckets))
			if err != nil {
				log.Panic(err)
			}
			//存储区块信息
			err = bucket.Put(zeroBlock.Hash, zeroBlock.Serialize())
			if err != nil {
				log.Panic(err)
			}

			err = bucket.Put([]byte("l"), zeroBlock.Hash)
			if err != nil {
				log.Panic(err)
			}
			tip = zeroBlock.Hash

		} else {
			tip = bucket.Get([]byte("l"))
		}
		return nil
	})

	return &BlockChain{tip, db}
}
