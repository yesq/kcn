package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
)

// CheckFile : a file to be checked whether changed
type CheckFile struct {
	Path    string
	content string
	Sha     []byte
}

// NewCheckFile : New a CheckFile struct
func NewCheckFile(path string, withSha bool) (*CheckFile, error) {
	if path == "" {
		return nil, fmt.Errorf("no file to check")
	}
	aFile := new(CheckFile)
	aFile.Path = path
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
	b, err := ioutil.ReadFile(a.Path)
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
	b, err := ioutil.ReadFile(a.Path)
	if err != nil {
		return false, err
	}
	return bytes.Equal(a.Sha, sha(b)), nil
}
