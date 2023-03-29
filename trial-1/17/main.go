/*
Q17
int型のポインタ変数aを用意してください。
int型の変数bを用意して任意の数字を代入してください。
bをaに代入してください。（変数aとbではデータ型が異なるので注意すること）
aに代入された値とアドレスを出力してください。
*/
package main

import "fmt"

func main() {
	var a *int
	b := 7
	a = &b

	fmt.Println(*a)
	fmt.Println(a)
}
