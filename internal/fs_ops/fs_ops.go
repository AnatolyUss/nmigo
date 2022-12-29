package fs_ops

import (
	"fmt"
	"os"
	"path/filepath"
)

func ReadWholeFile(address string) []byte {
	contents, err := os.ReadFile(address)

	if err != nil {
		fmt.Printf("Error %v", err)
		panic(err)
	}

	return contents
}

func GetBaseDirectory() string {
	baseDir, ok := os.LookupEnv("aux_dir")

	if !ok {
		pwd, err := os.Getwd()

		if err != nil {
			panic(err)
		}

		baseDir = filepath.Join(pwd, "..", "..")
	}

	return baseDir
}

func CreateLogsDirectory(path string) {
	fmt.Printf("[CreateLogsDirectory] Creating directory %s...\n", path)
	_, ok := os.Stat(path)

	if ok != nil {
		err := os.Mkdir(path, 0777)

		if err != nil && os.IsExist(err) {
			fmt.Printf("[CreateLogsDirectory] Failed to create directory due to %s", err)
			panic(err)
		}
	} else {
		fmt.Printf("[CreateLogsDirectory] Directory %s already exists", path)
	}
}
