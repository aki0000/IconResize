package main

import (
	"flag"
	"fmt"
	"image/jpeg"
	"image/png"
	"log"
	"os"

	"github.com/nfnt/resize"
)

// iOSアプリに必要なアイコンのサイズ
var iconSizes [13]uint = [13]uint{20, 29, 40, 58, 60, 76, 80, 87, 120, 152, 167, 180, 1024}
var iconTitles [13]string = [13]string{"20", "29", "40", "58", "60", "76", "80", "87", "120", "152", "167", "180", "1024"}

func main() {
	f := flag.String("file", "", "File Path")

	flag.Parse()
	filePath := *f
	// オプショナル引数のチェック
	if filePath == "" {
		fmt.Println("Need to use an optional: -file")
		return
	}

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return
	}

	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	for index, size := range iconSizes {
		m := resize.Resize(size, size, img, resize.NearestNeighbor)
		out, err := os.Create("test_" + iconTitles[index] + ".png")
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()

		jpeg.Encode(out, m, nil)
	}
}
