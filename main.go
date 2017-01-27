package main

import "fmt"
import "os"
import "errors"
import "sort"

type dirtree struct {
	path     string      // string path
	info     os.FileInfo //
	hash     string      // hash of file if file under strict mode
	contents []dirtree   //Sorted contents of directory, empty if file
}

func (dt *dirtree) pString(depth int) string {
	//output := fmt.Sprintf(fmt.Sprintf("|%%s|%%-%d%%si\n", depth), dt.info.Mode(), dt.info.Name())
	output := fmt.Sprintf("|%s|%s\n", dt.info.Mode(), dt.info.Name())
	for i := 0; i < len(dt.contents); i++ {
		output += dt.contents[i].pString(depth + 1)
	}
	//"|mode|-*depth name"
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
		sort.Strings(rawcontents)
		for i := 0; i < len(rawcontents); i++ {
			filepath := rootdirname + string(os.PathSeparator) + rawcontents[i]
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
	tree2, err2 := getDirTree(dir2)
	if err2 != nil {
		panic("dir2 panic")
	}
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println(tree1.String())
	fmt.Println(tree2.String())
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	difftree, err3 := treediff(tree1, tree2)
	if err3 != nil {
		panic("err3 panic")
	}
	fmt.Println("differences: ")
	fmt.Println(difftree)
}
