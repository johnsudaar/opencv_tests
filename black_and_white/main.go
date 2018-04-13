package main

import (
	"log"
	"os"

	"github.com/lazywei/go-opencv/opencv"
)

func main() {
	if len(os.Args) != 3 {
		panic("Usage: main [input] [output]")
	}
	inputPath := os.Args[1]
	outputPath := os.Args[2]

	log.Printf("Taking %s and saving to %s\n", inputPath, outputPath)

	input := opencv.LoadImage(inputPath, opencv.CV_LOAD_IMAGE_COLOR)
	if input == nil {
		panic("Fail to open input file")
	}

	output := opencv.CreateImage(input.Width(), input.Height(), input.Depth(), 1)

	opencv.CvtColor(input, output, opencv.CV_BGR2GRAY)

	opencv.SaveImage(outputPath, output, nil)

	inputWindow := opencv.NewWindow("Base image", opencv.CV_WINDOW_AUTOSIZE)
	outputWindow := opencv.NewWindow("Black and White image", opencv.CV_WINDOW_AUTOSIZE)

	inputWindow.ShowImage(input)
	outputWindow.ShowImage(output)

	opencv.WaitKey(0)

}
