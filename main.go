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

	scanPaths := strings.Split(paths, ",")
	scanKeywords := strings.Split(keywords, ",")

	nfnf := filenamefinder.NewFileNameFinder(scanKeywords)

	for _, path := range scanPaths {
		nfnf.Scan(path)
	}

	for _, file := range nfnf.Files {
		fmt.Println(file)
	}
}
