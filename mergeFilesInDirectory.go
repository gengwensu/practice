/**
 * Merge k files with lexicographically sorted files and return the result as one sorted output file.
 * Analyze and describe its time and space complexity.
**/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

var fileDirPath, outFile *string

func main() {
	fileDirPath = flag.String("dir", "./test", "directory path for input files")
	outFile = flag.String("outFile", "outFile.txt", "output file")
	flag.Parse()
	fmt.Printf("directory for input is %s, outFile is %s\n", *fileDirPath, *outFile)
	if files := getFilesInDirectory(*fileDirPath); len(files) > 0 {
		for i := range files {
			if !mergeInToOut(files[i], *outFile) { // merge one by one and output to file to minimize space
				log.Fatalf("merge fails")
			}
			// fmt.Printf("Done with %s\n", files[i])
		}
	}
	fmt.Println("files merged successfully!")
}

/**
* get filenames in a list and check to make sure each file is sorted
**/
func getFilesInDirectory(path string) []string {
	files := []string{}
	filepath.Walk(*fileDirPath, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() {
			files = append(files, f.Name())
		}
		return nil
	})
	if err := os.Chdir(path); err != nil {
		log.Fatalf("failed to change directory: %s\n", err)
	}
	for _, f := range files {
		file, err := os.Open(f)
		if err != nil {
			log.Fatalf("failed opening file: %s\n", err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		var prevLine string
		for scanner.Scan() { // using line by line scan to handle large files
			textLine := scanner.Text() // read in a lines
			if len(textLine) > 0 {     // skip empty line
				if textLine < prevLine { // out of lexicographic order
					log.Fatalf("%s is not sort!\n")
				} else {
					prevLine = textLine
				}
			}
		}
	}
	return files
}

/**
* merge infile with lines in lexico. order to outFile with lines in lexico. order.
* contains no empty lines or duplicated lines
**/
func mergeInToOut(inf, outf string) bool {
	if len(inf) == 0 { // nothing to merge
		return true
	}

	tempfile, err := os.Create("temp.txt") // open a temp file for result
	if err != nil {
		log.Fatalf("failed opening temp file: %s\n", err)
	}
	os.Remove("temp.txt")

	in, err := os.Open(inf)
	if err != nil {
		log.Fatalf("failed opening file: %s\n", err)
	}
	defer in.Close()

	out, err := os.OpenFile(outf, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatalf("failed opening file: %s\n", err)
	}

	inScanner := bufio.NewScanner(in)
	inScanner.Split(bufio.ScanLines)
	outScanner := bufio.NewScanner(out)
	outScanner.Split(bufio.ScanLines)

	var inLine, outLine string
	for inScanner.Scan() {
		if inLine = inScanner.Text(); len(inLine) > 0 { // get the first non-empty line
			break
		}
	}
	if len(inLine) == 0 { // inf is empty, no need to merge
		tempfile.Close()
		os.Remove("temp.txt")
		//fmt.Println("temp.txt removed")
		out.Close()
		return true
	}
	for outScanner.Scan() {
		if outLine = outScanner.Text(); len(outLine) > 0 { // get the first non-empty line
			break
		}
	}

	//fmt.Printf("before loop, in - %s, out - %s\n", inLine, outLine)
	var prevLine string // previous line written to temp
	for len(inLine) > 0 || len(outLine) > 0 {
		//fmt.Printf("before switch, in - %s, out - %s\n", inLine, outLine)
		switch {
		case len(inLine) > 0 && len(outLine) > 0:
			if inLine <= outLine {
				if inLine < outLine { // output if not equal
					if inLine != prevLine { // de-dup
						if _, err := tempfile.Write([]byte(inLine + "\n")); err != nil {
							log.Fatal(err)
						}
						//fmt.Printf("%s write to tempfile\n", inLine)
					}
					prevLine = inLine
				} // skip output if equal

				inLine = ""
				for inScanner.Scan() { // read next line
					if inLine = inScanner.Text(); len(inLine) > 0 { // get the next non-empty line
						break
					}
				}
			} else {
				if outLine != prevLine { // de-dup
					if _, err := tempfile.Write([]byte(outLine + "\n")); err != nil {
						log.Fatal(err)
					} else {
						//fmt.Printf("%s write to tempfile\n", outLine)
						prevLine = outLine
						outLine = ""
						if outScanner.Scan() { // read next line from outFile
							outLine = outScanner.Text()
						} // otherwise at EOF
					}
				}
			}

		case len(inLine) == 0 && len(outLine) > 0:
			if _, err := tempfile.Write([]byte(outLine + "\n")); err != nil {
				log.Fatal(err)
			} else { // output the rest of outFile
				for outScanner.Scan() {
					outLine = outScanner.Text()
					if _, err := tempfile.Write([]byte(outLine + "\n")); err != nil {
						log.Fatal(err)
					}
				}
			}
			outLine = ""

		case len(inLine) > 0 && len(outLine) == 0:
			if inLine != prevLine {
				if _, err := tempfile.Write([]byte(inLine + "\n")); err != nil {
					log.Fatal(err)
				}
				//fmt.Printf("%s write to tempfile\n", inLine)
				prevLine = inLine
			}
			for inScanner.Scan() { // output the rest of inFile
				inLine = inScanner.Text()
				//fmt.Printf("next line is %s\n", inLine)
				if len(inLine) > 0 { // skip empty lines
					if inLine != prevLine {
						if _, err := tempfile.Write([]byte(inLine + "\n")); err != nil {
							log.Fatal(err)
						}
						//fmt.Printf("%s write to tempfile\n", inLine)
						prevLine = inLine
					}
				}
			}
			inLine = ""
		}
	}

	out.Close()
	tempfile.Close()

	tempfile, err = os.Open("temp.txt") // open temp file again
	if err != nil {
		log.Fatalf("failed opening temp file: %s.\n", err)
	}

	out, err = os.OpenFile(outf, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatalf("failed opening file: %s.\n", err)
	}
	defer out.Close()

	if _, err = io.Copy(out, tempfile); err != nil {
		log.Fatalf("Fail to replace the outFile")
	}

	//fmt.Println("temp.txt removed")
	return true
}
