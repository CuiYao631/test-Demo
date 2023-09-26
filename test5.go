package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {

	root := "data"
	//err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
	//	files = append(files, path)
	//	return nil
	//})
	//if err != nil {
	//	panic(err)
	//}
	//for _, file := range files {
	//	fmt.Println(file)
	//}
	GetAllFile(root)
}
func GetAllFile(pathname string) error {
	var files []string

	rd, err := ioutil.ReadDir(pathname)
	for _, fi := range rd {
		//最后一次出现的索引
		f := strings.Split(fi.Name(), ".")
		if f[1] == "json" {
			files = append(files, f[0])
		}
	}

	for _, v := range files {
		fmt.Println(v)
	}
	return err
}
