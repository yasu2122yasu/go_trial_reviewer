/*
Q20
任意のディレクトリの配下にmain.goを作成してください。
任意のディレクトリの配下にhelloディレクトリを作成し、その配下にhello.goを作成してください。
hello.go内に"Hello, World"を出力する関数を作り、main関数から呼び出してください。
*/
package main

import (
	"20/lib"
)

func main() {
	lib.Hello()
}
