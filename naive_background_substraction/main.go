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

	curFWindow := opencv.NewWindow("Cur", opencv.CV_WINDOW_AUTOSIZE)

	log.Println("Press key when empty background")
	for {
		curFWindow.ShowImage(webcam.QueryFrame())

		if opencv.WaitKey(10) != -1 {
			break
		}

	}
	firstFrameColor := webcam.QueryFrame()
	firstFrame := opencv.CreateImage(firstFrameColor.Width(), firstFrameColor.Height(), firstFrameColor.Depth(), 1)
	opencv.CvtColor(firstFrameColor, firstFrame, opencv.CV_BGR2GRAY)
	ffWindow := opencv.NewWindow("BG", opencv.CV_WINDOW_AUTOSIZE)
	ffWindow.ShowImage(firstFrame)

	xWindow := opencv.NewWindow("XORED", opencv.CV_WINDOW_AUTOSIZE)
	thVal := float64(38)
	xWindow.CreateTrackbar("Value", int(thVal), 255, func(pos int) {
		thVal = float64(pos)
	})

	for {
		curFrameColor := webcam.QueryFrame()
		curFrame := opencv.CreateImage(curFrameColor.Width(), curFrameColor.Height(), curFrameColor.Depth(), 1)
		opencv.CvtColor(curFrameColor, curFrame, opencv.CV_BGR2GRAY)
		curFWindow.ShowImage(curFrame)

		diff := opencv.CreateImage(curFrame.Width(), curFrame.Height(), curFrame.Depth(), 1)
		threshold := opencv.CreateImage(curFrame.Width(), curFrame.Height(), curFrame.Depth(), 1)
		opencv.AbsDiff(curFrame, firstFrame, diff)
		opencv.Threshold(diff, threshold, thVal, 255, opencv.CV_THRESH_BINARY)

		xWindow.ShowImage(threshold)

		if opencv.WaitKey(10) == 1048603 {
			break
		}
	}

}
