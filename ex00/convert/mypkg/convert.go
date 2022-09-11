package mypkg

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

func Do() {
	fmt.Println("hello world")
}

func isDir(directory string) bool {
	fInfo, _ := os.Stat(directory)
	if fInfo.IsDir() == false {
		return false //It's file
	}
	return true
}

func isPng(str string) bool {
	if ext := filepath.Ext(str); ext == ".png" {
		return true
	} else {
		return false
	}
}

func check_error(err error, msg string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", msg, err)
		os.Exit(1)
	}
}

func Jpg_to_png(path string) {
	file, err := os.Open(path)
	check_error(err, "open")
	defer file.Close()

	img, _, err := image.Decode(file)
	check_error(err, "decode")

	png_file := strings.Replace(path, "jpg", "png", -1)
	out, err := os.Create(png_file)
	check_error(err, "create")
	defer out.Close()

	// 画像ファイル出力
	//    jpeg.Encode(out, img, nil)
	png.Encode(out, img)
}

// JPGファイルを探し出す
func MyWalk(dirname string) {
	if isDir(dirname) == false {
		fmt.Fprintf(os.Stderr, "error: %s is not directory\n", dirname)
		os.Exit(0)
	}
	err := filepath.Walk(dirname,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if filepath.Ext(path) == ".jpg" {
				Jpg_to_png(path)
			} else if isDir(path) == false && isPng(path) == false {
				fmt.Fprintf(os.Stderr, "error: %s is not a valid file\n", path)
			}
			return nil
		})
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s: no such file or directory\n", dirname)
		os.Exit(1)
	}
}
