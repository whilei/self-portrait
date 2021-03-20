package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var dirIn string
var montageMax = 64
var montageTile = "8x8"
var montageGeo = "32x32+1+1"
var outFilePrefix string

func montage(files []string, montageIndex int) {
	args := files
	args = append(args, "-tile", montageTile)
	args = append(args, "-geometry", montageGeo)
	args = append(args, fmt.Sprintf("%s%d.png", outFilePrefix, montageIndex))
	proc := exec.Command("montage", args...)
	out, err := proc.CombinedOutput()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%s\n", string(out))
}

func run() {
	index := 1
	batch := []string{}
	runMontage := func() {
		if len(batch) == 0 {
			return
		}
		montage(batch, index)
		batch = []string{}
		index++
	}
	if err := filepath.Walk(dirIn, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if ext := filepath.Ext(path); !strings.Contains(ext, "png") {
			log.Println("Skipping (not .png): %s", path)
			return nil
		}
		batch = append(batch, path)
		if len(batch) >= montageMax {
			runMontage()
		}
		return nil
	}); err != nil {
		log.Fatalln(err)
	}
	runMontage()
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	flag.StringVar(&dirIn, "dirIn", "", "input directory holding images to montage")
	flag.IntVar(&montageMax, "montage-max", 64, "Max number of images per montage")
	flag.StringVar(&montageTile, "montage-tile", "8x8", "Tile dimensions for montage")
	flag.StringVar(&montageGeo, "montage-geo", "64x64+1+1", "Geometry for each montage image")
	flag.StringVar(&outFilePrefix, "montage-file-pre", "montage-test-", "File prefix for each montaged file (eg. montage-prefix-1.png, montage-prefix-2.png)")

	flag.Parse()

	run()
}
