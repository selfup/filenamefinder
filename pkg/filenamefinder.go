package filenamefinder

import (
	"io/ioutil"
	"os"
	"runtime"
	"strings"
	"sync"
)

// NewFileNameFinder creates a pointer to FileSizeFinder with default values
func NewFileNameFinder(keywords []string) *FileSizeFinder {
	fsf := new(FileSizeFinder)

	if runtime.GOOS == "windows" {
		fsf.Direction = "\\"
	} else {
		fsf.Direction = "/"
	}

	fsf.Keywords = keywords

	return fsf
}

// FileSizeFinder struct contains needed data to perform concurrent operations
type FileSizeFinder struct {
	mutex     sync.Mutex
	Direction string
	Files     []string
	Keywords  []string
}

// Scan is a concurrent/parallel directory walker
func (f *FileSizeFinder) Scan(directory string) {
	_, err := ioutil.ReadDir(directory)
	if err != nil {
		panic(err)
	}

	f.findFiles(directory, "")
}

// The prefix is needed when the goroutines fire off
func (f *FileSizeFinder) findFiles(directory string, prefix string) {
	paths, _ := ioutil.ReadDir(directory)

	var dirs []os.FileInfo
	var files []os.FileInfo

	for _, path := range paths {
		if path.IsDir() {
			dirs = append(dirs, path)
		} else {
			files = append(files, path)
		}
	}

	for _, file := range files {
		for _, keyword := range f.Keywords {
			if strings.Contains(file.Name(), keyword) {
				f.mutex.Lock()
				f.Files = append(f.Files, directory+f.Direction+file.Name())
				f.mutex.Unlock()
			}
		}
	}

	dirLen := len(dirs)

	if dirLen > 0 {
		var dirGroup sync.WaitGroup
		dirGroup.Add(dirLen)

		for _, dir := range dirs {
			go func(localDirectory os.FileInfo, localPrefix string, osDirection string) {
				f.findFiles(localPrefix+osDirection+localDirectory.Name(), localPrefix)
				dirGroup.Done()
			}(dir, directory, f.Direction)
		}

		dirGroup.Wait()
	}
}
