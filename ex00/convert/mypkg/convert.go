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

// IsPng は.pngファイルかどうかbool値で返す関数です。
func IsPng(path string) bool {
	if ext := filepath.Ext(path); ext == ".png" {
		return true
	} else {
		return false
	}
}

// CheckError はerrorが生じた際にその内容を出力しプログラムを終了する関数です。
func CheckError(err error, msg ErrMsg) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v: %v\n", msg, err)
		os.Exit(1)
	}
}

// JpgToPng は.jpgファイルから.pngファイルに変換する関数です。
func JpgToPng(path string) {
	file, err := os.Open(path)
	CheckError(err, "open")
	defer file.Close()

	img, _, err := image.Decode(file)
	CheckError(err, "decode")

	png_file := strings.Replace(path, "jpg", "png", -1)
	out, err := os.Create(png_file)
	CheckError(err, "create")
	defer out.Close()

	png.Encode(out, img)
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

// FindJpg は.jpgファイルを探す関数です。
func FindJpg(dirname string) {
	if _, err := os.Stat(dirname); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", TrimSpaceLeft(err))
		os.Exit(1)
	}
	err := filepath.Walk(dirname,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if filepath.Ext(path) == ".jpg" {
				JpgToPng(path)
			} else if info.IsDir() == false && IsPng(path) == false {
				fmt.Fprintf(os.Stderr, "error: %s is not a valid file\n", path)
			}
			return nil
		})
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", TrimSpaceLeft(err))
		os.Exit(1)
	}
}

// Convert はconvertのmainとなる関数です。
func Convert() {
	flag.Parse()
	if dirname := flag.Arg(0); dirname == "" {
		fmt.Fprintf(os.Stderr, "error: invalid argument\n")
		os.Exit(0)
	} else if flag.Arg(1) != "" {
		fmt.Fprintf(os.Stderr, "error: multiple arguments\n")
		os.Exit(0)
	} else {
		FindJpg(dirname)
	}
}
