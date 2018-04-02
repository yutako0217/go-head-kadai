package main

import (
	"fmt"
	"os"
	"strconv"
	"io"
	"bufio"
)

var lineDefault = 10

func filePrint(filePath string, maxLine int, fileDisp bool) {
	if fileDisp {
		fmt.Printf("==> %s <==\n", filePath)
	}

	var file *os.File
	var err error

	// ファイルオープン
	file, err = os.Open(filePath)
	if err != nil {
		fmt.Printf("%s : No such file or directory \n", filePath)
		return
	}

	// 関数を抜ける時にクローズ
	defer file.Close()

	// bufferio
	reader := bufio.NewReaderSize(file, 4096)

	// 表示処理
	l := 0
	for l < maxLine {
		// 1行読む
		line, _, err := reader.ReadLine()
		fmt.Println(string(line))
		l++
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}
}

func main() {
	var maxLine int
	var filesPath []string
	var err error

	// デフォルトの表示行数
	maxLine = lineDefault
	// 引数を取得
	args := os.Args

	// 引数のチェック（-nとその値は除去する）
	// argsの0は自身になるので1から開始
	i := 1
	for i < len(args) {
		// fmt.Println(lineNum)
		arg := args[i]
		if arg == "-n" {
			//-nの次が数字かどうかのチェック
			maxLine, err = strconv.Atoi(args[i+1])
			if err == nil {
				// -n の次の引数は行数が取得できているので飛ばす
				i++
			} else {
				// -n オプションの次に、数字以外が来ている。
				fmt.Println("head: option requires an argument -- n")
				os.Exit(1)
			}
		} else {
			// スライスに追加する
			filesPath = append(filesPath, arg)
		}
		i++
	}

	// 表示処理を行う
	for _, filePath := range filesPath {
		filePrint(filePath, maxLine, len(filesPath) > 1)
	}
}
