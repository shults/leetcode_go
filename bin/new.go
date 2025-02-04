package main

import (
	"fmt"
	"os"
	"path"
	"strconv"
)

const (
	minRange = 1
	maxRange = 99999
)

func main() {
	if len(os.Args) != 2 {
		printAndExit(nil)
	}

	taskNumber, err := strconv.Atoi(os.Args[1])
	if err != nil {
		printAndExit(err)
	}

	if isOutOfRange(taskNumber) {
		printAndExit(fmt.Errorf("expected task number to be in range [%d;%d]", minRange, maxRange))
	}

	pkgName := fmt.Sprintf("t%05d", taskNumber)

	if err = os.Mkdir(pkgName, 0775); err != nil {
		printAndExit(fmt.Errorf("unable to create tast: %s", err))
	}

	for _, f := range []fData{
		fileData(pkgName),
		testFileData(pkgName),
	} {
		if err = os.WriteFile(f.path, f.content, 0664); err != nil {
			printAndExit(fmt.Errorf("failed to crate file %s: %s", f.path, err))
		}
	}

}

func printAndExit(err error) {
	fmt.Printf("Usage:\n")
	fmt.Printf("[program] <nr_of_task>\n")

	if err != nil {
		fmt.Printf("error: %s\n", err)
	}

	os.Exit(1)
}

func isOutOfRange(val int) bool {
	return val < minRange || val > maxRange
}

func fileData(pkgName string) fData {
	return fData{
		path:    path.Join(pkgName, pkgName+".go"),
		content: []byte(fmt.Sprintf("package %s\n", pkgName)),
	}
}

func testFileData(pkgName string) fData {
	return fData{
		path:    path.Join(pkgName, pkgName+"_test.go"),
		content: []byte(fmt.Sprintf("package %s\n", pkgName)),
	}
}

type fData struct {
	content []byte
	path    string
}
