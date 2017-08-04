package block

import (
	"crypto/sha1"
	"hash"
	"strconv"
	"time"
)

type Block struct {
	Index        uint64
	Timestamp    time.Time
	Data         string
	PreviousHash hash.Hash
	Hash         hash.Hash
}

func New(i uint64, t time.Time, d string, p hash.Hash) Block {
	b := Block{i, t, d, p, nil}
	b.Hash = b.HashIt()
	return b
}

func (b *Block) HashIt() hash.Hash {
	i := strconv.FormatUint(b.Index, 16)
	t := b.Timestamp.String()
	ph := b.PreviousHash.Sum(nil)
	h := sha1.New()
	h.Write([]byte(i + t + string(ph)))
	return h
}

func Genesis() Block {
	return New(0, time.Now(), "Genesis Block", sha1.New())
}

func Next(pb Block) Block {
	i := pb.Index + 1
	t := time.Now()
	d := "Block #" + string(i)
	return New(i, t, d, pb.Hash)
}
