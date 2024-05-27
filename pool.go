package hashstructure

import (
	"crypto/sha256"
	"hash"
	"sync"
)

var shaPool = sync.Pool{New: func() interface{} {
	return sha256.New()
}}

func getSHA256() hash.Hash {
	return shaPool.Get().(hash.Hash)
}

func putSHA256(h hash.Hash) {
	h.Reset()
	shaPool.Put(h)
}
