package main

import (
	"fmt"
	"os"
	"io"
	"bufio"
	"strings"
	"strconv"
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

		// -n= で始まる引数があった場合、表示行数を取得する
		if strings.HasPrefix(arg, "-n=") {

			// 分割して表示行数を取得
			maxLine = getLine(arg)
		} else {
			// -n= で始まる引数以外は、スライスに追加する
			filesPath = append(filesPath, arg)
		}
		i++
	}

	// 表示処理を行う
	for _, filePath := range filesPath {
		filePrint(filePath, maxLine, len(filesPath) > 1)
	}
}
func getLine(value string) (int) {
	valueList := strings.Split(value, "=")
	line, err := strconv.Atoi(valueList[1])
	if err != nil {
		panic(err)
	}
	return line
}
