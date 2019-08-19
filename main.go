package main

import (
	"fmt"
	"zeroChain/blc"
)

func main() {

	block := blc.NewBlock("First block", []byte{'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0'})
	fmt.Println("当前时间戳:", block.TimeStamp)
	fmt.Printf("前一个区块哈希: %x \n", block.PrevBlockHash)
	fmt.Println("当前交易数据:", block.Data)
	fmt.Printf("当前区块哈希: %x", block.Hash)

}
