package main

import (
	"flag"
	"fmt"
	"strings"

	filenamefinder "github.com/selfup/filenamefinder/pkg"
)

func main() {
	var paths string
	flag.StringVar(&paths, "path", "", "absolute path - can be comma delimited")

	var keywords string
	flag.StringVar(&keywords, "keywords", "", "keyword(s) for the filename - can be comma delimited")

	flag.Parse()

	scanKeywords := strings.Split(keywords, ",")
	scanPaths := strings.Split(paths, ",")

	nfsf := filenamefinder.NewFileNameFinder(scanKeywords)

	for _, path := range scanPaths {
		nfsf.Scan(path)
	}

	for _, file := range nfsf.Files {
		fmt.Println(file)
	}
}
