package blockchain

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"math/big"
)

const Difficulty = 10

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

func NewProofOfWork(block *Block) *ProofOfWork {
	// Initialize as 1, then left shift to set the difficulty
	targetVal := big.NewInt(1)
	targetVal.Lsh(targetVal, uint(256-Difficulty))

	return &ProofOfWork{block, targetVal}
}

/*
Takes a nonce and returns the computed data for hashing the block
*/
func (pow *ProofOfWork) ComputeData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			[]byte(pow.Block.PrevHash),
			[]byte(pow.Block.Data),
			make([]byte, 8), // Empty array for nonce
			make([]byte, 8), // Empty array for Difficulty
		},
		[]byte{},
	)

	binary.BigEndian.PutUint64(data[len(data)-16:], uint64(nonce))
	binary.BigEndian.PutUint64(data[len(data)-8:], uint64(Difficulty))
	return data
}

/*
Mine to find a valid block hash that meets the target
*/
func (pow *ProofOfWork) MineBlock() (int, []byte) {
	var intHash big.Int
	var computedHash [16]byte // Result of MD5 hash

	nonce := 0 // Start from 0, and continuously increase in the for loop

	for {
		computedData := pow.ComputeData(nonce)
		computedHash = md5.Sum(computedData)

		fmt.Printf("\r%x", computedHash)

		intHash.SetBytes(computedHash[:])

		if intHash.Cmp(pow.Target) == -1 {
			break
		}

		nonce++
	}
	fmt.Println()
	return nonce, computedHash[:]
}

/*
Validate the derived hash generated by the PoW work
*/
func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int

	computedData := pow.ComputeData(pow.Block.Nonce)

	computedHash := md5.Sum(computedData)
	intHash.SetBytes(computedHash[:])

	if intHash.Cmp(pow.Target) == -1 {
		return true
	} else {
		return false
	}
}
