package blc

//区块链结构
type BlockChain struct {
	//存储有序区块信息
	Blocks []*Block
}

//添加新区块
func (chain *BlockChain) AddBlock(data string) {

	//根据索引获取上一个区块
	prevBlock := chain.Blocks[(len(chain.Blocks) - 1)]
	//创建新区块
	newBlock := NewBlock(data, prevBlock.Hash)

	chain.Blocks = append(chain.Blocks, newBlock)
}

// 创建区块链
func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{NewZeroBlock()}}
}
