package scan

import (
	"crypto/sha256"
	"io"
	"os"

	"github.com/spaolacci/murmur3"
)

func HashFileContentSha256(filepath string) ([]byte, uint32, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, 0, err
	}
	defer f.Close()
	h := sha256.New()
	written, err := io.Copy(h, f)
	_ = written
	if err != nil {
		return nil, 0, err
	}
	contentHash := h.Sum(nil)
	lower32 := Lower32AsUint32(contentHash)
	return contentHash, lower32, nil
}

func Lower32AsUint32(sum []byte) uint32 {
	return uint32(sum[28])<<24 | uint32(sum[29])<<16 | uint32(sum[30])<<8 | uint32(sum[31])
}

// not really useful for anything right now
// func HashStringSha256Lower32(str string) uint32 {
// 	h := sha256.Sum256([]byte(str))
// 	return Lower32AsUint32(h[:]) // converts h from [32]byte into []byte without copying
// }

func HashStringMurmer32(str string) uint32 {
	return murmur3.Sum32([]byte(str))
}
func HashStringMurmer64(str string) uint64 {
	return murmur3.Sum64([]byte(str))
}
func HashStringMurmer63(str string) uint64 {
	hash := murmur3.Sum64([]byte(str))
	return hash & 0x7fffffffffffffff
}
