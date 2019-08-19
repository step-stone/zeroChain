package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

//定义数据库文件
const dbFile = "blockChain.db"

//定义数据仓库
const dbBucket = "block"

func main() {

	db, err := bolt.Open(dbFile, 0600, nil)

	//插入或更新数据库
	db.Update(func(tx *bolt.Tx) error {
		//获取数据表
		b := tx.Bucket([]byte(dbBucket))

		if b == nil {
			fmt.Println("No existing blockChain.db Creating a new dbFile")
			//创建表
			b, err := tx.CreateBucket([]byte(dbBucket))
			if err != nil {
				log.Panic(err)
			}

			err = b.Put([]byte("zero"), []byte("123456"))
			if err != nil {
				log.Panic(err)
			}
		}
		//value := b.Get([]byte("zero"))
		//fmt.Println("查询数据库值:", value)
		return nil
	})

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(dbBucket))
		value := b.Get([]byte("zero"))
		fmt.Println("查询数据:", string(value))
		return nil
	})
	defer db.Close()
	if err != nil {
		log.Panic(err)
	}
}
