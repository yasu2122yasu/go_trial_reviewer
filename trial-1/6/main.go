/*
Q6.
自分の誕生日の下二桁を用意し、変数に代入してください。
15より小さければ、"月の前半が誕生日です"、
15より大きければ、"月の後半が誕生日です"
と出力してください。
ただし、if文/else文を使用すること。
*/
package main

import "fmt"

func main() {
	var b int = 12
	if b < 15 {
		fmt.Println("月の前半が誕生日です")
	} else {
		fmt.Println("月の後半が誕生日です")
	}
}
