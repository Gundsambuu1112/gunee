package blockchain

import (
	"bytes"
	"crypto/sha256"
)

// Блокчэйны бүтэц
type BlockChain struct {
	blocks []*Block
}

// Блокын бүтэц.
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

type Commandline struct{}

//өгөгдөл болон өмнөх блокийн хэш дээр үндэслэн блокийн хэшийг тооцоолно.
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

//өгөгдсөн өгөгдөл болон өмнөх хэштэй шинэ блок үүсгэнэ.
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

// Шинэ блок нэмнэ.
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

//Genesis функц нь тодорхой өгөгдөл бүхий эхний блокийг үүсгэх замаар блокчэйныг эхлүүлдэг.
func Genesis() *Block {
	return CreateBlock("Gundsambuu", []byte{})
}

// эхний блокийг үүсгэж, блокчейн бүтцийг тухайн блокоор эхлүүлснээр шинэ блок гинжийг бий болгоно.
func InitBlockChain() *BlockChain {
	// BlockChain бүтцийн шинэ жишээ үүсгэх
	blockchain := BlockChain{
		blocks: []*Block{Genesis()},
	}

	// Заагчийг BlockChain бүтэц рүү буцаана
	return &blockchain
}
