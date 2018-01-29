package main

import (
	"io/ioutil"
	"fmt"
	"path/filepath"
	"os"
	"github.com/axgle/mahonia"
)

func main() {
	//按等级重命名
	var dir = "F:/tmp/170523temp_CNC_DTPinstructions";
	files, _ := ioutil.ReadDir(dir)

	for _, file := range files {
		if file.IsDir() {
			continue
		} else {
			fmt.Println(file.Name())
		}
	}

	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		fmt.Println(path)
		//println(f.Name())
		mahonia.NewDecoder("")
		fmt.Println(filepath.Join(filepath.Dir(path),f.Name()))
		//os.Rename(path, strings.Replace(path, dir, dir + ".out", -1));
		//os.Rename(path, filepath.Join());
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}