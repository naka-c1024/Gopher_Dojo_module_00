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

var exitStatus int

// IsPng は.pngファイルかどうかbool値で返す関数です。
func IsPng(path string) bool {
	return filepath.Ext(path) == ".png"
}

// TrimSpaceLeft はエラーメッセージにおいて不要なスペースから左部分を除く関数です。
func TrimSpaceLeft(err error) string {
	str := err.Error()
	spaceIndex := strings.Index(str, " ")
	if spaceIndex == -1 {
		return str
	}
	return str[spaceIndex+1:]
}

// JPGtoPng は.jpgファイルから.pngファイルに変換する関数です。
func JPGtoPng(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	var pngFile string
	switch filepath.Ext(path) {
	case ".jpg":
		pngFile = strings.TrimSuffix(path, ".jpg") + ".png"
	case ".jpeg":
		pngFile = strings.TrimSuffix(path, ".jpeg") + ".png"
	}
	out, err := os.Create(pngFile)
	if err != nil {
		return err
	}

	err = png.Encode(out, img)
	if err != nil {
		return err
	}

	return nil
}

// FindJPG は.jpgファイルを探す関数です。
func FindJPG(dirname string) {
	err := filepath.Walk(dirname,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if filepath.Ext(path) == ".jpg" || filepath.Ext(path) == ".jpeg" {
				err = JPGtoPng(path)
				if err != nil {
					fmt.Fprintf(os.Stderr, "error: %s: %s\n", path, err.Error())
					exitStatus = 1
				}
			} else if info.IsDir() == false && IsPng(path) == false {
				fmt.Fprintf(os.Stderr, "error: %s is not a valid file\n", path)
				exitStatus = 1
			}
			return nil
		})
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		exitStatus = 1
	}
}

// Convert はconvertのmainとなる関数です。
func Convert() {
	flag.Parse()
	if dirname := flag.Arg(0); dirname == "" {
		fmt.Fprintf(os.Stderr, "error: invalid argument\n")
		os.Exit(1)
	}
	for i := 0; flag.Arg(i) != ""; i++ {
		if _, err := os.Stat(flag.Arg(i)); err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", TrimSpaceLeft(err))
			exitStatus = 1
			continue
		}
		FindJPG(flag.Arg(i))
	}
	os.Exit(exitStatus)
}
