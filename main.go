package main

import "fmt"
import "os"
import "errors"

type dirtree struct {
	//fd       File
	path     string
	hash     string
	contents []dirtree
}

func getDirTree(rootdirname string) (dirtree, error) {
	fmt.Println("Getting Dir tree: " + rootdirname)
	//os.Readdir{names}(*File)...
	//TODO
	return dirtree{}, nil
}

func treediff(tree1, tree2 dirtree) (string, error) {
	if len(tree1.contents) != len(tree2.contents) { //TODO
		return "", errors.New("testing?")
	}
	return "OK", nil
}

func main() {
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
	difftree, err3 := treediff(tree1, tree2)
	if err3 != nil {
		panic("err3 panic")
	}
	fmt.Println("differences: ")
	fmt.Println(difftree)
}
