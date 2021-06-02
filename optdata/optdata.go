package optdata

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func ExtractOpdataLogs(basedirectory string, outputpath string) ([]string, error) {

	archiveFP := filepath.Join(basedirectory, ".")

	fileinfo, err := ioutil.ReadDir(archiveFP)
	if err != nil {
		return nil, err
	}

	var targzfiles []string
	targzfiles = make([]string, 0)

	basepath, err := filepath.Abs(archiveFP)
	if err != nil {
		log.Fatalln(err)
	}

	for _, f := range fileinfo {
		if f.IsDir() == false {
			//fmt.Println(f.Name())
			//fmt.Println(len(f.Name()))
			if len(f.Name()) > 6 {
				end := f.Name()[(len(f.Name()) - 6):]
				if end == "tar.gz" {

					ff := filepath.Join(basepath, f.Name())
					targzfiles = append(targzfiles, ff)

				}
			}

		}
	}

	fmt.Println(targzfiles)

	fmt.Println(outputpath)
	err = os.Chdir(outputpath)
	if err != nil {
		log.Fatalln(err)
	}

	wd, _ := os.Getwd()
	fmt.Println(wd)
	fmt.Println("here")

	for _, targzfile := range targzfiles {

		fmt.Println(targzfile)

		f, err := os.Open(targzfile)
		if err != nil {
			log.Fatal(err)
		}
		ExtractTarGz(f)

	}

	return targzfiles, nil

}
