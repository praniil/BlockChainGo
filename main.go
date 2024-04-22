package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	data      map[string]interface{}
	hash      string
	prevHash  string
	timestamp time.Time
	nonce     int
}

//Blockchain containing Blocks
type BlockChain struct {
	genesisBlock Block
	chain        []Block
	target       int
}

//to include a block into a blockchain a miner mines a new block calculating a hash puzzle
//calculation of the hash of the block
//it includes the hashing of transactions(data), nonce, prevHash to create the new hash of the block

func (b Block) calculateHash() string {
	data, err := json.Marshal(b.data)
	if err != nil {
		fmt.Printf("error in encoding the data: %v", err)
	}
	blockData := b.prevHash + string(data) + b.timestamp.String() + strconv.Itoa(b.nonce)
	newHash := sha256.Sum256([]byte(blockData))
	return fmt.Sprintf("%x", newHash)
}

func (b *Block) mine (target int) {
	for !strings.HasPrefix(b.hash, strings.Repeat("0", target)) {
		b.nonce++
		b.hash = b.calculateHash()
	} 
}

func CreatBlockchain(target int) BlockChain{
	genesisblock := Block{
		hash: "0",
		timestamp: time.Now(),
	}
	return BlockChain{
		genesisblock, 
		[]Block{genesisblock}, 
		target,
	}
}

func (b *BlockChain) addBlock (from, to string, amount float64) {
	blockData := map[string]interface{}{
		"from" : from, 
		"to" : to,
		"amount" : amount,
	}
	prevBlock := b.chain[len(b.chain) - 1]
	newBlock := Block{
		data : blockData, 
		prevHash: prevBlock.hash,
		timestamp: time.Now(),
	}
	newBlock.mine(b.target)
	b.chain = append(b.chain, newBlock)
}

func (b *BlockChain) isValid() bool{
	for i:= range b.chain[1:] {
		previousBlock := b.chain[i]
		currentBlock := b.chain[i + 1]
		if currentBlock.hash != currentBlock.calculateHash() || currentBlock.prevHash != previousBlock.hash{
			return false
		}
	}
	return true
}

func main() {
	blockchain := CreatBlockchain(2)
	blockchain.addBlock("Alice", "Bob", 5)
	blockchain.addBlock("John", "Bob", 2)
	fmt.Println(blockchain.chain)
	fmt.Println(blockchain.isValid())

}
