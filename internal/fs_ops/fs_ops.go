package fs_ops

import (
	"fmt"
	"os"
)

func ReadWholeFile(address string) []byte {
	contents, err := os.ReadFile(address)

	if err != nil {
		fmt.Printf("Error %v", err)
		panic(err)
	}

	return contents
}
