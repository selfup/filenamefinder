# File Name Finder

Concurrent, recursive, file name finder.

Can search for multiple keywords in file names.

Pipe friendly so you can grep or xargs all day.

### Install

```
go get github.com/selfup/filenamefinder
go install github.com/selfup/filenamefinder
```

### Use

```
$ filenamefinder -h
Usage of filenamefinder:
  -k string
        keyword(s) for the filename - can be comma delimited
  -p string
        absolute path - can be comma delimited
```

Example looking for all README files in $HOME and doing a count of matches:

```
$ time filenamefinder -k='README' -p="$HOME" | wc -l
4815

real    0m1.389s
user    0m4.960s
sys     0m2.708s
```

---

Performance is quite similar to `find` but this is such little code I figured why not.

You can also use this in your go projects and not have to exec a shell command.

That's nice too.

### Use as a lib

```go
import (
    "fmt"
    "strings"

    filenamefinder "github.com/selfup/filenamefinder/pkg"
)

scankeywords := strings.Split("first_keyword,second_keyword", ",")
scanPaths := strings.Split("/tmp,/etc,/home", ",")

nfsf := filenamefinder.NewFileNameFinder(scanKeywords)

for _, path := range scanPaths {
    nfsf.Scan(path)
}

for _, file := range nfsf.Files {
    fmt.Println(file)
}
```

### Caveats

Currently fuzzy find which I prefer. The rest can be narrowed down with grep or any other pipe freindly util.
