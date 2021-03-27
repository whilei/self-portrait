package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

var dirIn string
var montageMax = 64 // will be calculated from tile configuration (eg. 8 * 8 = 64)
var montageTile = "8x8"
var montageGeo = "32x32+1+1"
var outFilePrefix string

func montage(files []string, montageIndex int) {
	args := files
	args = append(args, "-tile", montageTile)
	args = append(args, "-geometry", montageGeo)

	// This is how the resulting montage file will be named.
	outFile := fmt.Sprintf("%s_%s_%d.png", outFilePrefix, montageTile, montageIndex)

	args = append(args, outFile)
	proc := exec.Command("montage", args...)
	out, err := proc.CombinedOutput()
	if err != nil {
		log.Printf( "montage errored: %s err=%v files(%d)=%v\n", outFile, err, len(files), files)
	} else {
		fmt.Printf("ok: %s %s\n", outFile, string(out))
	}
}

func mustCalculateMontageMax(tileParam string) int {
	xy := strings.Split(tileParam, "x")
	if len(xy) != 2 {
		panic("tile config must be delimited by a single 'x'")
	}
	x := xy[0]
	y := xy[1]
	xx, err := strconv.ParseInt(x, 10, 16)
	if err != nil {
		panic(err)
	}
	yy, err := strconv.ParseInt(y, 10, 16)
	if err != nil {
		panic(err)
	}
	return int(xx * yy)
}

func run() {
	montageMax = mustCalculateMontageMax(montageTile)

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
		if ext := filepath.Ext(path); !strings.HasSuffix(ext, "png") {
			log.Printf("Skipping (not .png): %s\n", path)
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
	// flag.IntVar(&montageMax, "montage-max", 64, "Max number of images per montage")
	flag.StringVar(&montageTile, "montage-tile", "8x8", "Tile dimensions for montage")
	flag.StringVar(&montageGeo, "montage-geo", "64x64+1+1", "Geometry for each montage image")
	flag.StringVar(&outFilePrefix, "montage-file-pre", "montage", "File prefix for each montaged file (eg. montage-prefix-1.png, montage-prefix-2.png)")

	flag.Parse()

	run()
}
