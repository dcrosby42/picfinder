package scan

import (
	"crypto/sha256"
	"io"
	"os"

	"github.com/spaolacci/murmur3"
)

func HashFileContentSha256(filepath string) ([]byte, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	h := sha256.New()
	written, err := io.Copy(h, f)
	_ = written
	if err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
	// fmt.Printf("Sum256() (%d written) => %x\n", w, h.Sum(nil))
}

func HashStringMurmer32(str string) uint32 {
	return murmur3.Sum32([]byte(str))
}
