package main

// lil bullet and detectomatic selfies
import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/lazywei/go-opencv/opencv"
)

//Find the files in input directory wanted
func getFiles(dirIn string, filetype string) []string {
	fmt.Printf("Finding .png files in %s\n", dirIn)
	files, err := filepath.Glob(dirIn + "*" + filetype)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found %d files", len(files))
	return files
}

//detects faces and crops em out
func cropFaces(inputs []string, dirOut string, harrcascade string) {
	err := os.MkdirAll(dirOut, 0777)
	if err != nil {
		fmt.Printf("Could not create directory %s", dirOut)
		log.Fatal(err)
	}
	for _, element := range inputs {
		image := opencv.LoadImage(element)
		if image == nil {
			panic("Loading Image " + element + "failed")
		}
		defer image.Release()

		//detect some faces
		cascade := opencv.LoadHaarClassifierCascade(harrcascade)

		faces := cascade.DetectObjects(image)
		for _, value := range faces {
			//Some invalid stuff filtering
			if value.Width() < image.Width() && value.Height() < image.Height() && value.Width() > 0 && value.Height() > 0 {
				//crop out the face
				crop := opencv.Crop(image, value.X(), value.Y(), value.Width(), value.Height())
				//save cropped
				opencv.SaveImage(dirOut+"face_"+filepath.Base(element), crop, 0)
				crop.Release()
			} else {
				fmt.Println(value)
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

	flag.StringVar(&dirIn, "dirIn", "./selfless/", "input directory holding selfies")
	flag.StringVar(&dirOut, "dirOut", "./selfmore/", "output directory")
	flag.StringVar(&filetype, "filetype", ".png", "file type to detect faces, searches input directory")
	flag.StringVar(&harrcascade, "harrcascade", "/Users/Kitty/Go/src/github.com/lazywei/go-opencv/samples/haarcascade_frontalface_alt.xml", "harrcascade thing")

	flag.Parse()

	cropFaces(getFiles(dirIn, filetype), dirOut, harrcascade)

}
