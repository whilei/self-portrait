package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var dirIn string
var dryRun = true

func run() {

	err := filepath.Walk(dirIn, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// check if valid image file
		if !strings.HasSuffix(path, "png") {
			return nil
		}
		if info.IsDir() {
			return nil
		}

		// get date info from filename
		base := filepath.Base(path)
		// form:
		/*
			date
			Sun Mar 21 08:07:22 CDT 2021
			echo "$(date +%y%m%d)-$(date +%H%M).png"
			210321-0806.png
		*/

		// double check eligibility
		if len(base) != len("210321-0806.png") {
			log.Printf("unexpected png doesn't match length: %s", path)
			return nil
		}

		baseDir := filepath.Dir(path)
		yearMonth := base[0:4]
		newbase := base[4:]
		newpath := filepath.Join(baseDir, yearMonth, newbase)

		yearMonthDir := filepath.Dir(newpath)
		yearMonthDirInfo, dirErr := os.Stat(yearMonthDir)
		if dirErr != nil {
			if os.IsNotExist(dirErr) {
				log.Printf("[dry=%v] mkdir: %s", dryRun, yearMonthDir)
				// create it
				if !dryRun {
					if err := os.Mkdir(yearMonthDir, os.ModePerm); err != nil {
						log.Fatalln(yearMonthDir, err)
					}
				}
			} else if !yearMonthDirInfo.IsDir() {
				return fmt.Errorf("not a directory: %s", yearMonthDir)
			} else {
				return err
			}
		}
		// now we have our year/month subdirectory
		log.Printf("[dry=%v] mv: %s -> %s", dryRun, path, newpath)
		if !dryRun {
			if err := os.Rename(path, newpath); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	flag.StringVar(&dirIn, "dirIn", "", "input directory")
	flag.BoolVar(&dryRun, "dryrun", true, "whether to actually create dirs and move files")
	flag.Parse()
	run()
}
