// iPhone/iOSアプリの各アイコンサイズのリサイズするプログラム

package main

import (
	"flag"
	"fmt"
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
	// 指定したファイルを開く
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return
	}
	// png形式にデコード
	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
	// 各サイズにリサイズ
	for index, size := range iconSizes {
		m := resize.Resize(size, size, img, resize.NearestNeighbor)
		out, err := os.Create("test_" + iconTitles[index] + ".png")
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()
		// png形式にエンコード
		png.Encode(out, m)
	}
}
