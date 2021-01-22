package gols

import (
	"flag"
	"io/ioutil"
	"os"
	"regexp"

	"github.com/fatih/color"
)

func Run() error {
	allFile := flag.Bool("a", false, "include directory entries whose names with a dot (.)")
	flag.Parse()

	files, err := ioutil.ReadDir(".")
	if err != nil {
		return err
	}

	if *allFile == false {
		tmpFiles := make([]os.FileInfo, 0, len(files))
		for _, file := range files {
			name := file.Name()
			match, _ := regexp.MatchString(`^\w`, name)
			if match == true {
				tmpFiles = append(tmpFiles, file)
			}
		}
		files = tmpFiles
	}

	for _, file := range files {
		if file.IsDir() {
			color.Cyan(file.Name())
		} else {
			color.Yellow(file.Name())
		}
	}

	return nil
}
