package packs

import (
	"fmt"
	"os"
	"path/filepath"
)

func CalFileSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(s string, info os.FileInfo, err error) error {
		fmt.Println(s)
		fmt.Println(info, info.Name())
		fmt.Println(err)
		fmt.Println("-------------")
		if err != nil {
			return err
		}
		if !info.IsDir() {
			fmt.Println(info.Name())
			size += info.Size()
		}
		return err
	})
	return size, err
}
