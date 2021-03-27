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
var fileOut string

func run() {

	convertResizeStaticArgs := []string{
		"-background", "transparent",
		"-gravity", "center",
		"-resize", "640x480",
	}

	resizeTmpFile := filepath.Join(os.TempDir(), "resized_tmp_file.png")

	i := 0
	if err := filepath.Walk(dirIn, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if ext := filepath.Ext(path); !strings.HasSuffix(ext, "png") {
			log.Printf("Skipping (not .png): %s\n", path)
			return nil
		}
		defer func() { i++ }()

		procResizeArgs := append([]string{path}, convertResizeStaticArgs...)

		// Normally use a temp file as the resize target.
		resizeTarget := resizeTmpFile
		if i == 0 {
			// But for the first image, we'll initialize the average file directly.
			resizeTarget = fileOut
		}

		procResizeArgs = append(procResizeArgs, resizeTarget)

		// Resize the image to standard dimensions.
		procResize := exec.Command("convert", procResizeArgs...)

		out, err := procResize.CombinedOutput()
		if err != nil {
			log.Printf("resize error: %v", err)
			return err
		} else {
			log.Printf("resize OK: %s %s", path, string(out))
		}

		// For the first file, this is all we need to do.
		if i == 0 {
			return nil
		}
		defer os.Remove(resizeTarget)

		// For all later files, run the averaging command between the
		// target image file and the standard-resized temporary image.
		procFx := exec.Command("convert", resizeTarget, fileOut,

			// This flag is magic.
			// I don't remember what its doing, just working from my shell script as a reference.
			"-fx", fmt.Sprintf(`(u+%d*v)/%d`, i, i+1),
			fileOut,
			)

		out, err = procFx.CombinedOutput()
		if err != nil {
			log.Printf("fx error: %v", err)
			return err
		} else {
			log.Printf("fx OK: %s %s", path, string(out))
		}

		return nil
	}); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	flag.StringVar(&dirIn, "dirIn", "", "input directory holding images to montage")
	flag.StringVar(&fileOut, "fileOut", "", "output file for average image (parent dir will be created if not existing)")

	flag.Parse()

	if dirIn == "" {
		log.Fatal("-dirIn cannot be empty")
	}
	if fileOut == "" {
		log.Fatal("-fileOut cannot be empty")
	}

	dirOut := filepath.Dir(fileOut)
	if fi, err := os.Stat(dirOut); os.IsNotExist(err) {
		if err := os.MkdirAll(dirOut, os.ModePerm); err != nil {
			panic(err)
		}
	} else if err != nil {
		panic(err)
	} else if !fi.IsDir() {
		panic("is not a directory: " + dirOut)
	}

	run()
}
