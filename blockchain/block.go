package blockchain

import (
	"bytes"
	"crypto/md5"
	"math/rand"
	"time"
)

type Block struct {
	Hash     string
	Data     string
	PrevHash string
	Nonce    int
}

func (b *Block) ComputeHash() {
	concatenatedData := bytes.Join([][]byte{[]byte(b.Data), []byte(b.PrevHash)}, []byte{})
	computedHash := md5.Sum(concatenatedData)
	b.Hash = string(computedHash[:])
}

// func (b *Block) ComputeHash() {
// 	concatenatedData := bytes.Join([][]byte{[]byte(b.Data), []byte(b.PrevHash)}, []byte{})
// 	computedHash := sha256.Sum256(concatenatedData)
// 	b.Hash = hex.EncodeToString(computedHash[:])
// }

func CreateBlock(data string, prevHash string) *Block {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	initialNonce := rand.Intn(10000)

	block := &Block{"", data, prevHash, initialNonce}

	newPow := NewProofOfWork(block)

	nonce, hash := newPow.MineBlock()

	block.Hash = string(hash[:])
	block.Nonce = nonce

	return block
}

func Genesis() *Block {
	return CreateBlock("Genesis", "")
}
