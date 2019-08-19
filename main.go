package main

import (
	"fmt"
	"time"
	"zeroChain/blc"
)

func main() {

	block := blc.NewZeroBlock()
	fmt.Println("当前时间戳:", block.TimeStamp)
	fmt.Printf("前一个区块哈希: %x \n", block.PrevBlockHash)
	fmt.Println("当前交易数据:", string(block.Data))
	fmt.Printf("当前区块哈希: %x\n", block.Hash)
	fmt.Println("---------------------------------------------")

	blockChain := blc.NewBlockChain()
	blockChain.AddBlock("send 20 btc")

	//遍历区块链
	for _, block := range blockChain.Blocks {
		fmt.Println("当前时间戳:", time.Unix(block.TimeStamp, 0).Format("2006-01-02 03:04:05 PM"))
		fmt.Printf("前一个区块哈希: %x \n", block.PrevBlockHash)
		fmt.Println("当前交易数据:", string(block.Data))
		fmt.Printf("当前区块哈希: %x\n", block.Hash)
		fmt.Printf("当前区块Nonce: %d\n", block.Nonce)
	}

}
