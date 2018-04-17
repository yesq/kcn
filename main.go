package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

var path string
var script string
var aFile *CheckFile
var ok bool

const dataJSON string = "./data.json"

func init() {
	log.Printf("%v", os.Args)
	if len(os.Args) != 1 {
		fmt.Printf(`
kcn:
  a toole to check service
  need two ENV :
    KCN_CHECKFILE : a file. if file changed, return false.
    KCN_SCRIPT    : a script path. run script.
	
`)
	}
	if err := initCheckFile(); err != nil {
		log.Fatal(err)
	}
	if err := initScript(); err != nil {
		log.Fatal(err)
	}
}

func initScript() error {
	script = os.Getenv("KCN_SCRIPT")
	if script == "" {
		return nil
	}
	_, err := os.Stat(script)
	if os.IsNotExist(err) {
		return fmt.Errorf("File not exist %v", err)
	}
	return err
}

func initCheckFile() error {
	path = os.Getenv("KCN_CHECKFILE")
	_, err := os.Stat(dataJSON)
	if os.IsNotExist(err) {
		log.Println("First run. Save file info.")
		// 第一次运行，没有保存文件信息
		// 读取文件信息
		newFile, err := NewCheckFile(path, true)
		if newFile == nil {
			// 没有指定文件
			log.Println("No file")
			return nil
		}
		if err != nil {
			return fmt.Errorf("Can not new a CheckFile: %v", err)
		}
		// 保存文件信息
		err = Save(dataJSON, newFile)
		if err != nil {
			dir, erro := filepath.Abs(filepath.Dir(os.Args[0]))
			return fmt.Errorf("Can not save info to disk %v : %v, %v", dir, err, erro)
		}
		return nil
	}
	log.Println("Not first run. Load info.")
	aFile, err = NewCheckFile(path, false)
	err = Load(dataJSON, aFile)
	if err != nil {
		return fmt.Errorf("Can not load info : %v", err)
	}
	return err
}

func main() {
	if aFile != nil {
		b, e := aFile.Check()
		if e != nil {
			log.Fatal(e)
		}
		if b == false {
			log.Fatal("file changed")
		}
		log.Printf("%v file same\n", path)
		ok = true
	}

	if script != "" {
		println(script)
		out, err := exec.Command("sh", script).Output()
		fmt.Print(string(out))
		if err != nil {
			log.Fatal(err)
		} else {
			ok = true
		}
	}
	if ok == false {
		log.Fatal("No one check exec. Please set one at least")
	}
}
