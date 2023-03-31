package poker

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"

	"golang.org/x/exp/slices"
)

// 今回切り出すテスト
// 1. "Q1 S7 D3 H3 Q6 D9"のように5枚(つまり文字列の長さが14文字)以外のカードが入力されたケース
// 2. "Q1 S7 H3 H3 Q6"のように同じカードが複数入力されるケース
// 3. "Q1 S7 D3 H3 D34"のように入力されたカードの数値が異常であるケース
//

// 1. "Q1 S7 D3 H3 Q6 D9"のように5枚(つまり文字列の長さが14文字)以外のカードが入力されたケース
func CheckInputLength(t string, err string) (string, string) {
	var b = utf8.RuneCountInString(t)
	fmt.Printf("%d\n", b)
	if b != 14 {
		err = "400(Bad Request): 入力値の長さが異なります"
	} else {
		fmt.Println("入力値の長さは正常です。")
	}
	return t, err
}

// 2. "Q1 S7 H3 H3 Q6"のように同じカードが複数入力されるケース
func ExDuplicatedCards(cards, err string) (string, string) {
	str := strings.Join(strings.Fields(cards), "")
	var dupSlice []string

	// 文字列をQ3,D4,H2のようにカードごとの要素にして配列に格納する。
	for i := 0; i < len(str); i += 2 {
		if i+1 < len(str) {
			dupSlice = append(dupSlice, str[i:i+2])
		} else {
			dupSlice = append(dupSlice, str[i:])
		}
	}

	// 重複を削除する。
	slices.Sort(dupSlice)
	unique := slices.Compact(dupSlice)

	// 重複を削除した後に、要素数が一致しなければエラーを返す。
	if len(unique) != 5 {
		err = "400(Bad Request): 同じカードを複数枚入力しています"
	} else {
		fmt.Println("重複するカードはありません。")
	}

	return cards, err
}

// 3. "Q1 S7 D3 H3 D34"のように入力されたカードの数値が異常であるケース
func MakeIntSlice(cards string, err string) ([]int, string) {

	r, _ := regexp.Compile("[^0-9| ]")

	q := strings.Split(r.ReplaceAllString(cards, ""), " ")

	s := [5]string{}
	copy(s[:], q)

	var ab = []int{}

	for _, i := range q {
		j, _ := strconv.Atoi(i)
		if j > 13 {
			err = "400(Bad Request): 入力されたカードの値が異常です"
		} else {
			fmt.Println("正常：入力されたカードの値は13以下です。")
		}
		ab = append(ab, j)
	}

	return ab, err
}
