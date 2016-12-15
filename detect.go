package main

// lil bullet and detectomatic selfies
// TODO: get it a cron to do it on the server

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	difflib "github.com/aryann/difflib"
	opencv "github.com/lazywei/go-opencv/opencv"
)

// Exists reports whether the named file or directory exists.
// apparently this can be wrong if permissions or something else thangle with it, would say true when not
func exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

//Find the files in input directory wanted
func getFiles(dirIn string, filetype string) []string {
	fmt.Printf("Finding %s files in %s\n", filetype, dirIn)
	files, err := filepath.Glob(dirIn + "*" + filetype)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found %d files\n", len(files))
	return files
}

// remove abs path and "face_" from either/both original and face file names
func stripFileName(name string) string {
	return string(bytes.Replace([]byte(filepath.Base(name)), []byte("face_"), []byte(""), 1)) // remove 'face_' prefix
}

func getDifferentFiles(listA, listB []string) (diff []difflib.DiffRecord) {
	return difflib.Diff(listA, listB)
}

func getUniqueOriginals(dirIn, dirOut, filetype string) []string {

	// get file list from originals
	originals := getFiles(dirIn, ".png")
	faces := getFiles(dirOut, ".png")
	var originalsStripped = make([]string, len(originals))
	var facesStripped = make([]string, len(faces))

	// strip file names, removing "face_" from faces
	// and abs path from both
	for i, faceFileName := range faces {
		facesStripped[i] = stripFileName(faceFileName)
	}
	for j, originalsFileName := range originals {
		originalsStripped[j] = stripFileName(originalsFileName)
	}

	//  Constant    Code   Meaning
	// ----------  ------ ---------------------------------------
	//  Common      " "    The element occurs in both sequences.
	//  LeftOnly    "-"    The element is unique to sequence 1.
	//  RightOnly   "+"    The element is unique to sequence 2.
	diffs := getDifferentFiles(originalsStripped, facesStripped)
	fmt.Printf("Diffing...\n")
	var uniques []string

	for _, diff := range diffs {
		// fmt.Printf("%s\n", diff)
		// should only have LeftOnly's, ie only have orignals that are not in faces'
		if diff.Delta == difflib.LeftOnly {
			uniqueAbsPath := dirIn + diff.Payload
			fmt.Printf("%s\n", uniqueAbsPath)
			uniques = append(uniques, uniqueAbsPath)
		}
	}

	fmt.Printf("Found %d unique files.\n", len(uniques))

	return uniques
}

//detects faces and crops em out
func cropFaces(inputs []string, dirOut string, harrcascade string) {
	err := os.MkdirAll(dirOut, 0777) // makes dir if not exists
	if err != nil {
		fmt.Printf("Could not create directory %s\n", dirOut)
		log.Fatal(err)
	}

	fmt.Printf("Cropping %d images.\n", len(inputs))

	for _, element := range inputs {

		outPath := dirOut + "face_" + filepath.Base(element)

		image := opencv.LoadImage(element)
		if image == nil {
			panic("Loading Image " + element + "failed")
		}
		defer image.Release()

		//detect some faces
		cascade := opencv.LoadHaarClassifierCascade(harrcascade)
		faces := cascade.DetectObjects(image)
		if len(faces) == 0 {
			fmt.Printf("Found no face in %s\n", element)
		}

		for _, value := range faces {
			//Some invalid stuff filtering
			if value.Width() < image.Width() && value.Height() < image.Height() && value.Width() > 0 && value.Height() > 0 {
				//crop out the face
				crop := opencv.Crop(image, value.X(), value.Y(), value.Width(), value.Height())
				//save cropped
				opencv.SaveImage(outPath, crop, 0)
				crop.Release()
			} else {
				fmt.Printf("Couldn't save: %s", value)
			}
		}
	}
}

func main() {

	cmd := os.Args[0]

	fmt.Printf("Program Name: %s\n", cmd)

	var dirIn string
	var dirOut string
	var filetype string
	var harrcascade string

	flag.StringVar(&dirIn, "dirIn", "/Users/ia/dev/self-portrait/originals_examples/", "input directory holding selfies")
	flag.StringVar(&dirOut, "dirOut", "/Users/ia/dev/self-portrait/faces_examples/", "output directory")
	flag.StringVar(&filetype, "filetype", ".png", "file type to detect faces, searches input directory")
	flag.StringVar(&harrcascade, "harrcascade", "/Users/ia/gocode/src/github.com/lazywei/go-opencv/samples/haarcascade_frontalface_alt.xml", "harrcascade thing")

	flag.Parse()

	cropFaces(getUniqueOriginals(dirIn, dirOut, filetype), dirOut, harrcascade)

}
