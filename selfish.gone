package main

import (
	"bytes"
	"fmt"
	"log"
	"path/filepath"

	difflib "github.com/aryann/difflib"
	// "os"
	// "os/exec"
)

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

	fmt.Printf("--> Found %d unique files.\n", len(uniques))

	return uniques
}

func main() {
	originalsDir := "/Users/ia/dev/self-portrait/originals_examples/"
	facesDir := "/Users/ia/dev/self-portrait/faces_examples/"
	filetype := ".png"
	getUniqueOriginals(originalsDir, facesDir, filetype)
}
