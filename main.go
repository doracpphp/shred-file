package main

import (
	"fmt"
	"os"
)

func zeroFillFile(filepath string) error {
	// ファイルをオープンしてサイズを取得
	file, err := os.OpenFile(filepath, os.O_WRONLY, 0)
	if err != nil {
		return err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}
	fileSize := fileInfo.Size()

	// ファイルをゼロで上書き（ゼロフィリング）
	zeroes := make([]byte, fileSize)
	err = file.Truncate(0)
	if err != nil {
		return err
	}
	_, err = file.Write(zeroes)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	filePath := "path/to/file.txt" // ゼロフィリングしたいファイルのパス

	err := zeroFillFile(filePath)
	if err != nil {
		fmt.Println("エラー:", err)
		return
	}
}
