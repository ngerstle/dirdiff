package main

import "fmt"
import "os"
import "errors"
import "sort"
import "strings"

type dirtree struct {
	path     string      // string path
	info     os.FileInfo //
	hash     string      // hash of file if file under strict mode
	contents []dirtree   //Sorted contents of directory, empty if file
}

func (dt *dirtree) dString(depth, dsize int) string {
	//TODO convert depth/dsize to int[] for more control over |, add bool for last in dir..
	drune := map[bool]string{false: "▸", true: "┬"}[(dt.info.IsDir())]
	output := fmt.Sprintf("│%s┃%s%s\n", dt.info.Mode(), strings.Repeat(" ", depth)+"├"+strings.Repeat("─", dsize)+drune, dt.info.Name())
	for i := 0; i < len(dt.contents); i++ {
		output += dt.contents[i].dString(depth+1+dsize, len(dt.info.Name()))
	}
	return output
}
func (dt *dirtree) cString(depth []int, last bool) string {
	//TODO convert depth/dsize to int[] for more control over |, add bool for last in dir..
	drune := map[bool]string{false: "▸", true: "┬"}[(dt.info.IsDir())]
	prune := map[bool]string{false: "├", true: "└"}[last]
	spacing := ""
	for i := 0; i < len(depth)-1; i++ {
		spacing += "│" + strings.Repeat(" ", depth[i])
	}
	output := fmt.Sprintf("│%s┃%s%s%s\n", dt.info.Mode(), spacing, prune+strings.Repeat("─", depth[len(depth)-1])+drune, dt.info.Name())
	ld := len(depth)
	if last && (ld > 1) {
		ndepth := make([]int, (ld - 1))
		copy(ndepth, depth)
		ndepth[ld-2] += depth[ld-1]
		depth = ndepth
	}
	for i := 0; i < len(dt.contents); i++ {
		//output += dt.contents[i].dString(depth+1+dsize, len(dt.info.Name()))
		ddepths := make([]int, (ld + 1))
		copy(ddepths, depth)
		ddepths[len(ddepths)-1] = len(dt.info.Name())
		output += dt.contents[i].cString(ddepths, (i == len(dt.contents)-1))
	}
	return output
}
func (dt *dirtree) String() string {
	//output := dt.dString(0, 0)
	output := dt.cString([]int{0}, false)
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
	//add ignore hidden (.filename) files (os compat?)
	//add ignore/exclude options

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
