package main

import "github.com/go-opencv/go-opencv/opencv"

func main() {
	imagePath := "../data/img/lena.jpg"
	image := opencv.LoadImage(imagePath, opencv.CV_LOAD_IMAGE_COLOR)
	if image == nil {
		panic("empty")
	}

	window := opencv.NewWindow("Image display", opencv.CV_WINDOW_AUTOSIZE)
	window.ShowImage(image)
	opencv.WaitKey(0)
}
