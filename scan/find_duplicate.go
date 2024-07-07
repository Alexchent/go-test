package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
)

type FileInfo struct {
	Path string
	Hash []byte
}

func getFileHash(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return nil, err
	}

	return hash.Sum(nil), nil
}

func scanDirectory(dir string, fileChan chan<- FileInfo, wg *sync.WaitGroup) {
	defer wg.Done()

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		//fmt.Println(path)
		if !info.IsDir() {
			hash, err := getFileHash(path)
			if err != nil {
				return err
			}
			fileChan <- FileInfo{Path: path, Hash: hash}
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error scanning directory:", err)
	}
}

func findDuplicates(fileChan <-chan FileInfo) {
	fileMap := make(map[string][]string)

	for fileInfo := range fileChan {
		hashStr := fmt.Sprintf("%x", fileInfo.Hash)
		fileMap[hashStr] = append(fileMap[hashStr], fileInfo.Path)
	}

	for hash, paths := range fileMap {
		if len(paths) > 1 {
			fmt.Printf("Duplicate files with hash %s:\n", hash)
			for _, path := range paths {
				fmt.Println(path)
			}
			fmt.Println()
		}
	}
	fmt.Println("finish")
}

func main() {
	var path string
	fmt.Printf("请输入要扫描的目录:\n")

	_, err := fmt.Scan(&path)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	dir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if path == "" || path == "/" {
		path = dir + "/Downloads"
	}
	fmt.Println("正在扫描：" + path)

	fileChan := make(chan FileInfo)
	var wg sync.WaitGroup

	wg.Add(1)
	go scanDirectory(path, fileChan, &wg)

	go func() {
		wg.Wait()
		close(fileChan)
	}()

	findDuplicates(fileChan)
}
