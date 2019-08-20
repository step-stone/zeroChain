package blc

import (
	"github.com/boltdb/bolt"
	"log"
)

// 迭代器
type BlockchainIterator struct {
	//当前正在遍历的哈希
	CurrentHash []byte

	DB *bolt.DB
}

//迭代器
func (chain *BlockChain) Iterator() *BlockchainIterator {

	return &BlockchainIterator{chain.Tip, chain.DB}
}

func (iterator *BlockchainIterator) Next() *BlockchainIterator {

	var nextHash []byte

	err := iterator.DB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBuckets))

		currentBlockBytes := bucket.Get(iterator.CurrentHash)

		currentBlock := Deserialize(currentBlockBytes)
		nextHash = currentBlock.PrevBlockHash
		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	return &BlockchainIterator{nextHash, iterator.DB}

}
