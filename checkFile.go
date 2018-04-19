package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

// CheckFile : a file to be checked whether changed
type CheckFile struct {
	Path    string
	content string
	Sha     []byte
}

// NewCheckFile : New a CheckFile struct
func NewCheckFile(filePath string, withSha bool) (*CheckFile, error) {
	if filePath == "" {
		return nil, fmt.Errorf("no file to check")
	}
	aFile := new(CheckFile)
	aFile.Path = filePath
	if withSha == true {
		return aFile, aFile.GenSha()
	}
	return aFile, nil
}

func sha(input []byte) []byte {
	first := sha256.New()
	first.Write([]byte(input))

	return first.Sum(nil)
}

// GenSha : generate sha of file
func (a *CheckFile) GenSha() error {
	return a.genSha()
}

func (a *CheckFile) genSha() error {
	b, err := getModTime([]byte(a.Path))
	if err != nil {
		return fmt.Errorf("Can not read file %s", a.Path)
	}
	a.Sha = sha(b)
	return nil
}

// Check : check current file with old one
func (a CheckFile) Check() (bool, error) {
	return a.check()
}

func (a CheckFile) check() (bool, error) {
	b, err := getModTime([]byte(a.Path))
	if err != nil {
		return false, err
	}
	return bytes.Equal(a.Sha, sha(b)), nil
}

func getModTime(filePath []byte) ([]byte, error) {
	filePathString := string(filePath)
	info, err := os.Stat(filePathString)
	if err != nil {
		return nil, err
	}
	allTimeBytes, err := info.ModTime().MarshalText()
	if err != nil {
		return nil, err
	}
	// file 是文件。返回修改时间
	if info.IsDir() == false {
		return allTimeBytes, nil
	}
	// file 是文件夹。返回子文件的修改时间。递归。
	files, err := ioutil.ReadDir(filePathString)
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		timeBytes, err := getModTime([]byte(path.Join(filePathString, f.Name())))
		if err != nil {
			return nil, err
		}
		allTimeBytes = append(allTimeBytes, timeBytes...)
	}
	return allTimeBytes, nil
}
