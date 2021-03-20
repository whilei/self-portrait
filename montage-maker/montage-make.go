package main

import (
	"flag"
	"os"
)

var dirIn string
var dirOut string
var filetype string
var harrcascade string
var knownEmptyStore string

func main() {
	home := os.Getenv("HOME")
	if home == "" {
		panic("empty environment variable: HOME")
	}
	flag.StringVar(&dirIn, "dirIn", "/Users/ia/dev/self-portrait/data/examples/originals/", "input directory holding selfies")
	flag.StringVar(&dirOut, "dirOut", "/Users/ia/dev/self-portrait/data/examples/faces/", "output directory")
	flag.StringVar(&filetype, "filetype", ".png", "file type to detect faces, searches input directory")
	flag.StringVar(&harrcascade, "harrcascade", "/Users/ia/gocode/src/github.com/lazywei/go-opencv/samples/haarcascade_frontalface_alt.xml", "harrcascade thing")
	flag.StringVar(&knownEmptyStore, "empty", "/Users/ia/dev/self-portrait/data/knownnofaces", "file in which to store list of known no-face images")

	flag.Parse()
}
