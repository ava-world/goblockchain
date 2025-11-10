package main

import (
	"fmt"
	"log"
	"time"
)

type Block struct {
	nonce        int
	previousHash string
	timestamp    int64
	transactions []string
}

func newBlock(nonce int, previousHash string) *Block {

	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	return b
}

func (b *Block) Print() {

	fmt.Printf("timestamp %d\n", b.timestamp)
	fmt.Printf("nonce %d\n", b.nonce)
	fmt.Printf("previous_hash %s\n", b.previousHash)
	fmt.Printf("transactions %s\n", b.transactions)
}

type Blockchain struct {
	transactionnPool []stringchain []*Block
}

unc NewBlockchain() *Blockchain {
	bc := new(Blockchin)
	bc.Create(0, "Init hash")
	return bc
}

func (bc *Blockchain) CreateBlock(nonce int, previousHash string) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b

}
// ----------
func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	b := newBlock(0, "init hash")
	b.Print()

}
