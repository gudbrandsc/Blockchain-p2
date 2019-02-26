package p1

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type BlockChain struct {
	Chain   map[int32][]Block
	Length int32
}
//Structure used to create a JSON object.
type EncodedBlockChain struct {
	Hash		[]Encoded_block

}

// Initialise a new BlockChain
func (b *BlockChain) Initial() {

	//Set values from param
	b.Length = 0;
	b.Chain = make(map[int32][]Block);

}


func (b *BlockChain) Get(height int32)  ([]Block, error) {
	b.Chain[height] = append(b.Chain[height], *new(Block))
	result := b.Chain[height];

	return result, errors.New("path_not_found");
}


func (b *BlockChain) Insert(block Block)  {

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

func (b *BlockChain) EncodeToJSON() (string, error) {
	d := "[";
	fmt.Println("Lets encode the data")

	for k, v := range b.Chain {
		for _, element := range v {
			d += element.EncodeToJSON() + ","
		}
		k = k;
	}
	d = strings.TrimRight(d, ",")
	d += "]"

	return d, nil

}



func DecodeJsonToBlockChain(data string)  (BlockChain, error) {
	blockChain := new(BlockChain);
	blockChain.Initial();
	var blocks []Encoded_block
	json.Unmarshal([]byte(data), &blocks)



	for _, block := range blocks {
		//fmt.Println(block)

		val, err := json.Marshal(block)
		if err != nil {
			fmt.Println("Unable")
			fmt.Println(err)
		}
		block := DecodeFromJson(string(val));
		fmt.Println("insert block ")

		blockChain.Insert(block);
	}

	return *blockChain, nil
}







