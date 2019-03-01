package p1

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type BlockChain struct {
	Chain  map[int32][]Block
	Length int32
}

// Initialise a new BlockChain
func (b *BlockChain) Initial() {
	//Set values from param
	b.Length = 0
	b.Chain = make(map[int32][]Block)

}

//TODO fix this
func (b *BlockChain) Get(height int32) ([]Block, error) {
	b.Chain[height] = append(b.Chain[height], *new(Block))
	result := b.Chain[height]
	if len(result) == 0 {
		return nil, errors.New("No Blocks")
	}
	return result, errors.New("")
}

// Insert a block into the BlockChain
func (b *BlockChain) Insert(block Block) {
	//Check if block is stored in array if to insert it.
	if !(hashInArray(block.Header.Hash, b.Chain[block.Header.Height])) {
		b.Chain[block.Header.Height] = append(b.Chain[block.Header.Height], block)
	}
}

//Check if the hash value of the block is already stored in the array.
func hashInArray(blockHash string, list []Block) bool {
	for _, b := range list {
		if b.Header.Hash == blockHash {
			return true
		}
	}
	return false
}

// Function that encodes a BlockChain into a json array string
func (b *BlockChain) EncodeToJSON() (string, error) {
	encodedBlockChain := "["

	// Iterate each index in the hashmap.
	for _, v := range b.Chain {
		// For each index, iterate array of blocks.
		for _, element := range v {
			encodedBlockChain += element.EncodeToJSON() + ","
		}
	}
	encodedBlockChain = strings.TrimRight(encodedBlockChain, ",")
	encodedBlockChain += "]"

	return encodedBlockChain, nil

}

//Function that takes a json array string of blocks and creates a BlockChain containing every block
func DecodeJsonToBlockChain(data string) (BlockChain, error) {
	//Create new a Blockchain
	blockChain := new(BlockChain)
	blockChain.Initial()

	// Store each block object as a json struct in blocks
	var blocks []Encoded_block
	err := json.Unmarshal([]byte(data), &blocks)
	if err != nil {
		fmt.Println("error:", err)
	}

	//Iterate each block stored in the chain
	for _, block := range blocks {
		val, err := json.Marshal(block)
		if err != nil {
			fmt.Println("error:", err)
		}
		blockChain.Insert(DecodeFromJson(string(val)))
	}
	return *blockChain, err
}
