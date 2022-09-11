// mypkg は自作パッケージです。
package mypkg

import (
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

// ErrMsg はエラーメッセージを表すユーザー定義型です。
type ErrMsg string

func isDir(directory string) bool {
	fInfo, _ := os.Stat(directory)
	if fInfo == nil {
		fmt.Fprintf(os.Stderr, "error: %s: no such file or directory\n", directory)
		os.Exit(1)
	}
	if fInfo.IsDir() == false {
		return false
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

func checkError(err error, msg ErrMsg) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v: %v\n", msg, err)
		os.Exit(1)
	}
}

func jpgToPng(path string) {
	file, err := os.Open(path)
	checkError(err, "open")
	defer file.Close()

	img, _, err := image.Decode(file)
	checkError(err, "decode")

	png_file := strings.Replace(path, "jpg", "png", -1)
	out, err := os.Create(png_file)
	checkError(err, "create")
	defer out.Close()

	png.Encode(out, img)
}

func myWalk(dirname string) {
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
				jpgToPng(path)
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

// Convert は.jpgファイルを.pngファイルに変換する関数です。
func Convert() {
	flag.Parse()
	if dirname := flag.Arg(0); dirname == "" {
		fmt.Fprintf(os.Stderr, "error: invalid argument\n")
		os.Exit(0)
	} else if flag.Arg(1) != "" {
		fmt.Fprintf(os.Stderr, "error: multiple arguments\n")
		os.Exit(0)
	} else {
		myWalk(dirname)
	}
}
