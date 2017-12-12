/*
findFileDuplicates - given a directory, output files that have identical contents in [][]string
example: input "c:/temp"
output: [
			[file1, file2],
			[file3, file4, file5,...],
		]
*/
package main

import (
	"fmt"
)

func main() {
	fileDirPath := flag.String("dir", "./", "directory path for input XML files")
	flag.Parse()
	fmt.Printf("directory for input is %s\n", *fileDirPath)
	ans:=findFileDuplicates(*fileDirPath)
	fmt.Printf("%v\n",ans)
}

func findFileDuplicates(path string) [][]string {
	files := []string{}
	filepath.Walk(*fileDirPath, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() {
			files = append(files, f.Name())
		}
		return nil
	})
	hmCont:=map[string][]string{}
	
	for _, f := range files {
		dat, err := ioutil.ReadFile(f)
		if err != nil {
			panic(err)
		} else {
			hmCont[string(dat)]=append(hmCont[string(dat)],f) 
		}
	}
	ans:=[][]string{}
	for _,v:=range hmCont {
		ans=append(ans,v)
	}
	return ans
}
