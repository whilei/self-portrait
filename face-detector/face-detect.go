package main

// lil bullet and detectomatic selfies
// TODO: get it a cron to do it on the server

import (
	"bufio"
	"flag"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/aryann/difflib"
	"github.com/oliamb/cutter"
	"gocv.io/x/gocv"
)

var dirIn string
var dirOut string
var filetype string
var harrcascade string
var knownEmptyStore string

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
func getFiles(dirIn string, filetype string) (files []string) {
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
	return filepath.Base(name)
	// return string(bytes.Replace([]byte(filepath.Base(name)), []byte("face_"), []byte(""), 1)) // remove 'face_' prefix
}

func getDifferentFiles(listA, listB []string) (diff []difflib.DiffRecord) {
	return difflib.Diff(listA, listB)
}

func getUniqueOriginals(dirIn, dirOut, filetype string) []string {

	// get file list from originals
	originals := getFiles(dirIn, filetype /*eg. .png*/)
	faces := getFiles(dirOut, filetype /*eg. .png*/)
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
			uniqueAbsPath := filepath.Join(dirIn, diff.Payload)
			fmt.Printf("%s\n", uniqueAbsPath)
			uniques = append(uniques, uniqueAbsPath)
		}
	}

	fmt.Printf("Found %d unique files.\n", len(uniques))

	return uniques
}

func readFileLinesToStrSlice(fpath string) (out []string) {
	file, err := os.Open(fpath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" || len(scanner.Text()) < 4 {
			continue
		}
		out = append(out, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return out
}

var faceCropScale = 33 // percent bigger
func enlargeCrop(rect image.Rectangle, maxCols, maxRows int) (nanchor image.Point, ncols, nrows int) {
	width, height := (rect.Max.X - rect.Min.X), (rect.Max.Y - rect.Min.Y)
	ncols = width * (100 + faceCropScale) / 100
	nrows = width * (100 + faceCropScale + 62) / 100 // golden ratio

	// adjust anchor (top left == rect Min) given scaled rect size
	x, y := rect.Min.X, rect.Min.Y
	x = x - ((ncols - width) / 2)
	y = y - ((nrows - height) / 2)

	// ensure fit within max bounds
	if x < 0 {
		x = 0
	}
	if y < 0 {
		y = 0
	}
	if x+ncols > maxCols {
		ncols = maxCols - x
	}
	if y+nrows > maxRows {
		nrows = maxRows - y
	}
	nanchor = image.Point{
		X: x,
		Y: y,
	}
	return
}

//detects faces and crops em out
func cropFaces(inputs []string, dirOut string, harrcascade string) {
	err := os.MkdirAll(dirOut, os.ModePerm) // makes dir if not exists
	if err != nil {
		fmt.Printf("Could not create directory %s\n", dirOut)
		log.Fatal(err)
	}

	fmt.Printf("Cropping %d images.\n", len(inputs))
	//detect some faces
	// load classifier to recognize faces
	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()

	if !classifier.Load(harrcascade) {
		fmt.Printf("Error reading cascade file: %v\n", harrcascade)
		return
	}
	// track files without faces to check against
	// removes a string from string slice if it exists in the slice
	// this "widdles" the size of the in-memory known-no-faces list as previously checked files are referenced
	// Since order is preserved, it should be relatively fast.
	spliceIfContains := func(sl []string, s string) ([]string, bool) {
		for i, ss := range sl {
			if ss == s {
				sl = append(sl[:i], sl[i+1:]...)
				return sl, true
			}
		}
		return sl, false
	}
	nofaceslistfilePath := filepath.Clean(knownEmptyStore)
	var nofaces []string
	if _, e := os.Stat(nofaceslistfilePath); e == nil {
		nofaces = readFileLinesToStrSlice(nofaceslistfilePath)
	} else if os.IsNotExist(e) {
		// simulate touch
		if f, e := os.Create(nofaceslistfilePath); e != nil {
			panic(e)
		} else {
			f.Close()
		}
	}
	// // ugly roll for the nofaces cache file
	// defer func() {
	// 	tail := exec.Command("tail", "-3000", nofaceslistfilePath)
	// 	tmpFile, err := ioutil.TempFile(os.TempDir(), "no-face-tmp")
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	// 	stdout, err := tail.StdoutPipe()
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	// 	if err := tail.Start(); err != nil {
	// 		log.Fatalln(err)
	// 	}
	// 	buf := bytes.NewBuffer([]byte{})
	// 	_, err = buf.ReadFrom(stdout)
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	// 	err = ioutil.WriteFile(tmpFile.Name(), buf.Bytes(), os.ModePerm)
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	// 	if err := tail.Wait(); err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	stdout.Close()
	// 	os.Rename(tmpFile.Name(), nofaceslistfilePath)
	// }()

	nofaceFile, err := os.OpenFile(nofaceslistfilePath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer nofaceFile.Close()

	l := len(inputs)
	for i, element := range inputs {

		// outPath := dirOut + "face_" + filepath.Base(element)
		outPath := filepath.Join(dirOut, filepath.Base(element))
		fmt.Println(i+1, "/", l, ":", element)

		var cachedKnownNoFace bool
		nofaces, cachedKnownNoFace = spliceIfContains(nofaces, element)
		if cachedKnownNoFace {
			log.Println("known no-face, skipping")
			continue
		}

		imageMat := gocv.IMRead(element, gocv.IMReadColor)
		if imageMat.Empty() {
			log.Println("empty image, skipping", element)
			if _, has := spliceIfContains(nofaces, element); !has {
				nofaceFile.WriteString(element + "\n")
			}
			continue
		}

		// check to see if completely black image.
		// sometimes this happens because my computer doesn't go to sleep correctly when connected to external monitor
		// ... and sometimes it's just dark
		// scal := imageMat.Mean()
		// if scal.Val1 < 10 && scal.Val2 < 10 && scal.Val3 < 10 {
		// 	log.Println("dark image, skipping")
		// 	noface.WriteString(element)
		// 	continue
		// }

		rects := classifier.DetectMultiScale(imageMat)
		if len(rects) == 0 {
			log.Println("no faces detected")
			if _, has := spliceIfContains(nofaces, element); !has {
				nofaceFile.WriteString(element + "\n")
			}
			continue
		}

		fs, err := os.Open(element)
		if err != nil {
			panic(err)
		}
		pngF, err := png.Decode(fs)
		if err != nil {
			panic(err)
		}
		for _, rect := range rects {
			if rect.Empty() {
				log.Println("no rect")
				continue
			}
			log.Println("rect bounds", rect.Bounds().String())
			face := imageMat.Region(rect)
			if face.Empty() {
				log.Println("empty face")
				continue
			}
			// don't chop the chin off
			a, w, h := enlargeCrop(rect, imageMat.Cols(), imageMat.Rows())
			croppedImg, err := cutter.Crop(pngF, cutter.Config{
				Width:  w,
				Height: h,
				Anchor: a,
				Mode:   cutter.TopLeft, // optional, default value
			})
			if err != nil {
				panic(err)
			}
			toimg, _ := os.Create(outPath)
			if err := png.Encode(toimg, croppedImg); err != nil {
				log.Println("err encoding png", err)
			} else {
				log.Println("->", outPath)
			}
			toimg.Close()
			if err != nil {
				log.Println("err writing face bytes", err)
				os.Remove(outPath)
			}
		}
	}
}

func main() {

	cmd := os.Args[0]

	fmt.Printf("Program Name: %s\n", cmd)

	flag.StringVar(&dirIn, "dirIn", "/Users/ia/dev/self-portrait/data/examples/originals/", "input directory holding selfies")
	flag.StringVar(&dirOut, "dirOut", "/Users/ia/dev/self-portrait/data/examples/faces/", "output directory, will be create-alled if DNE")
	flag.StringVar(&filetype, "filetype", ".png", "file type to detect faces, searches input directory")
	flag.StringVar(&harrcascade, "harrcascade", "/Users/ia/gocode/src/github.com/lazywei/go-opencv/samples/haarcascade_frontalface_alt.xml", "harrcascade thing")

	// This file acts as a semi-persistent cache to avoid checking the same image for faces twice.
	flag.StringVar(&knownEmptyStore, "cache-nofacelist", filepath.Join(os.TempDir(), "face-detector-nofacelist"), "file in which to store list of known no-face images")

	flag.Parse()

	// Sanity
	if dirIn == dirOut {
		log.Fatalln("dirIn cannot also be dirOut")
	}

	// Sanitize
	if !strings.HasSuffix(dirIn, string(filepath.Separator)) {
		dirIn += string(filepath.Separator)
	}

	cropFaces(getUniqueOriginals(dirIn, dirOut, filetype), dirOut, harrcascade)

}
