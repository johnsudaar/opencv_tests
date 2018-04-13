package main

import (
	"log"

	"github.com/lazywei/go-opencv/opencv"
)

func main() {
	webcam := opencv.NewCameraCapture(0)
	if webcam == nil {
		panic("Fail to open camera")
	}

	window := opencv.NewWindow("Webcam", opencv.CV_WINDOW_AUTOSIZE)
	for {
		curFrame := webcam.QueryFrame()
		if curFrame == nil {
			log.Println("No capture, skipping")
		}
		window.ShowImage(curFrame)

		// WTF 1048603 ???
		if opencv.WaitKey(10) == 1048603 {
			break
		}
	}
	log.Printf("Exiting")
}
