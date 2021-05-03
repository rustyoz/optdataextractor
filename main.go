package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	optdata "github.com/rustyoz/optdataextractor/optdata"
)

func main() {

	optdatapath := flag.String("d", "test", "location of optdata folder")
	outputpath := flag.String("o", "archive", "output folder")

	if filepath.IsAbs(*outputpath) == false {
		wd, _ := os.Getwd()
		*outputpath = filepath.Join(wd, *outputpath)
	}

	flag.Parse()

	_, err := optdata.ExtractOpdataLogs(*optdatapath, *outputpath)

	if err != nil {
		log.Fatalln(err)
	}

}
