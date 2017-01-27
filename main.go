package main

import "fmt"
import "os"
import "errors"

type dirtree struct {
	//fd       File
	path     string      // string path
	info     os.FileInfo //
	hash     string      // hash of file if file under strict mode
	contents []dirtree   //Sorted contents of directory, empty if file
}

func (dt *dirtree) pString(depth int) string {
	output := ""
	//output = fmt.sprintf()
	return output
}
func (dt *dirtree) String() string {
	output := dt.pString(0)
	return output
}

func getDirTree(rootdirname string) (dirtree, error) {
	fmt.Println("Getting Dir tree: " + rootdirname)

	fileinfo, err := os.Stat(rootdirname) //Stat/LStat
	if err != nil {
		panic("os.Stat panic")
	}

	contents := make([]dirtree, 0)
	//contents := []dirtree{}
	hash := ""
	if fileinfo.Mode().IsDir() {
		file, err := os.Open(rootdirname)
		if err != nil {
			panic("os.open panic")
		}
		rawcontents, err := file.Readdirnames(-1)
		if err != nil {
			panic("readdir panic")
		}
		err = file.Close()
		if err != nil {
			panic("can't close file")
		}
		//TODO sort rawcontents by name
		for i := 0; i < len(rawcontents); i++ {
			filepath := rootdirname + rawcontents[i] //TODO fix
			newnode, err := getDirTree(filepath)
			if err == nil {
				contents = append(contents, newnode)
			} else {
				return dirtree{rootdirname, fileinfo, hash, contents}, err
			}

		}
	}
	return dirtree{rootdirname, fileinfo, hash, contents}, nil
}

func treediff(tree1, tree2 dirtree) (string, error) {
	if len(tree1.contents) != len(tree2.contents) { //TODO
		return "", errors.New("testing?")
	}
	return "OK", nil
}

func main() {

	//add strict flag (hash instead of size)
	//add links flag (follow/don't follow links)

	fmt.Printf("Hellow\n")
	fmt.Println("All args:")
	dir1 := os.Args[1]
	dir2 := os.Args[2]
	//fmt.Println(os.Args)
	fmt.Printf("Dir 1: ")
	fmt.Println(dir1)
	fmt.Printf("Dir 2: ")
	fmt.Println(dir2)
	tree1, err1 := getDirTree(dir1) //convert to goroutine..
	if err1 != nil {
		panic("dir1 panic")
	}
	tree2, err2 := getDirTree(dir1)
	if err2 != nil {
		panic("dir2 panic")
	}
	fmt.Println(tree1)
	fmt.Println(tree2)
	difftree, err3 := treediff(tree1, tree2)
	if err3 != nil {
		panic("err3 panic")
	}
	fmt.Println("differences: ")
	fmt.Println(difftree)
}
