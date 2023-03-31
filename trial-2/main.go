package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"regexp"
	"sort"
	"strings"

	"trial-2/poker"

	"github.com/gorilla/mux"
)

//ここから本実装

type Poker struct {
	Title string `json:"title"`
}

var pokers []Poker

func process(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var poker Poker

	_ = json.NewDecoder(r.Body).Decode(&poker)

	t := poker.Title

	y := testPass(t)

	json.NewEncoder(w).Encode(y)
}

// テストのため関数を切り出した。
func testPass(t string) string {

	var err string

	//文字列の長さが14でない場合にエラーを返す関数
	t, err = poker.CheckInputLength(t, err)

	p, err := poker.ExDuplicatedCards(t, err)

	num, err := poker.MakeIntSlice(p, err)

	u, err := makeStringSlice(p, err)

	flush := judgeFlush(p, u)

	straight := judgeStraight(num)

	dup := findDup(num)

	j := sliceToArray(dup)

	a := judge(j, num, straight, flush)

	var x string

	if err == "" {
		x = a
	} else {
		x = err
	}

	return x

}

func main() {
	// Initiate Router
	r := mux.NewRouter()

	r.HandleFunc("/", process).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", r))
}

// テストのためpokerモジュールに切り出した
// func checkInputLength(t string, err string) (string, string) {
// 	var b = utf8.RuneCountInString(t)
// 	fmt.Printf("%d\n", b)
// 	if b != 14 {
// 		err = "400(Bad Request): 入力値の長さが異なります"
// 	} else {
// 		fmt.Println("入力値の長さは正常です。")
// 	}
// 	return t, err
// }

// テストのためpokerモジュールに切り出した
// func exDuplicatedCards(cards string, err string) (string, string) {
// 	str := strings.Join(strings.Fields(cards), "")
// 	var dupSlice []string

// 	// 文字列をQ3,D4,H2のようにカードごとの要素にして配列に格納する。
// 	for i := 0; i < len(str); i += 2 {
// 		if i+1 < len(str) {
// 			dupSlice = append(dupSlice, str[i:i+2])
// 		} else {
// 			dupSlice = append(dupSlice, str[i:])
// 		}
// 	}

// 	// 重複を削除する。
// 	slices.Sort(dupSlice)
// 	unique := slices.Compact(dupSlice)

// 	// 重複を削除した後に、要素数が一致しなければエラーを返す。
// 	if len(unique) != 5 {
// 		err = "400(Bad Request): 同じカードを複数枚入力しています"
// 	} else {
// 		fmt.Println("重複するカードはありません。")
// 	}

// 	return cards, err
// }

// テストのためpokerモジュールに切り出した
// func makeIntSlice(cards string, err string) ([]int, string) {

// 	r, _ := regexp.Compile("[^0-9| ]")

// 	q := strings.Split(r.ReplaceAllString(cards, ""), " ")

// 	s := [5]string{}
// 	copy(s[:], q)

// 	var ab = []int{}

// 	for _, i := range q {
// 		j, _ := strconv.Atoi(i)
// 		if j > 13 {
// 			err = "400(Bad Request): 入力されたカードの値が異常です"
// 		} else {
// 			fmt.Println("正常：入力されたカードの値は13以下です。")
// 		}
// 		ab = append(ab, j)
// 	}

// 	sort.Ints(ab)

// 	if len(ab) == 5 {
// 		fmt.Println("正常：入力されたカードの数字の数は正常です")
// 	} else {
// 		err = "400(Bad Request): 入力されたカードの枚数が異常です"
// 	}

// 	return ab, err
// }

// 文字列から文字のみを抜き出して文字列型のスライスに変換するコード
func makeStringSlice(cards string, err string) ([]string, string) {
	r, _ := regexp.Compile("[^A-Z| ]")

	p := strings.Split(r.ReplaceAllString(cards, ""), " ")

	var u []string

	u, err = suitValidation(p, err)

	if len(p) == 5 {
		fmt.Println("正常：入力されたカードのスートの数は正常です")
	} else {
		err = "400(Bad Request): 入力されたカードの枚数が異常です"
		return u, err
	}

	return u, err
}

func judgeFlush(cards string, test2 []string) int {
	flush := 0

	if strings.Count(cards, test2[0]) == len(test2) {
		flush = 1
	} else {
		flush = 0
	}

	return flush

}

func judgeStraight(num []int) int {
	straight := 0

	if num[1] == num[0]+1 && num[2] == num[0]+2 && num[3] == num[0]+3 && num[4] == num[0]+4 || num[0] == 1 && num[1] == 1 {
		straight = 1
	} else {
		straight = 0
	}
	return straight
}

func findGreat(array []int) int {
	greatest := 0
	for i := 0; i < len(array); i++ {
		if array[i] > greatest {
			greatest = array[i]
		}
	}
	return greatest + 1
}

func findDup(array []int) []int {
	dup := make([]int, findGreat(array))

	for i := 0; i < len(array); i++ {
		dup[array[i]] += 1
	}

	sort.Ints(dup)
	return dup
}

func sliceToArray(dup []int) []int {
	j := []int{}

	for _, i := range dup {
		if i != 0 {
			j = append(j, i)
		}
	}
	return j
}

func judge(j []int, num []int, straight, flush int) string {
	//二次元スライスを作成する。2ペアなどの判定で利用。

	p := [][]int{{1, 4}, {2, 3}, {1, 1, 3}, {1, 2, 2}, {1, 1, 1, 2}, {1, 1, 1, 1, 1}}
	var a string = ""

	if reflect.DeepEqual(j, p[0]) {
		a = "4カード"
	} else if reflect.DeepEqual(j, p[1]) {
		a = "フルハウス"
	} else if reflect.DeepEqual(j, p[2]) {
		a = "スリーカード"
	} else if reflect.DeepEqual(j, p[3]) {
		a = "ツーペア"
	} else if reflect.DeepEqual(j, p[4]) {
		a = "ワンペア"
	} else if reflect.DeepEqual(j, p[5]) {
		// 比較をするにはスライスを配列に変換する必要がある。numスライスをs配列にコピーする。
		s := [5]int{}
		copy(s[:], num)

		if s == [5]int{1, 10, 11, 12, 13} && flush == 1 {
			a = "ロイヤルストレートフラッシュ"
		} else if s != [5]int{1, 10, 11, 12, 13} && flush == 1 && straight == 1 {
			a = "ストレートフラッシュ"
		} else if flush == 1 {
			a = "フラッシュ"
		} else if straight == 1 {
			a = "ストレート"
		} else {
			a = "ノーペア"
		}
	}
	return a
}

func suitValidation(y []string, err string) ([]string, string) {
	// 排除したい文字がなければ正しいパラメータを表示するだけ
	for _, n := range y {
		switch n {
		case "Q":
			// 正常系
		case "D":
			// 正常系
		case "S":
			// 正常系
		case "H":
			// 正常系
		default:
			err = "400(Bad Request): 入力されたカードのスートが異常です"
			return y, err
		}
	}
	return y, err
}
